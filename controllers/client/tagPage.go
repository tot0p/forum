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
	"strings"
)

type TagPage struct {
	Path                 string
	Subjects             []models.Subject
	Posts                []models.Post
	Connected            bool
	User                 models.User
	ProfilePictureBase64 string
	Usernames            map[string]string
	Typ, Tag             string
}

//Function to upvote or downvote a post
func (i TagPage) GetUpVoteDownVotePost(post models.Post) models.Vote {
	return post.GetVote()
}

//Function to upvote or downvote a subject
func (i TagPage) GetUpVoteDownVoteSubject(subject models.Subject) models.Vote {
	return subject.GetVote()
}

//Method to get the username of the actual session
func (p TagPage) GetOwnerUsername(UUID string) string {
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

func (p *TagPage) ServeHTTP(w http.ResponseWriter, r *http.Request, m map[string]string) {
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
	p.Typ = m["type"]
	p.Tag = m["tag"]
	if strings.ToLower(m["type"]) == "subject" {
		p.Subjects, err = request.GetSubjectsByTag(m["tag"])
		if err != nil {
			p.Subjects = nil
		}
	} else if strings.ToLower(m["type"]) == "post" {
		p.Posts, err = request.GetPostsByTag(m["tag"])
		if err != nil {
			p.Posts = nil
		}
	}
	w.WriteHeader(http.StatusOK)
	tmpl := template.Must(template.ParseFiles(CurrentFolder + p.Path))
	tmpl.Execute(w, p)
}
