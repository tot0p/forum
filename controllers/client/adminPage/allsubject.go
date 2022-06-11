package adminpage

import (
	"forum/models"
	"forum/tools/request"
	"net/http"
	"text/template"
)

type AllSubject struct {
	Path        string
	AllSubjects []models.Subject
	Connected   bool
	User        models.User
}

func (a *AllSubject) ServeHTTP(w http.ResponseWriter, r *http.Request, m map[string]string) {
	a.AllSubjects = []models.Subject{}
	cookie, err := r.Cookie("SID")
	if err != nil {
		a.Connected = false
	} else {
		a.User, err = request.GetMe(cookie.Value)
		a.Connected = err == nil
	}
	if !a.Connected {
		w.Write([]byte("{\"err\":\"403\",\"msg\":\"Session invalide\"}"))
		return
	}
	if a.User.Role != "admin" {
		w.Write([]byte("{\"err\":\"403\",\"msg\":\"Forbidden\"}"))
		return
	}
	if r.Method == "POST" {
		err := request.DeleteSubjectById(r.PostFormValue("id"), cookie.Value)
		if err != nil {
			w.Write([]byte("{\"err\":\"403\",\"msg\":\"" + err.Error() + "\"}"))
			return
		}
	}
	a.AllSubjects, err = request.GetAllSubject()
	if err != nil {
		w.Write([]byte("{\"err\":\"403\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	tmpl := template.Must(template.ParseFiles(CurrentFolder + a.Path))
	tmpl.Execute(w, a)
}
