package authorization

import (
	"net/http"
	"strings"
)

//Function to set the token's authorization
func SetAuthorizationBearer(token string, r *http.Request) {
	var bearer = "Bearer " + token
	r.Header.Add("Authorization", bearer)
}

//Function to get the token's authorization
func GetAuthorizationBearer(w http.ResponseWriter, r *http.Request) string {
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer")
	if len(splitToken) != 2 {
		return ""
	}
	reqToken = strings.TrimSpace(splitToken[1])
	return reqToken
}
