package client

import (
	"html/template"
	"net/http"
)

type Page404 struct {
	Path string
	Url  string
}

//Function for the 404 error page
func (p *Page404) ServeHTTP(w http.ResponseWriter, r *http.Request, m map[string]string) {
	w.WriteHeader(http.StatusNotFound)
	tmpl := template.Must(template.ParseFiles(CurrentFolder + p.Path))
	p.Url = r.URL.Path
	tmpl.Execute(w, p)
}
