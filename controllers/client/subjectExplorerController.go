package client

import (
	"encoding/base64"
	"forum/models"
	"forum/tools/request"
	"html/template"
	"net/http"
)

type SubjectExplorerPage struct {
	Path                 string
	User                 models.User
	ProfilePictureBase64 string
	Connected            bool
}

func (p *SubjectExplorerPage) ServeHTTP(w http.ResponseWriter, r *http.Request, m map[string]string) {
	cookie, err := r.Cookie("SID")
	if err != nil {
		p.Connected = false
	} else {
		p.User, err = request.GetMe(cookie.Value)
		p.Connected = err == nil
		p.ProfilePictureBase64 = base64.StdEncoding.EncodeToString(p.User.ProfilePicture)
	}
	w.WriteHeader(http.StatusOK)
	tmpl := template.Must(template.ParseFiles(CurrentFolder + p.Path))
	tmpl.Execute(w, p)
}
