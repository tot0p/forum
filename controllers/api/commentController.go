package api

import (
	"encoding/json"
	"forum/models"
	"forum/repository"
	"forum/tools/authorization"
	"forum/tools/session"
	"net/http"
	"strings"
)

//Function to get all the comment
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

//Function to get a comment using the id
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

//Function to delete a comment using the id
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

//Function to create a new Comment
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
		params["parent"].(string),
	)
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	w.Write([]byte("{\"msg\":\"success\"}"))
}

//Function to modify a comment using the id
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

func CommentLike(paramsURL map[string]string, params map[string]interface{}, w http.ResponseWriter, r *http.Request) {
	sess, err := session.GlobalSessions.Provider.SessionRead(authorization.GetAuthorizationBearer(w, r))
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	if !session.GlobalSessions.SessionExist(sess.SessionID()) {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"Session Invalid\"}"))
		return
	}
	UserUUID, err := sess.Get("UUID")
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	User, err := repository.GetUser("UUID", UserUUID.(string))
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	Comment, err := repository.GetComment("id", paramsURL["id"])
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	remove := false
	new := ""
	upvote := Comment.ConvertUpVotes()
	for _, i := range upvote {
		if User.UUID == i {
			remove = true
		} else {
			new += "#" + i
		}
	}
	if !remove {
		new = ""
		for _, i := range Comment.ConvertDownVotes() {
			if User.UUID == i {
				remove = true
			} else {
				new += "#" + i
			}
		}
		if remove {
			Comment.DownVotes = new
		}
		Comment.UpVotes += "#" + User.UUID
	} else {
		Comment.UpVotes = new
	}
	repository.PostComment(*Comment, "id", paramsURL["id"])
	w.Write([]byte("{\"msg\":\"success\"}"))
}

func CommentHate(paramsURL map[string]string, params map[string]interface{}, w http.ResponseWriter, r *http.Request) {
	sess, err := session.GlobalSessions.Provider.SessionRead(authorization.GetAuthorizationBearer(w, r))
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	if !session.GlobalSessions.SessionExist(sess.SessionID()) {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"Session Invalid\"}"))
		return
	}
	UserUUID, err := sess.Get("UUID")
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	User, err := repository.GetUser("UUID", UserUUID.(string))
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	Comment, err := repository.GetComment("id", paramsURL["id"])
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	remove := false
	new := ""
	Downvote := Comment.ConvertDownVotes()
	for _, i := range Downvote {
		if User.UUID == i {
			remove = true
		} else {
			new += "#" + i
		}
	}
	if !remove {
		new = ""
		for _, i := range Comment.ConvertUpVotes() {
			if User.UUID == i {
				remove = true
			} else {
				new += "#" + i
			}
		}
		if remove {
			Comment.UpVotes = new
		}
		Comment.DownVotes += "#" + User.UUID
	} else {
		Comment.DownVotes = new
	}
	repository.PostComment(*Comment, "id", paramsURL["id"])
	w.Write([]byte("{\"msg\":\"success\"}"))
}

//Function to Count the number of subject on the website
func CommentCount(paramsURL map[string]string, params map[string]interface{}, w http.ResponseWriter, r *http.Request) {
	Comment, err := repository.GetComment("id", paramsURL["id"])
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	Votes := models.Vote{UpVote: strings.Count(Comment.UpVotes, "#"), DownVote: strings.Count(Comment.DownVotes, "#")}
	result, err := json.Marshal(Votes)
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	w.Write(result)
}

//function to say if a user already like or dislike subject
func UserLikeOrHateComment(paramsURL map[string]string, params map[string]interface{}, w http.ResponseWriter, r *http.Request) {
	sess, err := session.GlobalSessions.Provider.SessionRead(authorization.GetAuthorizationBearer(w, r))
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	if !session.GlobalSessions.SessionExist(sess.SessionID()) {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"Session Invalid\"}"))
		return
	}
	UserUUID, err := sess.Get("UUID")
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	User, err := repository.GetUser("UUID", UserUUID.(string))
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	Comment, err := repository.GetComment("id", paramsURL["id"])
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	for _, i := range Comment.ConvertUpVotes() {
		if i == User.UUID {
			w.Write([]byte("{\"downvote\":false,\"upvote\":true}"))
			return
		}
	}
	for _, i := range Comment.ConvertDownVotes() {
		if i == User.UUID {
			w.Write([]byte("{\"downvote\":true,\"upvote\":false}"))
			return
		}
	}
	w.Write([]byte("{\"downvote\":false,\"upvote\":false}"))
}

func GetCommentsByPostId(paramsURL map[string]string, params map[string]interface{}, w http.ResponseWriter, r *http.Request) {
	allComments, err := repository.GetCommentsByPostId(paramsURL["id"])
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
