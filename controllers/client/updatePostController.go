package client

import (
	"encoding/base64"
	"forum/models"
	"forum/tools/request"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

type UpdatePostPage struct {
	Path                 string
	User                 models.User
	Connected            bool
	ProfilePictureBase64 string
	Post                 models.Post
}

func (p *UpdatePostPage) ServeHTTP(w http.ResponseWriter, r *http.Request, m map[string]string) {
	cookie, err := r.Cookie("SID")
	if err != nil {
		p.Connected = false
	} else {
		p.User, err = request.GetMe(cookie.Value)
		p.Connected = err == nil
		p.ProfilePictureBase64 = base64.StdEncoding.EncodeToString(p.User.ProfilePicture)
	}
	p.Post, err = request.GetPostById(m["id"])
	if err != nil {
		w.Write([]byte("error"))
		return
	}
	if p.Post.Owner != p.User.UUID && p.User.Role != "admin" {
		w.Write([]byte("{\"error\":\"forbiden access\"}"))
		return
	}
	if r.Method == "POST" {
		err := r.ParseMultipartForm(5 << 20) // allocate 5mb of ram for the form
		if err != nil {
			log.Fatal(err)
		}
		p.Post.Title = r.PostFormValue("title")
		p.Post.Description = r.PostFormValue("description")
		nsfw := r.PostFormValue("nsfw")
		p.Post.Tags = r.PostFormValue("tags")
		if nsfw == "on" {
			p.Post.NSFW = 1
		} else {
			p.Post.NSFW = 0
		}
		ppFile, _, err := r.FormFile("image")
		if err == nil {
			ppFileBytes, err := ioutil.ReadAll(ppFile)
			if err != nil {
				log.Fatal(err)
			}
			p.Post.Image = ppFileBytes
		} else {
			p.Post.Image = []byte("")
		}
		err = request.PutPost(p.Post, cookie.Value)
		if err != nil {
			w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
			return
		}
	}
	w.WriteHeader(http.StatusOK)
	tmpl := template.Must(template.ParseFiles(CurrentFolder + p.Path))
	tmpl.Execute(w, p)
}
