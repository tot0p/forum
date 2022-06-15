package client

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"forum/models"
	"forum/tools/request"
	"forum/tools/riot"
	modelsriot "forum/tools/riot/modelsRiot"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

type UserPage struct {
	Path                 string
	UserOrigin           models.User
	ProfilePictureBase64 string
	UserRiot             []modelsriot.Rank
	SummonerName         string
	Connected            bool
	User                 models.User
	Posts                []models.Post
	Ranked               bool
	Subjects             []models.Subject
	Usernames            map[string]string
}

//Method to get the username of the actual owner for the current post
func (i UserPage) GetOwnerUsername(UUID string) string {
	if _, ok := i.Usernames[UUID]; ok {
		return i.Usernames[UUID]
	}
	resp, err := http.Get(os.Getenv("url_api") + fmt.Sprintf("username/%s", UUID))
	if err != nil {
		return ""
	}
	reqBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	var jsonReqBody map[string]string
	err = json.Unmarshal(reqBody, &jsonReqBody)
	if err != nil {
		return ""
	}
	i.Usernames[UUID] = jsonReqBody["username"]
	return jsonReqBody["username"]
}

//Function to create a user page
func (p *UserPage) ServeHTTP(w http.ResponseWriter, r *http.Request, m map[string]string) {
	p.Usernames = make(map[string]string)
	cookie, err := r.Cookie("SID")
	if err != nil {
		p.Connected = false
	} else {
		p.UserOrigin, err = request.GetMe(cookie.Value)
		p.Connected = err == nil
		p.ProfilePictureBase64 = base64.StdEncoding.EncodeToString(p.UserOrigin.ProfilePicture)
	}
	p.User, err = request.GetUserByName(m["name"])
	if err != nil {
		w.Write([]byte("{\"user\":\"not exist\"}"))
		return
	}
	p.UserRiot = riot.API.GetRankById(p.User.RiotId)
	if len(p.UserRiot) < 1 {
		p.SummonerName = ""
	} else {
		p.SummonerName = p.UserRiot[0].SummonerName
	}
	p.Posts, _ = request.GetPostsByUserId(p.User.UUID)
	if err != nil {
		fmt.Println("ds")
	}
	p.Subjects, _ = request.GetSubjectByUserId(p.User.UUID)
	p.Ranked = !(len(p.UserRiot) == 0)
	w.WriteHeader(http.StatusOK)
	tmpl := template.Must(template.ParseFiles(CurrentFolder + p.Path))
	tmpl.Execute(w, p)
}
