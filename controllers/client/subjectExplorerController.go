package client

import (
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

type SubjectExplorerPage struct {
	Path      string
	User      models.User
	Connected bool
	Usernames map[string]string
	Subjects  []models.Subject
}

func (i SubjectExplorerPage) GetUpVoteDownVoteSubject(subj models.Subject) models.Vote {
	return subj.GetVote()
}

//Function to get the username of the actual session
func (i SubjectExplorerPage) GetOwnerUsername(UUID string) string {
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

//Method to create a page to explore the subjects
func (p *SubjectExplorerPage) ServeHTTP(w http.ResponseWriter, r *http.Request, m map[string]string) {
	p.Usernames = make(map[string]string)
	cookie, err := r.Cookie("SID")
	if err != nil {
		p.Connected = false
	} else {
		p.User, err = request.GetMe(cookie.Value)
		p.Connected = err == nil
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
	resp, err := http.Get(os.Getenv("url_api") + "subject/GetNBSubject/15")
	if err != nil {
		p.Subjects = nil
	}
	rep, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		p.Subjects = nil
	}
	err = json.Unmarshal(rep, &p.Subjects)
	if err != nil {
		p.Subjects = nil
	}
	w.WriteHeader(http.StatusOK)
	tmpl := template.Must(template.ParseFiles(CurrentFolder + p.Path))
	tmpl.Execute(w, p)
}
