package client

import (
	"forum/models"
	"forum/tools/request"
	"forum/tools/riot"
	modelsriot "forum/tools/riot/modelsRiot"
	"html/template"
	"net/http"
)

type ProfilePage struct {
	Path         string
	User         models.User
	UserRiot     []modelsriot.Rank
	SummonerName string
	Connected    bool
	Ranked       bool
}

//Function to create a profile page
func (p *ProfilePage) ServeHTTP(w http.ResponseWriter, r *http.Request, m map[string]string) {
	cookie, err := r.Cookie("SID")
	if err != nil {
		p.Connected = false
	} else {
		p.User, err = request.GetMe(cookie.Value)
		p.Connected = err == nil
	}
	if !p.Connected {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	}
	p.UserRiot = riot.API.GetRankById(p.User.RiotId)
	if len(p.UserRiot) < 1 {
		p.SummonerName = ""
	} else {
		p.SummonerName = p.UserRiot[0].SummonerName
	}
	p.Ranked = !(len(p.UserRiot) == 0)
	w.WriteHeader(http.StatusOK)
	tmpl := template.Must(template.ParseFiles(CurrentFolder + p.Path))
	tmpl.Execute(w, p)
}
