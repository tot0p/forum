package verif

import (
	"net/mail"
	"time"
	"unicode"
)

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

func EmailVerif(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func DateVerif(Date string) bool {
	_, err := time.Parse("2006-01-02", Date)
	return err == nil
}

func NSFWVerif(NSFW int) bool {
	if NSFW == 0 || NSFW == 1 {
		return true
	}
	return false
}

func RiotVerif(RiotId string) bool {
	if RiotId == "" {
		return true
	} else {
		return true
	}
}

func ImageVerif(Image string) bool {
	return true
}
