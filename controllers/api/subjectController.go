package api

import (
	"encoding/hex"
	"encoding/json"
	"forum/models"
	"forum/repository"
	"forum/tools"
	"forum/tools/authorization"
	"forum/tools/session"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

//Get all subjects
func GetAllSubject(paramsURL map[string]string, params map[string]interface{}, w http.ResponseWriter, r *http.Request) {
	allSubjects, err := repository.GetAllSubject()
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	subjects, err := json.Marshal(allSubjects)
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	w.Write(subjects)
}

//Get a subject by his Id
func GetSubjectById(paramsURL map[string]string, params map[string]interface{}, w http.ResponseWriter, r *http.Request) {
	subjectById, err := repository.GetSubject("id", paramsURL["id"])
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

//Search a subject
func SearchSubject(paramsURL map[string]string, params map[string]interface{}, w http.ResponseWriter, r *http.Request) {
	allSubjects, err := repository.GetAllSubject()
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	searchSubjects := []models.Subject{}
	word, err := url.QueryUnescape(paramsURL["word"])
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	for i := range *allSubjects {
		if strings.Contains(strings.ToLower((*allSubjects)[i].Title), strings.ToLower(word)) {
			searchSubjects = append(searchSubjects, (*allSubjects)[i])
		}
	}
	subjects, err := json.Marshal(searchSubjects)
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	w.Write(subjects)
}

//Create a subject
func CreateSubject(paramsURL map[string]string, params map[string]interface{}, w http.ResponseWriter, r *http.Request) {
	sess, err := session.GlobalSessions.Provider.SessionRead(authorization.GetAuthorizationBearer(w, r))
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	if !session.GlobalSessions.SessionExist(sess.SessionID()) {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"Session Invalid\"}"))
		return
	}
	var tags = []string{}
	for _, i := range params["tags"].([]interface{}) {
		tags = append(tags, i.(string))
	}
	UUID, err := sess.Get("UUID")
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	var imageData []byte
	if params["image"].(string) == "" {
		imageData = tools.CreateImg()
	} else {
		imageData, err = hex.DecodeString(params["image"].(string))
		if err != nil {
			w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
			return
		}
	}
	err = repository.CreateSubject(
		params["title"].(string),
		params["description"].(string),
		UUID.(string),
		imageData,
		int(params["nsfw"].(float64)),
		tags,
	)
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	w.Write([]byte("{\"msg\":\"success\"}"))
}

//Delete a subject by his Id
func DeleteSubjectById(paramsURL map[string]string, params map[string]interface{}, w http.ResponseWriter, r *http.Request) {
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
	err = repository.DeleteSubject("id", paramsURL["id"])
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	w.Write([]byte("{\"msg\":\"success\"}"))
}

//Modify a subject by his Id
func PutSubjectsById(paramsURL map[string]string, params map[string]interface{}, w http.ResponseWriter, r *http.Request) {
	sess, err := session.GlobalSessions.Provider.SessionRead(authorization.GetAuthorizationBearer(w, r))
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	if !session.GlobalSessions.SessionExist(sess.SessionID()) {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"Session Invalid\"}"))
		return
	}
	subjectById, err := repository.GetSubject("id", params["id"].(string))
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
		if subjectById.Owner != user.UUID {
			w.Write([]byte("{\"err\":\"500\",\"msg\":\"You are not the owner of this subject\"}"))
			return
		}
	}
	if params["title"] != nil {
		subjectById.Title = params["title"].(string)
	}
	if params["description"] != nil {
		subjectById.Description = params["description"].(string)
	}
	if params["nsfw"] != nil {
		subjectById.NSFW = int(params["nsfw"].(float64))
	}
	if params["image"] != nil {
		if params["image"].(string) != "" {
			imageData, err := hex.DecodeString(params["image"].(string))
			if err != nil {
				w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
				return
			}
			subjectById.Image = imageData
		}
	}
	if params["tags"] != nil {
		subjectById.Tags = params["tags"].(string)
	}
	err = repository.PostSubject(*subjectById, "id", subjectById.Id)
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	subject, err := json.Marshal(subjectById)
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	if subject != nil {
		w.Write(subject)
	}
}

