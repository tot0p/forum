package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"forum/tools"
	"forum/tools/riot"
	"forum/tools/verif"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type RegisterPage struct {
	Path  string
	Error string
}

//Function to be register on the api
func (p *RegisterPage) ServeHTTP(w http.ResponseWriter, r *http.Request, m map[string]string) {
	if r.Method == "POST" {
		err := r.ParseMultipartForm(5 << 20) // allocate 5mb of ram for the form
		if err != nil {
			log.Fatal(err)
		}
		// admin
		// adminADMIN1234!
		PostDB := true
		var ppFileBytes []byte = []byte(" ")
		Data := make(map[string]string)
		ppFile, _, err := r.FormFile("profilepicture")
		if err == nil {
			ppFileBytes, err = ioutil.ReadAll(ppFile)
			if err != nil {
				ppFileBytes = tools.CreateImg()
			}
			Data["profilepicture"] = fmt.Sprintf("%02x", ppFileBytes)
		} else {
			Data["profilepicture"] = fmt.Sprintf("%02x", tools.CreateImg())
		}
		Data["bio"] = r.PostFormValue("bio")
		Data["birthdate"] = r.PostFormValue("birthdate")
		Data["email"] = r.PostFormValue("email")
		Data["firstname"] = r.PostFormValue("firstname")
		Data["genre"] = r.PostFormValue("genre")
		Data["lastname"] = r.PostFormValue("lastname")
		Data["username"] = r.PostFormValue("username")
		Data["password"] = r.PostFormValue("password")
		Data["riotid"] = r.PostFormValue("riotid")
		Data["Confirmation_Password"] = r.PostFormValue("Confirmation_Password")
		if Data["password"] != Data["Confirmation_Password"] {
			fmt.Println("Not Same Password")
			PostDB = false
		}
		if !verif.PasswordVerif(Data["password"]) {
			fmt.Println("ALED LE MOT DE PASSE EST PAS BON")
			PostDB = false
		}
		Data["riotid"] = riot.API.GetUserByName(Data["riotid"]).Id
		if PostDB {
			url := os.Getenv("url_api")
			url += "user"
			delete(Data, "Confirmation_Password")
			jsonData, err := json.Marshal(Data)
			if err != nil {
				log.Fatal(err)
			} else {
				resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
				if err != nil {
					log.Fatal(err)
				} else {
					_, err = ioutil.ReadAll(resp.Body)
					if err != nil {
						log.Fatal(err)
					} else {
						http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
					}
				}
			}

		}
	}
	w.WriteHeader(http.StatusOK)
	tmpl := template.Must(template.ParseFiles(CurrentFolder + p.Path))
	tmpl.Execute(w, p)
}
