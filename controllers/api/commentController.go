package api

import (
	"encoding/json"
	"forum/repository"
	"forum/tools/authorization"
	"forum/tools/session"
	"net/http"
)

//Get All Comment
func GetAllComment(paramsURL map[string]string, params map[string]interface{}, w http.ResponseWriter, r *http.Request) {
	allComments, err := repository.GetAllComment()
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	comments, err := json.Marshal(allComments)
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	w.Write(comments)
}

//Get a comment by an Id
func GetCommentById(paramsURL map[string]string, params map[string]interface{}, w http.ResponseWriter, r *http.Request) {
	commentById, err := repository.GetComment("id", paramsURL["id"])
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	comment, err := json.Marshal(commentById)
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	w.Write(comment)
}

//Delete a comment by an Id
func DeleteCommentById(paramsURL map[string]string, params map[string]interface{}, w http.ResponseWriter, r *http.Request) {
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
	err = repository.DeleteComment("id", paramsURL["id"])
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	w.Write([]byte("{\"msg\":\"success\"}"))
}

//Create a comment
func CreateComment(paramsURL map[string]string, params map[string]interface{}, w http.ResponseWriter, r *http.Request) {
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
	err = repository.CreateComment(
		params["content"].(string),
		UUID.(string),
	)
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	w.Write([]byte("{\"msg\":\"success\"}"))
}

//Put comment by an Id
func PutCommentById(paramsURL map[string]string, params map[string]interface{}, w http.ResponseWriter, r *http.Request) {
	sess, err := session.GlobalSessions.Provider.SessionRead(authorization.GetAuthorizationBearer(w, r))
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	if !session.GlobalSessions.SessionExist(sess.SessionID()) {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"Session Invalid\"}"))
		return
	}
	commentById, err := repository.GetComment("id", paramsURL["id"])
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
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
		if commentById.Owner != UUID.(string) {
			w.Write([]byte("{\"err\":\"500\",\"msg\":\"You are not the owner of this comment\"}"))
			return
		}
	}
	if params["content"] != nil {
		commentById.Content = params["content"].(string)
	}
	if params["upvotes"] != nil {
		commentById.UpVotes = params["upvotes"].(string)
	}
	if params["downvotes"] != nil {
		commentById.DownVotes = params["downvotes"].(string)
	}
	if params["responses"] != nil {
		commentById.Responses = params["responses"].(string)
	}
	err = repository.PostComment(*commentById, "id", commentById.Id)
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	comment, err := json.Marshal(commentById)
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	if comment != nil {
		w.Write(comment)
	}
}
