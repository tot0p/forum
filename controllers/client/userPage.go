package client

import (
	"encoding/base64"
	"forum/models"
	"forum/tools/request"
	"forum/tools/riot"
	modelsriot "forum/tools/riot/modelsRiot"
	"net/http"
	"text/template"
)

type UserPage struct {
	Path                 string
	User                 models.User
	ProfilePictureBase64 string
	UserRiot             []modelsriot.Rank
	SummonerName         string
	Connected            bool
	UserP                models.User
}

func (p *UserPage) ServeHTTP(w http.ResponseWriter, r *http.Request, m map[string]string) {
	cookie, err := r.Cookie("SID")
	if err != nil {
		p.Connected = false
	} else {
		p.User, err = request.GetMe(cookie.Value)
		p.Connected = err == nil
		p.ProfilePictureBase64 = base64.StdEncoding.EncodeToString(p.User.ProfilePicture)
	}

	p.UserRiot = riot.API.GetRankById(p.User.RiotId)
	if len(p.UserRiot) < 1 {
		p.SummonerName = ""
	} else {
		p.SummonerName = p.UserRiot[0].SummonerName
	}

	p.UserP, err = request.GetUserByName(m["name"])
	if err != nil {
		w.Write([]byte("{\"user\":\"not exist\""))
		return
	}
	w.WriteHeader(http.StatusOK)
	tmpl := template.Must(template.ParseFiles(CurrentFolder + p.Path))
	tmpl.Execute(w, p)
}
