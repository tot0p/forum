package client

import (
	"forum/models"
	"html/template"
	"net/http"
)

type UpdateCommentPage struct {
	Path string
	User models.User
}

//Function to create a page on which you can update a comment
func (p *UpdateCommentPage) ServeHTTP(w http.ResponseWriter, r *http.Request, m map[string]string) {
	w.WriteHeader(http.StatusOK)
	tmpl := template.Must(template.ParseFiles(CurrentFolder + p.Path))
	tmpl.Execute(w, p)
}
