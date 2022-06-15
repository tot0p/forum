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

//Method to create a page for the team
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
	user, err = request.GetUserByName("Aless")
	if err != nil {
		w.Write([]byte("error"))
		return
	}
	t.AllTeam = append(t.AllTeam, user)
	w.WriteHeader(http.StatusOK)
	tmpl := template.Must(template.ParseFiles(CurrentFolder + t.Path))
	tmpl.Execute(w, t)
}

//Method to get the summoner name of an user
func (t *TeamPage) GetSummonerName(user models.User) string {
	summonerName := riot.API.GetUserById(user.RiotId).Name
	return summonerName
}

//Method to get the rank of an user
func (t *TeamPage) GetRank(user models.User) []modelsriot.Rank {
	rank := riot.API.GetRankById(user.RiotId)
	var inOrder []modelsriot.Rank
	Solo := 9
	Flex := 9
	for i, elem := range rank {
		if elem.QueueType == "RANKED_SOLO_5x5" {
			Solo = i
		}
		if elem.QueueType == "RANKED_FLEX_SR" {
			Flex = i
		}
	}
	var void modelsriot.Rank
	if Solo != 9 {
		inOrder = append(inOrder, rank[Solo])
	} else {
		inOrder = append(inOrder, void)
	}
	if Flex != 9 {
		inOrder = append(inOrder, rank[Flex])
	} else {
		inOrder = append(inOrder, void)
	}
	return inOrder
}
