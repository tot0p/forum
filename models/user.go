package models

import "encoding/base64"

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
	Follows        string `json:"follows"`
}

// type UserSec struct {
// 	UUID           string `json:"uuid"`
// 	ProfilePicture []byte `json:"profilepicture"`
// 	Username       string `json:"username"`
// 	Email          string `json:"email"`
// 	FirstName      string `json:"firstname"`
// 	LastName       string `json:"lastname"`
// 	RiotId         string `json:"riotid"`
// 	BirthDate      string `json:"birthdate"`
// 	Genre          string `json:"genre"`
// 	Role           string `json:"role"`
// 	Title          string `json:"title"`
// 	Bio            string `json:"bio"`
// 	Premium        int    `json:"premium"`
// 	Follows        string `json:"follows"`
// }

// func (u User) ToUserSec() UserSec {
// 	return UserSec{UUID: u.UUID, ProfilePicture: u.ProfilePicture, Username: u.Username, Email: u.Email, FirstName: u.FirstName,
// 		LastName: u.LastName, RiotId: u.RiotId, BirthDate: u.BirthDate, Genre: u.Genre, Role: u.Role, Title: u.Title, Bio: u.Bio,
// 		Premium: u.Premium, Follows: u.Follows,
// 	}
// }

func (user *User) ToBase64() string {
	return base64.StdEncoding.EncodeToString(user.ProfilePicture)
}

func (u *User) Sec() {
	u.Password = ""
	u.OauthToken = ""
}
