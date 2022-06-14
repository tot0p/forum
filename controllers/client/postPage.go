package client

import (
	"encoding/json"
	"fmt"
	"forum/models"
	"forum/tools/request"
	"html/template"
	"io/ioutil"
	"net/http"
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

//Function to get the Username of the Actual owner for the current post
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

//Function to create a post page
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
	if r.Method == "POST" {
		Comment := models.Comment{}
		Comment.Parent = p.Post.Id
		Comment.Content = r.PostFormValue("content")
		err = request.PostComment(Comment, cookie.Value)
		if err != nil {
			w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
			return
		}
		p.Post, err = request.GetPostById(m["id"])
		if err != nil || p.Post.Id == "" {
			w.Write([]byte("{\"err\":\"not exist\"}"))
			return
		}
	}
	p.Comments, err = request.GetCommentsByPostId(p.Post.Id)
	if err != nil {
		fmt.Println("oh non err dans post page")
	}
	w.WriteHeader(http.StatusOK)
	tmpl := template.Must(template.ParseFiles(CurrentFolder + p.Path))
	tmpl.Execute(w, p)
}
