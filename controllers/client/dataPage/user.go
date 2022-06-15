package datapage

import (
	"fmt"
	"forum/models"
	"forum/tools/request"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

var CurrentFolder = "src/html/data/"

type UserDataPage struct {
	Path      string
	Users     []models.User
	User      models.User
	Connected bool
}

//Funtion to create a data page
func (i *UserDataPage) ServeHTTP(w http.ResponseWriter, r *http.Request, m map[string]string) {
	cookie, err := r.Cookie("SID")
	if err != nil {
		i.Connected = false
	} else {
		i.User, err = request.GetMe(cookie.Value)
		i.Connected = err == nil
	}
	if !i.Connected || i.User.Role != "admin" {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	i.Users, err = request.GetAllUser(cookie.Value)
	if err != nil {
		w.Write([]byte("{\"msg\":\"internal error\"}"))
		return
	}
	//genre
	var cam = map[string]int{}
	for _, elem := range i.Users {
		cam[elem.Genre]++
	}
	w.WriteHeader(http.StatusOK)
	rep, err := CamenbertGenerator(cam, "Genre", "chart_div")
	if err != nil {
		fmt.Println(err)
	}
	//premium
	cam = map[string]int{}
	for _, elem := range i.Users {
		if elem.Premium >= 1 {
			cam["Premium"]++
		} else {
			cam["NonPremium"]++
		}
	}
	rep2, err := CamenbertGenerator(cam, "Premium", "zaoaz")
	if err != nil {
		log.Fatal(err)
	}
	//with riot id
	riotid := map[string]int{}
	for _, elem := range i.Users {
		if elem.RiotId != "" {
			riotid["RiotId"]++
		} else {
			riotid["With out Riot Id"]++
		}
	}
	rep3, err := CamenbertGenerator(riotid, "RiotId", "riotid")
	if err != nil {
		log.Fatal(err)
	}

	tmpl := template.Must(template.ParseFiles(CurrentFolder + i.Path))
	tmpl.Execute(w, i)
	w.Write([]byte(rep))
	w.Write([]byte(rep2))
	w.Write([]byte(rep3))
	// connection_journaliere
}

//Function to create a camenbert
func CamenbertGenerator(data map[string]int, title string, id string) (string, error) {
	res := "["
	first := true
	for i, f := range data {
		if first {
			res += fmt.Sprintf("[\"%s\",%d]", i, f)
			first = false
		} else {
			res += fmt.Sprintf(",[\"%s\",%d]", i, f)
		}
	}
	res += "]"
	body, err := ioutil.ReadFile("src/jsNoShare/data/camenbert.js")
	if err != nil {
		return "", err
	}
	res = fmt.Sprintf(string(body), id, id, res, title, id)
	return "<script>" + res + "</script>", nil

}