//Check if a subject is in a array of subjects
func ContainsSubject(AllSubjects []models.Subject, subject models.Subject) bool {
	for i := range AllSubjects {
		if AllSubjects[i].Id == subject.Id {
			return true
		}
	}
	return false
}

//Get a random subject
func GetNbRandomSubject(paramsURL map[string]string, params map[string]interface{}, w http.ResponseWriter, r *http.Request) {
	allPosts, err := repository.GetAllSubject()
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	nb, err := strconv.Atoi(paramsURL["nb"])
	if err != nil {
		w.Write([]byte("{\"err\":\"403\",\"msg\":\"Need number in url \"}"))
		return
	}
	if nb > len(*allPosts) {
		posts, err := json.Marshal(allPosts)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
			return
		}
		w.Write(posts)
		return
	}
	result := []models.Subject{}
	for i := 0; i < nb; i++ {
		temp := rand.Intn(len(*allPosts))
		for ContainsSubject(result, (*allPosts)[temp]) {
			temp = rand.Intn(len(*allPosts))
		}
		result = append(result, (*allPosts)[temp])
	}
	posts, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	w.Write(posts)
}

//Get the last updated subject
func GetSubjectLastUpdate(paramsURL map[string]string, params map[string]interface{}, w http.ResponseWriter, r *http.Request) {
	allPosts, err := repository.GetSubjectLastUpdate()
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	nb, err := strconv.Atoi(paramsURL["nb"])
	if err != nil {
		w.Write([]byte("{\"err\":\"403\",\"msg\":\"Need number in url \"}"))
		return
	}
	if nb > len(*allPosts) {
		posts, err := json.Marshal(allPosts)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
			return
		}
		w.Write(posts)
		return
	}
	*allPosts = (*allPosts)[:nb]

	posts, err := json.Marshal(allPosts)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	w.Write(posts)
}

func SubjectLike(paramsURL map[string]string, params map[string]interface{}, w http.ResponseWriter, r *http.Request) {
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
	Subject, err := repository.GetSubject("id", paramsURL["id"])
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	remove := false
	new := ""
	upvote := Subject.ConvertUpVotes()
	for _, i := range upvote {
		if User.UUID == i {
			remove = true
		} else {
			new += "#" + i
		}
	}
	if !remove {
		Subject.UpVotes += "#" + User.UUID
	} else {
		Subject.UpVotes = new
	}
	repository.PostSubject(*Subject, "id", paramsURL["id"])
	w.Write([]byte("{\"msg\":\"success\"}"))
}

func SubjectHate(paramsURL map[string]string, params map[string]interface{}, w http.ResponseWriter, r *http.Request) {
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
	Subject, err := repository.GetSubject("id", paramsURL["id"])
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	remove := false
	new := ""
	Downvote := Subject.ConvertDownVotes()
	for _, i := range Downvote {
		if User.UUID == i {
			remove = true
		} else {
			new += "#" + i
		}
	}
	if !remove {
		Subject.DownVotes += "#" + User.UUID
	} else {
		Subject.DownVotes = new
	}
	repository.PostSubject(*Subject, "id", paramsURL["id"])
	w.Write([]byte("{\"msg\":\"success\"}"))
}

func SubjectCount(paramsURL map[string]string, params map[string]interface{}, w http.ResponseWriter, r *http.Request) {
	Subject, err := repository.GetSubject("id", paramsURL["id"])
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	Votes := models.Vote{UpVote: strings.Count(Subject.UpVotes, "#"), DownVote: strings.Count(Subject.DownVotes, "#")}
	result, err := json.Marshal(Votes)
	if err != nil {
		w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	w.Write(result)
}
