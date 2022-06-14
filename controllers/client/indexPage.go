package client

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"forum/models"
	"forum/tools/request"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

var CurrentFolder = "src/html/"

type Page struct {
	Path                 string
	User                 models.User
	ProfilePictureBase64 string
	LatestSubjects       []models.Subject
	LatestPosts          []models.Post
	Usernames            map[string]string
	Connected            bool
	Stats                models.AllCount
}

//Function to get the Username of the actual session
func (i Page) GetOwnerUsername(UUID string) string {
	if _, ok := i.Usernames[UUID]; ok {
		return i.Usernames[UUID]
	}
	resp, err := http.Get(os.Getenv("url_api") + fmt.Sprintf("username/%s", UUID))
	if err != nil {
		return ""
	}
	reqBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	var jsonReqBody map[string]string
	err = json.Unmarshal(reqBody, &jsonReqBody)
	if err != nil {
		return ""
	}
	i.Usernames[UUID] = jsonReqBody["username"]
	return jsonReqBody["username"]
}

//Function to create the owner session page
func (i *Page) ServeHTTP(w http.ResponseWriter, r *http.Request, m map[string]string) {
	i.Usernames = make(map[string]string)
	cookie, err := r.Cookie("SID")
	if err != nil {
		i.Connected = false
	} else {
		i.User, err = request.GetMe(cookie.Value)
		i.Connected = err == nil
		i.ProfilePictureBase64 = base64.StdEncoding.EncodeToString(i.User.ProfilePicture)
	}
	resp, err := http.Get(os.Getenv("url_api") + "subject/GetLastSubjectUpdate/15")
	if err != nil {
		i.LatestSubjects = nil
	}
	rep, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		i.LatestSubjects = nil
	}
	err = json.Unmarshal(rep, &i.LatestSubjects)
	if err != nil {
		i.LatestSubjects = nil
	}
	resp, err = http.Get(os.Getenv("url_api") + "post/GetLastPost/20")
	if err != nil {
		i.LatestSubjects = nil
	}
	rep, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		i.LatestSubjects = nil
	}
	err = json.Unmarshal(rep, &i.LatestPosts)
	if err != nil {
		i.LatestSubjects = nil
	}
	resp, err = http.Get(os.Getenv("url_api") + "count")
	if err != nil {
		i.LatestSubjects = nil
	}
	rep, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		i.LatestSubjects = nil
	}
	err = json.Unmarshal(rep, &i.Stats)
	if err != nil {
		i.LatestSubjects = nil
	}

	w.WriteHeader(http.StatusOK)
	tmpl := template.Must(template.ParseFiles(CurrentFolder + i.Path))
	tmpl.Execute(w, i)
}
