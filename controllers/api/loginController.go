package api

import (
	"encoding/json"
	"forum/models"
	"forum/repository"
	"forum/tools/authorization"
	"forum/tools/session"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

//Function to login an user on the website
func UserLogin(paramsURL map[string]string, params map[string]interface{}, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	if params["username"] == nil || params["password"] == nil {
		w.WriteHeader(500)
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"Username or Password cant be nil\"}"))
		return
	}
	user, err := repository.GetUser("username", params["username"].(string))
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	pass := params["password"].(string)
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass)) != nil {
		w.WriteHeader(500)
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"Wrong password\"}"))
		return
	}
	sess := session.GlobalSessions.SessionStart(w, r)
	sess.Set("UUID", user.UUID)
	data := make(map[string]string)
	data["SID"] = sess.SessionID()
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	w.WriteHeader(http.StatusOK)
	rep, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	w.Write(rep)
}

//Function to get the number of users connected on the website
func GetNBUserConnected(paramsURL map[string]string, params map[string]interface{}, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	count, err := json.Marshal(models.Count{Nb: session.GlobalSessions.GetNBSession()})
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	w.Write(count)
}

//Function to stop a session
func DeleteSession(paramsURL map[string]string, params map[string]interface{}, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	sess, err := session.GlobalSessions.Provider.SessionRead(authorization.GetAuthorizationBearer(w, r))
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	if !session.GlobalSessions.SessionExist(sess.SessionID()) {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"Session Invalid\"}"))
		return
	}
	session.GlobalSessions.SessionDestroy(sess.SessionID())
	w.Write([]byte("{\"msg\":\"success\"}"))
}
