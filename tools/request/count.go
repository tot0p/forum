package request

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

//Function to get the number of users created from the api
func GetCountUser() int {
	resp, err := http.Get(os.Getenv("url_api") + "count/user")
	if err != nil {
		return -1
	}
	rep, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return -1
	}
	var m map[string]interface{}
	err = json.Unmarshal(rep, &m)
	if err != nil {
		return -1
	}
	return int(m["Nb"].(float64))
}

//Function to get the number of sessions running from the api
func GetCountSession() int {
	resp, err := http.Get(os.Getenv("url_api") + "count/session")
	if err != nil {
		return -1
	}
	rep, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return -1
	}
	var m map[string]interface{}
	err = json.Unmarshal(rep, &m)
	if err != nil {
		return -1
	}
	return int(m["Nb"].(float64))
}

//Function to get the number of subjects created from the api
func GetCountSubject() int {
	resp, err := http.Get(os.Getenv("url_api") + "count/subject")
	if err != nil {
		return -1
	}
	rep, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return -1
	}
	var m map[string]interface{}
	err = json.Unmarshal(rep, &m)
	if err != nil {
		return -1
	}
	return int(m["Nb"].(float64))
}

//Function to get the number of posts created from the api
func GetCountPost() int {
	resp, err := http.Get(os.Getenv("url_api") + "count/post")
	if err != nil {
		return -1
	}
	rep, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return -1
	}
	var m map[string]interface{}
	err = json.Unmarshal(rep, &m)
	if err != nil {
		return -1
	}
	return int(m["Nb"].(float64))
}
