package client

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"forum/models"
	"forum/tools/request"
	"io/ioutil"
	"net/http"
	"os"
	"text/template"
)

type SubjectPage struct {
	Path                 string
	Connected            bool
	User                 models.User
	ProfilePictureBase64 string
	Usernames            map[string]string
	Subject              models.Subject
	AllPost              []models.Post
}

func (p SubjectPage) GetOwnerUsername(UUID string) string {
	if _, ok := p.Usernames[UUID]; ok {
		return p.Usernames[UUID]
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
	p.Usernames[UUID] = jsonReqBody["username"]
	return jsonReqBody["username"]
}

func (p *SubjectPage) ServeHTTP(w http.ResponseWriter, r *http.Request, m map[string]string) {
	p.Usernames = make(map[string]string)
	cookie, err := r.Cookie("SID")
	if err != nil {
		p.Connected = false
	} else {
		p.User, err = request.GetMe(cookie.Value)
		p.Connected = err == nil
		p.ProfilePictureBase64 = base64.StdEncoding.EncodeToString(p.User.ProfilePicture)
	}
	p.Subject, err = request.GetSubjectById(m["id"])
	if err != nil || p.Subject.Id == "" {
		w.Write([]byte("{\"err\":\"not exist\"}"))
		return
	}
	for _, elem := range p.Subject.ConvertAllPosts() {
		temp, err := request.GetPostById(elem)
		if err != nil {
			fmt.Println(err)
		} else {
			p.AllPost = append(p.AllPost, temp)
		}
	}
	w.WriteHeader(http.StatusOK)
	tmpl := template.Must(template.ParseFiles(CurrentFolder + p.Path))
	tmpl.Execute(w, p)
}
