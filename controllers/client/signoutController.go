package client

import (
	"forum/models"
	"forum/tools/request"
	"net/http"
)

type SignOutPage struct {
	Connected bool
	User      models.User
}

//Function to signout on the api
func (s *SignOutPage) ServeHTTP(w http.ResponseWriter, r *http.Request, m map[string]string) {
	cookie, err := r.Cookie("SID")
	if err != nil {
		s.Connected = false
	} else {
		s.User, err = request.GetMe(cookie.Value)
		s.Connected = err == nil
	}
	if s.Connected {
		request.SignOutUser(cookie.Value)
	}
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}
