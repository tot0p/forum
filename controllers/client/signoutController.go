package client

import (
	"forum/tools/session"
	"net/http"
)

type SignOutPage struct {
}

func (s *SignOutPage) ServeHTTP(w http.ResponseWriter, r *http.Request, m map[string]string) {
	session.GlobalSessions.SessionDestroy(w, r)
	http.Redirect(w, r, "/", http.StatusOK)
}
