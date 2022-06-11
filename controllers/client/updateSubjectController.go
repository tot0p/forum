package client

import (
	"encoding/base64"
	"fmt"
	"forum/models"
	"forum/tools/request"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

type UpdateSubjectPage struct {
	Path                 string
	User                 models.User
	Subject              models.Subject
	Connected            bool
	ProfilePictureBase64 string
}

func (p *UpdateSubjectPage) ServeHTTP(w http.ResponseWriter, r *http.Request, m map[string]string) {
	cookie, err := r.Cookie("SID")
	if err != nil {
		p.Connected = false
	} else {
		p.User, err = request.GetMe(cookie.Value)
		p.Connected = err == nil
		p.ProfilePictureBase64 = base64.StdEncoding.EncodeToString(p.User.ProfilePicture)
	}
	p.Subject, err = request.GetSubjectById(m["id"])
	if err != nil {
		fmt.Println(err.Error())
		w.Write([]byte("error"))
		return
	}
	if p.Subject.Owner != p.User.UUID && p.User.Role != "admin" {
		w.Write([]byte("{\"error\":\"forbiden access\"}"))
		return
	}
	if r.Method == "POST" {
		err := r.ParseMultipartForm(5 << 20) // allocate 5mb of ram for the form
		if err != nil {
			log.Fatal(err)
		}
		p.Subject.Title = r.PostFormValue("title")
		p.Subject.Description = r.PostFormValue("description")
		nsfw := r.PostFormValue("nsfw")
		p.Subject.Tags = r.PostFormValue("tags")
		if nsfw == "on" {
			p.Subject.NSFW = 1
		} else {
			p.Subject.NSFW = 0
		}
		ppFile, _, err := r.FormFile("image")
		if err == nil {
			ppFileBytes, err := ioutil.ReadAll(ppFile)
			if err != nil {
				log.Fatal(err)
			}
			p.Subject.Image = ppFileBytes
		} else {
			p.Subject.Image = []byte("")
		}
		err = request.PutSubject(p.Subject, cookie.Value)
		if err != nil {
			w.Write([]byte("{\"err\":\"500\",\"msg\":\"" + err.Error() + "\"}"))
			return
		}
	}
	w.WriteHeader(http.StatusOK)
	tmpl := template.Must(template.ParseFiles(CurrentFolder + p.Path))
	tmpl.Execute(w, p)
}
