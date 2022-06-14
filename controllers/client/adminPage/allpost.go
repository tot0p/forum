package adminpage

import (
	"forum/models"
	"forum/tools/request"
	"html/template"
	"net/http"
)

type AllPost struct {
	Path      string
	AllPosts  []models.Post
	Connected bool
	User      models.User
}

//Method linked to the page contain all the posts
func (a *AllPost) ServeHTTP(w http.ResponseWriter, r *http.Request, m map[string]string) {
	a.AllPosts = []models.Post{}
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
		err := request.DeletePostById(r.PostFormValue("id"), cookie.Value)
		if err != nil {
			w.Write([]byte("{\"err\":\"403\",\"msg\":\"" + err.Error() + "\"}"))
			return
		}
	}
	a.AllPosts, err = request.GetAllPost()
	if err != nil {
		w.Write([]byte("{\"err\":\"403\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	tmpl := template.Must(template.ParseFiles(CurrentFolder + a.Path))
	tmpl.Execute(w, a)
}
