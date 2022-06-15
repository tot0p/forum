package client

import (
	"forum/models"
	"forum/tools/request"
	"forum/tools/riot"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

type UpdateProfilePage struct {
	Path         string
	User         models.User
	SummonerName string
	Connected    bool
}

//Method to create a page on which you can update a profile
func (p *UpdateProfilePage) ServeHTTP(w http.ResponseWriter, r *http.Request, m map[string]string) {
	var err error
	p.User = models.User{}
	cookie, err := r.Cookie("SID")
	if err != nil {
		p.Connected = false
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	p.User, err = request.GetMe(cookie.Value)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	} else {
		p.User, err = request.GetMe(cookie.Value)
		p.Connected = err == nil
	}
	p.SummonerName = riot.API.GetUserById(p.User.RiotId).Name
	if r.Method == "POST" {
		err := r.ParseMultipartForm(5 << 20) // allocate 5mb of ram for the form
		if err != nil {
			log.Fatal(err)
		}
		p.User.Password = r.PostFormValue("password")
		p.User.Username = r.PostFormValue("username")
		p.User.Email = r.PostFormValue("email")
		p.User.FirstName = r.PostFormValue("firstname")
		p.User.LastName = r.PostFormValue("lastname")
		p.User.BirthDate = r.PostFormValue("birthdate")
		p.User.Genre = r.PostFormValue("genre")
		p.User.RiotId = riot.API.GetUserByName(r.PostFormValue("riotid")).Id
		p.User.Bio = r.PostFormValue("bio")
		ppFile, _, err := r.FormFile("profilepicture")
		if err == nil {
			ppFileBytes, err := ioutil.ReadAll(ppFile)
			if err != nil {
				log.Fatal(err)
			}
			p.User.ProfilePicture = ppFileBytes
		} else {
			p.User.ProfilePicture = []byte("")
		}
		err = request.PutUser(p.User, cookie.Value)
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
