package client

import (
	"encoding/base64"
	"forum/models"
	"forum/tools/request"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type UpdatePostPage struct {
	Path                 string
	User                 models.User
	Connected            bool
	ProfilePictureBase64 string
	Post                 models.Post
}

//Method to create a page on which you can update a post
func (p *UpdatePostPage) ServeHTTP(w http.ResponseWriter, r *http.Request, m map[string]string) {
	cookie, err := r.Cookie("SID")
	if err != nil {
		p.Connected = false
	} else {
		p.User, err = request.GetMe(cookie.Value)
		p.Connected = err == nil
		p.ProfilePictureBase64 = base64.StdEncoding.EncodeToString(p.User.ProfilePicture)
	}
	if !p.Connected {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	p.Post, err = request.GetPostById(m["id"])
	if err != nil {
		w.Write([]byte("error"))
		return
	}
	if p.Post.Owner != p.User.UUID && p.User.Role != "admin" && p.Post.Id == "" {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	if r.Method == "POST" && p.Connected {
		err := r.ParseMultipartForm(5 << 20) // allocate 5mb of ram for the form
		if err != nil {
			log.Fatal(err)
		}
		p.Post.Title = r.PostFormValue("title")
		p.Post.Description = r.PostFormValue("description")
		nsfw := r.PostFormValue("nsfw")
		p.Post.Tags = strings.Join(ParseTag(r.PostFormValue("tags")), "#")
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
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	}
	w.WriteHeader(http.StatusOK)
	tmpl := template.Must(template.ParseFiles(CurrentFolder + p.Path))
	tmpl.Execute(w, p)
}
