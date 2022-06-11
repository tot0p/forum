package client

import (
	"encoding/base64"
	"forum/models"
	"forum/tools/request"
	"net/http"
	"text/template"
)

type PostPage struct {
	Path                 string
	Post                 models.Post
	User                 models.User
	Connected            bool
	ProfilePictureBase64 string
}

func (p *PostPage) ServeHTTP(w http.ResponseWriter, r *http.Request, m map[string]string) {
	cookie, err := r.Cookie("SID")
	if err != nil {
		p.Connected = false
	} else {
		p.User, err = request.GetMe(cookie.Value)
		p.Connected = err == nil
		p.ProfilePictureBase64 = base64.StdEncoding.EncodeToString(p.User.ProfilePicture)
	}
	p.Post, err = request.GetPostById(m["id"])
	if err != nil || p.Post.Id == "" {
		w.Write([]byte("{\"err\":\"not exist\"}"))
		return
	}
	w.WriteHeader(http.StatusOK)
	tmpl := template.Must(template.ParseFiles(CurrentFolder + p.Path))
	tmpl.Execute(w, p)
}
