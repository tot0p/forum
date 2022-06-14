package client

import (
	"forum/models"
	"forum/tools/request"
	"forum/tools/riot"
	modelsriot "forum/tools/riot/modelsRiot"
	"html/template"
	"net/http"
)

type TeamPage struct {
	Path     string
	AllTeam  []models.User
	UserRiot []modelsriot.Rank
}

func (t *TeamPage) ServeHTTP(w http.ResponseWriter, r *http.Request, m map[string]string) {
	var err error

	user, err := request.GetUserByName("Axou")
	if err != nil {
		w.Write([]byte("error"))
		return
	}
	t.AllTeam = append(t.AllTeam, user)

	user, err = request.GetUserByName("Sambre")
	if err != nil {
		w.Write([]byte("error"))
		return
	}
	t.AllTeam = append(t.AllTeam, user)

	user, err = request.GetUserByName("Tot0p")
	if err != nil {
		w.Write([]byte("error"))
		return
	}
	t.AllTeam = append(t.AllTeam, user)

	user, err = request.GetUserByName("mkarten")
	if err != nil {
		w.Write([]byte("error"))
		return
	}
	t.AllTeam = append(t.AllTeam, user)

	w.WriteHeader(http.StatusOK)
	tmpl := template.Must(template.ParseFiles(CurrentFolder + t.Path))
	tmpl.Execute(w, t)
}

func (t *TeamPage) GetSummonerName(user models.User) string {
	summonerName := riot.API.GetUserById(user.RiotId).Name
	return summonerName
}

func (t *TeamPage) GetRank(user models.User) []modelsriot.Rank {
	rank := riot.API.GetRankById(user.RiotId)
	var inOrder []modelsriot.Rank
	tempo0 := 9
	tempo1 := 9
	for i, elem := range rank {
		if elem.QueueType == "RANKED_SOLO_5x5" {
			tempo0 = i
		}
		if elem.QueueType == "RANKED_FLEX_SR" {
			tempo1 = i
		}
	}
	if tempo0 == 0 || tempo0 == 1 {
		inOrder = append(inOrder, rank[tempo0])
	}
	if tempo1 == 0 || tempo1 == 1 {
		inOrder = append(inOrder, rank[tempo1])
	}
	return inOrder
}
