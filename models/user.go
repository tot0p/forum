package models

import (
	"encoding/base64"
)

type User struct {
	UUID           string `json:"uuid"`
	ProfilePicture []byte `json:"profilepicture"`
	Username       string `json:"username"`
	Password       string `json:"password"`
	Email          string `json:"email"`
	FirstName      string `json:"firstname"`
	LastName       string `json:"lastname"`
	RiotId         string `json:"riotid"`
	BirthDate      string `json:"birthdate"`
	OauthToken     string `json:"oauthtoken"`
	Genre          string `json:"genre"`
	Role           string `json:"role"`
	Title          string `json:"title"`
	Bio            string `json:"bio"`
	Premium        int    `json:"premium"`
}

//Method to convert Hex to Base64
func (user *User) ToBase64() string {
	return base64.StdEncoding.EncodeToString(user.ProfilePicture)
}

//Method to defend the user of giving his OauthToken and Password
func (u *User) Sec() {
	u.Password = ""
	u.OauthToken = ""
}
