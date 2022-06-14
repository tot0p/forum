package datapage

import (
	"encoding/json"
	"fmt"
	"forum/models"
	"forum/tools/authorization"
	"forum/tools/session"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var CurrentFolder = "src/html/data/"

type UserDataPage struct {
	Path string
	User []models.User
}

//Funtion to create a data page
func (i *UserDataPage) ServeHTTP(w http.ResponseWriter, r *http.Request, m map[string]string) {
	i.User = []models.User{}
	url := os.Getenv("url_api") + "users"
	client := &http.Client{}
	sess := session.GlobalSessions.SessionStart(w, r)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}
	authorization.SetAuthorizationBearer(sess.SessionID(), req)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	reqBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"error\":\"500\"}"))
		return
	}
	if string(reqBody) == "{\"err\":\"403\",\"msg\":\"Forbidden\"}" {
		w.Write(reqBody)
		return
	}
	err = json.Unmarshal(reqBody, &i.User)
	if err != nil {
		log.Fatal(err)
	}
	var cam = map[string]int{}
	for _, elem := range i.User {
		cam[elem.Genre]++
	}
	w.WriteHeader(http.StatusOK)
	rep, err := CamenbertGenerator(cam, "Genre", "chart_div")
	if err != nil {
		fmt.Println(err)
	}
	cam = map[string]int{}
	for _, elem := range i.User {
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
	tmpl := template.Must(template.ParseFiles(CurrentFolder + i.Path))
	tmpl.Execute(w, i)
	w.Write([]byte(rep))
	w.Write([]byte(rep2))
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
