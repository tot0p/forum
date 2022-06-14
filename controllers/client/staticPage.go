package client

import (
	"html/template"
	"net/http"
)

type StaticPage struct {
	Path string
}

//Function to serve the page who don't need template
func (s *StaticPage) ServeHTTP(w http.ResponseWriter, r *http.Request, m map[string]string) {
	w.WriteHeader(http.StatusOK)
	tmpl := template.Must(template.ParseFiles(CurrentFolder + s.Path))
	tmpl.Execute(w, s)
}
