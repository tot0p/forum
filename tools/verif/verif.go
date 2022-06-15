package verif

import (
	"net/mail"
	"time"
	"unicode"
)

//Function to check if the password is valid
func PasswordVerif(password string) bool {
	var (
		hasMinLen  = false
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)
	if len(password) >= 8 {
		hasMinLen = true
	}
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}
	return hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial
}

//Function to check if the email is valid
func EmailVerif(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

//Function to check if the date is valid
func DateVerif(Date string) bool {
	_, err := time.Parse("2006-01-02", Date)
	return err == nil
}

//Function to check if the nsfw is valid
func NSFWVerif(NSFW int) bool {
	if NSFW == 0 || NSFW == 1 {
		return true
	}
	return false
}

//Function to check if the Riot id is valid
func RiotVerif(RiotId string) bool {
	if RiotId == "" {
		return true
	} else {
		return true
	}
}

//Function to check if the image is valid
func ImageVerif(Image string) bool {
	return true
}
