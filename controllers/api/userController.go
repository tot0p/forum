package api

import (
	"encoding/hex"
	"encoding/json"
	"forum/models"
	"forum/repository"
	"forum/tools/authorization"
	"forum/tools/session"
	"forum/tools/verif"
	"net/http"
	"net/url"
	"strings"
)

//Function to get all users
func GetAllUsers(paramsURL map[string]string, params map[string]interface{}, w http.ResponseWriter, r *http.Request) {
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
	UUID, err := sess.Get("UUID")
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	user, err := repository.GetUser("UUID", UUID.(string))
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	if user.Role != "admin" {
		w.Write([]byte("{\"err\":\"403\",\"msg\":\"Forbidden\"}"))
		return
	}
	allUsers, err := repository.GetAllUser()
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	for i := range *allUsers {
		(*allUsers)[i].Sec()
	}
	users, err := json.Marshal(allUsers)
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	w.Write(users)
}

//Function to search an user
func SearchUser(paramsURL map[string]string, params map[string]interface{}, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	allUsers, err := repository.GetAllUser()
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	searchUser := []models.User{}
	word, err := url.QueryUnescape(paramsURL["word"])
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	for i := range *allUsers {
		(*allUsers)[i].Sec()
		if strings.Contains(strings.ToLower((*allUsers)[i].Username), strings.ToLower(word)) && (*allUsers)[i].UUID != "0" {
			searchUser = append(searchUser, (*allUsers)[i])
		}
	}
	users, err := json.Marshal(searchUser)
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	w.Write(users)
}

//Function to get an user by his id
func GetUserById(paramsURL map[string]string, params map[string]interface{}, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	userById, err := repository.GetUser("UUID", paramsURL["id"])
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	userById.Sec()
	user, err := json.Marshal(userById)
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	w.Write(user)
}

//Function to get an user by his username
func GetUserByUsername(paramsURL map[string]string, params map[string]interface{}, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	username, err := url.QueryUnescape(paramsURL["username"])
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	userByUsername, err := repository.GetUser("username", username)
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	userByUsername.Sec()
	user, err := json.Marshal(userByUsername)
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	w.Write(user)
}

//Function to create an user
func CreateUser(paramsURL map[string]string, params map[string]interface{}, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	var (
		oauth, password string
	)
	if params["oauthtoken"] != nil && params["password"] == nil {
		oauth = params["oauthtoken"].(string)
	} else if params["password"] != nil && params["oauthtoken"] == nil {
		password = params["password"].(string)
	} else {
		w.Write([]byte("{\"\":\"\"}"))
		return
	}

	if !verif.PasswordVerif(params["password"].(string)) || !verif.EmailVerif(params["email"].(string)) || !verif.DateVerif(params["birthdate"].(string)) {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte("{\"err\":\"406\",\"msg\":\"password doesn't have ...\"}")) // à mettre les descs
		return
	}

	err := repository.CreateUser(
		params["username"].(string),
		password,
		params["email"].(string),
		params["firstname"].(string),
		params["lastname"].(string),
		params["birthdate"].(string),
		oauth,
		params["genre"].(string),
		params["bio"].(string),
		params["riotid"].(string),
		[]byte(params["profilepicture"].(string)),
		0)
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}

	w.Write([]byte("{\"msg\":\"success\"}"))
}

//Delete an user by his id
func DeleteUserById(paramsURL map[string]string, params map[string]interface{}, w http.ResponseWriter, r *http.Request) {
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
	UUID, err := sess.Get("UUID")
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	user, err := repository.GetUser("UUID", UUID.(string))
	if user.Role == "admin" {
		err = repository.DeleteUser("UUID", paramsURL["id"])
	}
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	w.Write([]byte("{\"msg\":\"success\"}"))
}

//Function to modify an user using his id
func PutUserById(paramsURL map[string]string, params map[string]interface{}, w http.ResponseWriter, r *http.Request) {
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
	userById := new(models.User)
	UUID, err := sess.Get("UUID")
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	user, err := repository.GetUser("UUID", UUID.(string))
	if user.Role == "admin" && params["id"] != nil {
		userById, err = repository.GetUser("UUID", params["id"].(string))
	} else {
		userById = user
	}

	if !verif.EmailVerif(params["email"].(string)) || !verif.DateVerif(params["birthdate"].(string)) {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte("{\"err\":\"406\",\"msg\":\"password doesn't have ...\"}")) // à mettre les descs
		return
	}
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	if params["profilepicture"] != nil {
		if params["profilepicture"].(string) != "" {
			data, err := hex.DecodeString(params["profilepicture"].(string))
			if err != nil {
				w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
				return
			}
			userById.ProfilePicture = data
		}
	}
	if params["username"] != nil {
		userById.Username = params["username"].(string)
	}
	if params["password"] != nil {
		if params["password"].(string) != "" {
			if !verif.PasswordVerif(params["password"].(string)) {
				w.WriteHeader(http.StatusNotAcceptable)
				w.Write([]byte("{\"err\":\"406\",\"msg\":\"password doesn't have ...\"}")) // à mettre les descs
				return
			}
			userById.Password = repository.HashPassword(params["password"].(string))
		}
	}
	if params["email"] != nil {
		if params["email"].(string) != "" {
			userById.Email = params["email"].(string)
		}
	}
	if params["firstname"] != nil {
		userById.FirstName = params["firstname"].(string)
	}
	if params["lastname"] != nil {
		userById.LastName = params["lastname"].(string)
	}
	if params["riotid"] != nil {
		userById.RiotId = params["riotid"].(string)
	}
	if params["birthdate"] != nil {
		if params["birthdate"].(string) != "" {
			userById.BirthDate = params["birthdate"].(string)
		}
	}
	if params["oauthtoken"] != nil {
		userById.OauthToken = params["oauthtoken"].(string)
	}
	if params["genre"] != nil {
		userById.Genre = params["genre"].(string)
	}
	if params["role"] != nil {
		userById.Role = params["role"].(string)
	}
	if params["title"] != nil {
		userById.Title = params["title"].(string)
	}
	if params["bio"] != nil {
		userById.Bio = params["bio"].(string)
	}
	if params["premium"] != nil {
		userById.Premium = int(params["premium"].(float64))
	}
	err = repository.PostUser(*userById, "UUID", userById.UUID)
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	UpdatedUser, err := json.Marshal(userById)
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	if UpdatedUser != nil {
		w.Write(UpdatedUser)
	}
}

//Function to get an user's username
func GetUserUsername(paramsURL map[string]string, params map[string]interface{}, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	user := make(map[string]string)
	if paramsURL["id"] == "" {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"UUID Not Supplied\"}"))
		return
	}
	user["username"] = repository.GetUserUsername(paramsURL["id"])
	username, err := json.Marshal(user)
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	w.Write(username)
}

//Function to get your user
func GetUserMe(paramsURL map[string]string, params map[string]interface{}, w http.ResponseWriter, r *http.Request) {
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
	UUID, err := sess.Get("UUID")
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	userById, err := repository.GetUser("UUID", UUID.(string))
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	userById.Sec()
	user, err := json.Marshal(userById)
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	w.Write(user)
}
