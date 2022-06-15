package client

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"forum/models"
	"forum/tools"
	"forum/tools/request"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
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

//Function to upvote or downvote a post
func (i SubjectPage) GetUpVoteDownVotePost(post models.Post) models.Vote {
	return post.GetVote()
}

//Function to upvote or downvote a subject
func (i SubjectPage) GetUpVoteDownVoteSubject(subject models.Subject) models.Vote {
	return subject.GetVote()
}

//Method to get the username of the actual session
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

//Method to create a page for a subject
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
	if r.Method == "POST" && p.Connected {
		resp, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.Write([]byte("{\"msg\":\"error request\"}"))
			return
		}
		params := tools.PlaintTextToMap(resp)
		if typ, ok := params["type"]; ok && typ != "" {
			if id, ok := params["id"]; ok && id != "" {
				if why, ok := params["why"]; ok && why != "" {
					if params["type"] == "subject" {
						request.LikeSubject(params["id"], cookie.Value, params["why"])
					} else if params["type"] == "post" {
						request.LikePost(params["id"], cookie.Value, params["why"])
					}
				}
			}
		}
	}
	p.Subject, err = request.GetSubjectById(m["id"])
	if err != nil || p.Subject.Id == "" {
		temp := Page404{Path: "404.html"}
		temp.ServeHTTP(w, r, m)
		return
	}
	p.AllPost, _ = request.GetPostsBySubjectId(p.Subject.Id)
	w.WriteHeader(http.StatusOK)
	tmpl := template.Must(template.ParseFiles(CurrentFolder + p.Path))
	tmpl.Execute(w, p)
}
