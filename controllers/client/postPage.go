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
	"net/url"
	"os"
)

type PostPage struct {
	Path      string
	Post      models.Post
	Comments  []models.Comment
	User      models.User
	Usernames map[string]string
	Connected bool
}

//Function to upvote or downvote a post
func (i PostPage) GetUpVoteDownVotePost(post models.Post) models.Vote {
	return post.GetVote()
}

//Function to upvote or downvote a subject
func (i PostPage) GetUpVoteDownVoteComment(com models.Comment) models.Vote {
	return com.GetVote()
}

//Function to get the username of the actual owner for the current post
func (i PostPage) GetOwnerUsername(UUID string) string {
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

//Method to create a post page
func (p *PostPage) ServeHTTP(w http.ResponseWriter, r *http.Request, m map[string]string) {
	p.Usernames = make(map[string]string)
	cookie, err := r.Cookie("SID")
	if err != nil {
		p.Connected = false
	} else {
		p.User, err = request.GetMe(cookie.Value)
		p.Connected = err == nil
	}
	p.Post, err = request.GetPostById(m["id"])
	if err != nil || p.Post.Id == "" {
		w.Write([]byte("{\"err\":\"not exist\"}"))
		return
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
					if params["type"] == "comment" {
						request.LikeComment(params["id"], cookie.Value, params["why"])
					} else if params["type"] == "post" {
						request.LikePost(params["id"], cookie.Value, params["why"])
					}
				}
			}
		}
		if _, ok := params["content"]; ok {
			Comment := models.Comment{}
			Comment.Parent = p.Post.Id
			Comment.Content, err = url.QueryUnescape(params["content"])
			if err != nil {
				w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
				return
			}
			err = request.PostComment(Comment, cookie.Value)
			if err != nil {
				w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
				return
			}
		}
		p.Post, err = request.GetPostById(m["id"])
		if err != nil || p.Post.Id == "" {
			temp := Page404{Path: "404.html"}
			temp.ServeHTTP(w, r, m)
			return
		}
	}
	p.Comments, err = request.GetCommentsByPostId(p.Post.Id)
	if err != nil {
		p.Comments = []models.Comment{}
	}
	w.WriteHeader(http.StatusOK)
	tmpl := template.Must(template.ParseFiles(CurrentFolder + p.Path))
	tmpl.Execute(w, p)
}
