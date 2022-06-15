package api

import (
	"io/ioutil"
	"net/http"
)

//write Doc in html
func Doc(w http.ResponseWriter, r *http.Request) {
	fileBytes, err := ioutil.ReadFile("controllers/api/ReadMe.md")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/markdown")
	w.Write(fileBytes)
	return
}
