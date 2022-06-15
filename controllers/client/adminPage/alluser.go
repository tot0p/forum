package adminpage

import (
	"forum/models"
	"forum/tools/request"
	"html/template"
	"net/http"
)

var CurrentFolder = "src/html/admin/"

type AllUser struct {
	Path      string
	AllUsers  []models.User
	Connected bool
	User      models.User
}

//Method linked to the page containing all the Users
func (a *AllUser) ServeHTTP(w http.ResponseWriter, r *http.Request, m map[string]string) {
	a.AllUsers = []models.User{}
	cookie, err := r.Cookie("SID")
	if err != nil {
		a.Connected = false
	} else {
		a.User, err = request.GetMe(cookie.Value)
		a.Connected = err == nil
	}
	if !a.Connected {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	if a.User.Role != "admin" {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	if r.Method == "POST" {
		err := request.DeleteUserById(r.PostFormValue("id"), cookie.Value)
		if err != nil {
			w.Write([]byte("{\"err\":\"403\",\"msg\":\"" + err.Error() + "\"}"))
			return
		}
	}
	a.AllUsers, err = request.GetAllUser(cookie.Value)
	if err != nil {
		w.Write([]byte("{\"err\":\"403\",\"msg\":\"" + err.Error() + "\"}"))
		return
	}
	tmpl := template.Must(template.ParseFiles(CurrentFolder + a.Path))
	tmpl.Execute(w, a)
}
