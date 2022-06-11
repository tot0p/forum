package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"forum/models"
	"forum/tools/authorization"
	"forum/tools/request"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"
)

type CreatePostPage struct {
	Path       string
	AllSubject []models.Subject
}

func (p *CreatePostPage) ServeHTTP(w http.ResponseWriter, r *http.Request, m map[string]string) {
	p.AllSubject, _ = request.GetAllSubject()
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
		Data["parent"] = r.PostFormValue("parent")
		if Data["nsfw"] == "on" {
			Data["nsfw"] = 1
		} else {
			Data["nsfw"] = 0
		}
		Data["tags"] = strings.Split(fmt.Sprintf("%v", Data["tags"]), "#")
		url := os.Getenv("url_api")
		url += "post"
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
