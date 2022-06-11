package datapage

import "net/http"

type SubjectData struct {
	Path string
}

func (s *SubjectData) ServeHTTP(w http.ResponseWriter, r *http.Request, m map[string]string) {

}
