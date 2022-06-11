package api

import (
	"encoding/json"
	"forum/models"
	"forum/repository"
	"forum/tools/session"
	"net/http"
)

//Get the number of users created
func GetNBUser(paramsURL map[string]string, params map[string]interface{}, w http.ResponseWriter, r *http.Request) {
	userById, err := repository.Count("user")
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	user, err := json.Marshal(userById)
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	w.Write(user)
}

//Get the number of subjects created
func GetNBSubject(paramsURL map[string]string, params map[string]interface{}, w http.ResponseWriter, r *http.Request) {
	subjectById, err := repository.Count("subject")
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	subject, err := json.Marshal(subjectById)
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	w.Write(subject)
}

//Get the number of posts created
func GetNBPost(paramsURL map[string]string, params map[string]interface{}, w http.ResponseWriter, r *http.Request) {
	postById, err := repository.Count("post")
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	post, err := json.Marshal(postById)
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	w.Write(post)
}

//Get all counters
func GetCount(paramsURL map[string]string, params map[string]interface{}, w http.ResponseWriter, r *http.Request) {
	counters := models.AllCount{}
	count, err := repository.Count("post")
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	counters.Post = count.Nb
	count, err = repository.Count("subject")
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	counters.Subject = count.Nb
	count, err = repository.Count("user")
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	counters.User = count.Nb
	counters.Session = session.GlobalSessions.GetNBSession()
	user, err := json.Marshal(counters)
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	w.Write(user)
}
