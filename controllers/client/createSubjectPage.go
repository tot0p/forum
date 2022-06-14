package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"forum/tools/authorization"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type CreateSubjectPage struct {
	Path  string
	Error string
}

//Function to create page for a subject
func (p *CreateSubjectPage) ServeHTTP(w http.ResponseWriter, r *http.Request, m map[string]string) {
	if r.Method == "POST" {
		err := r.ParseMultipartForm(5 << 20) // allocate 5mb of ram for the form
		if err != nil {
			log.Fatal(err)
		}
		Data := make(map[string]interface{})
		var ppFileBytes []byte = []byte(" ")
		ppFile, _, err := r.FormFile("image")
		if err == nil {
			ppFileBytes, err = ioutil.ReadAll(ppFile)
			if err != nil {
				log.Fatal(err)
			}
			Data["image"] = fmt.Sprintf("%02x", ppFileBytes)
		} else {
			Data["image"] = ""
		}
		Data["title"] = r.PostFormValue("title")
		Data["description"] = r.PostFormValue("description")
		Data["nsfw"] = r.PostFormValue("nsfw")
		Data["tags"] = r.PostFormValue("tags")
		if Data["nsfw"] == "on" {
			Data["nsfw"] = 1
		} else {
			Data["nsfw"] = 0
		}
		Data["tags"] = strings.Split(fmt.Sprintf("%v", Data["tags"]), "#")
		url := os.Getenv("url_api") + "subject"
		jsonData, err := json.Marshal(Data)
		if err != nil {
			log.Fatal(err)
		}
		cookie, err := r.Cookie("SID")
		if err != nil {
			fmt.Println(err)
		}
		client := &http.Client{}
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
		if err != nil {
			log.Fatal(err)
		}
		req.Header.Set("content-Type", "application/json")
		authorization.SetAuthorizationBearer(cookie.Value, req)
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		_, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
	}
	w.WriteHeader(http.StatusOK)
	tmpl := template.Must(template.ParseFiles(CurrentFolder + p.Path))
	tmpl.Execute(w, p)
}
