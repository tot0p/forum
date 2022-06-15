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
	"strings"
)

type PageSearch struct {
	Path              string
	GoodSearch        bool
	Connected         bool
	Subjects          []models.Subject
	Users             []models.User
	Posts             []models.Post
	User              models.User
	Usernames         map[string]string
	Search, Searchval string
}

func (i PageSearch) GetUpVoteDownVoteSubject(subj models.Subject) models.Vote {
	return subj.GetVote()
}

//Function to upvote or downvote a post
func (i PageSearch) GetUpVoteDownVotePost(post models.Post) models.Vote {
	return post.GetVote()
}

func (i PageSearch) GetOwnerUsername(UUID string) string {
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

func (p *PageSearch) ServeHTTP(w http.ResponseWriter, r *http.Request, m map[string]string) {
	p.Usernames = make(map[string]string)
	cookie, err := r.Cookie("SID")
	if err != nil {
		p.Connected = false
	} else {
		p.User, err = request.GetMe(cookie.Value)
		p.Connected = err == nil
	}
	p.GoodSearch = true
	if r.Method == "POST" {
		resp, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.Write([]byte("{\"msg\":\"error request\"}"))
			return
		}
		params := tools.PlaintTextToMap(resp)
		fmt.Println(string(resp))
		if p.Connected {
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
		if _, ok := params["search"]; ok {
			p.Search = params["search"]
			if _, ok := params["search-value"]; ok {
				p.Searchval = params["search-value"]
			} else {
				p.Searchval = ""
			}
			switch params["search"] {
			case "Subject":
				p.Subjects, err = request.GetSearchSubjects(p.Searchval)
				if err != nil {
					p.Subjects = nil
					p.GoodSearch = false
				}
			case "Post":
				p.Posts, err = request.GetSearchPosts(p.Searchval)
				if err != nil {
					p.Posts = nil
					p.GoodSearch = false
				}
			case "User":
				p.Users, err = request.GetSearchUser(p.Searchval)
				if err != nil {
					p.Users = nil
					p.GoodSearch = false
				}
			default:
				p.GoodSearch = false
			}
		}
	} else if r.Method == "GET" {
		params := map[string]string{}
		for k, elem := range r.URL.Query() {
			params[k] = strings.Join(elem, ",")
		}
		if _, ok := params["search"]; ok {
			p.Search = params["search"]
			if _, ok := params["search-value"]; ok {
				p.Searchval = params["search-value"]
			} else {
				p.Searchval = ""
			}
			switch params["search"] {
			case "Subject":
				p.Subjects, err = request.GetSearchSubjects(p.Searchval)
				if err != nil {
					p.Subjects = nil
					p.GoodSearch = false
				}
			case "Post":
				p.Posts, err = request.GetSearchPosts(p.Searchval)
				if err != nil {
					p.Posts = nil
					p.GoodSearch = false
				}
			case "User":
				p.Users, err = request.GetSearchUser(p.Searchval)
				if err != nil {
					p.Users = nil
					p.GoodSearch = false
				}
			default:
				p.GoodSearch = false
			}
		}
	}
	tmpl := template.Must(template.ParseFiles(CurrentFolder + p.Path))
	tmpl.Execute(w, p)
}
