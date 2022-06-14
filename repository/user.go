package repository

import (
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"strings"

	"forum/models"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

func CreateUserTable() error {
	_, err := forumDatabase.ExecuteStatement(`CREATE TABLE user (
		UUID           STRING  PRIMARY KEY
							   NOT NULL ON CONFLICT ABORT,
		profilePicture STRING,
		username       STRING  NOT NULL,
		password       STRING  NOT NULL,
		email          STRING  NOT NULL,
		firstName      STRING,
		lastName       STRING,
		riotId         STRING,
		oauthToken     STRING,
		birthDate      STRING  NOT NULL,
		genre          STRING,
		role           STRING  NOT NULL,
		title          STRING,
		bio            STRING,
		premium        INTEGER,
		follows        STRING
	);
	
	`)
	return err
}

func CreateUser(username, password, email, firstname, lastName, birthDate, oauthToken, genre, bio, riotId string, profilePicture []byte, premium int) error {
	user := models.User{
		UUID:           GenerateUUID(),
		ProfilePicture: profilePicture,
		Username:       username,
		Password:       HashPassword(password),
		Email:          email,
		FirstName:      firstname,
		LastName:       lastName,
		BirthDate:      birthDate,
		OauthToken:     oauthToken,
		Genre:          genre,
		Bio:            bio,
		Premium:        premium,
		RiotId:         riotId,
	}
	return InsertUserTable(user)
}

func InsertUserTable(user models.User) error {
	_, err := GetUser("username", user.Username)
	if err == nil {
		return errors.New("username already taken")
	}
	_, err = GetUser("email", user.Email)
	if err == nil {
		return errors.New("email already taken")
	}
	_, err = forumDatabase.ExecuteStatement(fmt.Sprintf(`INSERT INTO user
	VALUES (
		'%s',
		'%s',
		'%s',
		'%s',
		'%s',
		'%s',
		'%s',
		'%s',
		'%s',
		'%s',
		'%s',
		'%s',
		'%s',
		'%s',
		'%d',
		'%s'
	);`, user.UUID, user.ProfilePicture, strings.Replace(user.Username, "'", "''", -1), user.Password, user.Email, strings.Replace(user.FirstName, "'", "''", -1), strings.Replace(user.LastName, "'", "''", -1), user.RiotId, user.OauthToken, user.BirthDate, strings.Replace(user.Genre, "'", "''", -1), strings.Replace(user.Role, "'", "''", -1), strings.Replace(user.Title, "'", "''", -1), strings.Replace(user.Bio, "'", "''", -1), user.Premium, strings.Replace(user.Follows, "'", "''", -1)))
	return err
}

func GetUser(searchColumn, searchValue string) (*models.User, error) {
	var str string
	rows, err := forumDatabase.QuerryData(fmt.Sprintf(`SELECT UUID,
	profilePicture,
	username,
	password,
	email,
	firstName,
	lastName,
	riotId,
	oauthToken,
	birthDate,
	genre,
	role,
	title,
	bio,
	premium,
	follows
FROM user WHERE
%s = '%s';`, searchColumn, strings.Replace(searchValue, "'", "''", -1)))
	isPresent := false
	user := new(models.User)
	for rows.Next() {
		rows.Scan(&user.UUID, &str, &user.Username, &user.Password, &user.Email, &user.FirstName, &user.LastName, &user.RiotId, &user.OauthToken, &user.BirthDate, &user.Genre, &user.Role, &user.Title, &user.Bio, &user.Premium, &user.Follows)
		if err != nil {
			log.Fatal(err)
		}
		isPresent = true
	}
	if !isPresent {
		return nil, errors.New("user does not exist")
	}
	user.ProfilePicture, err = hex.DecodeString(str)
	return user, err
}

func GetAllUser() (*[]models.User, error) {
	var str string
	rows, err := forumDatabase.QuerryData(`SELECT UUID,
	profilePicture,
	username,
	password,
	email,
	firstName,
	lastName,
	riotId,
	oauthToken,
	birthDate,
	genre,
	role,
	title,
	bio,
	premium,
	follows
	FROM user;`)
	userTable := new([]models.User)
	for rows.Next() {
		user := new(models.User)
		rows.Scan(&user.UUID, &str, &user.Username, &user.Password, &user.Email, &user.FirstName, &user.LastName, &user.RiotId, &user.OauthToken, &user.BirthDate, &user.Genre, &user.Role, &user.Title, &user.Bio, &user.Premium, &user.Follows)
		user.ProfilePicture, err = hex.DecodeString(str)
		*userTable = append(*userTable, *user)
		if err != nil {
			log.Fatal(err)
		}
	}
	return userTable, err
}

func PostUser(user models.User, searchColumn, searchValue string) error {
	_, err := forumDatabase.ExecuteStatement(fmt.Sprintf(`UPDATE user
	SET UUID = '%s',
		profilePicture = '%02x',
		username = '%s',
		password = '%s',
		email = '%s',
		firstName = '%s',
		lastName = '%s',
		riotId = '%s',
		oauthToken = '%s',
		birthDate = '%s',
		genre = '%s',
		role = '%s',
		title = '%s',
		bio = '%s',
		premium = '%d',
		follows = '%s'
  WHERE %s = '%s';`, user.UUID, user.ProfilePicture, strings.Replace(user.Username, "'", "''", -1), user.Password, user.Email, strings.Replace(user.FirstName, "'", "''", -1), strings.Replace(user.LastName, "'", "''", -1), user.RiotId, user.OauthToken, user.BirthDate, strings.Replace(user.Genre, "'", "''", -1), strings.Replace(user.Role, "'", "''", -1), strings.Replace(user.Title, "'", "''", -1), strings.Replace(user.Bio, "'", "''", -1), user.Premium, strings.Replace(user.Follows, "'", "''", -1), searchColumn, strings.Replace(searchValue, "'", "''", -1)))
	return err
}

func GetUserUsername(UUID string) string {
	rows, err := forumDatabase.QuerryData(fmt.Sprintf("SELECT username FROM user WHERE UUID = '%s' ", UUID))
	if err != nil {
		return ""
	}
	var username string
	for rows.Next() {
		rows.Scan(&username)
		if err != nil {
			return ""
		}
	}
	return username
}

func DeleteUser(searchColumn, searchValue string) error {
	_, err := forumDatabase.ExecuteStatement(fmt.Sprintf(`DELETE FROM user
	WHERE %s = '%s';`, searchColumn, searchValue))
	return err

}

func ResetUserTable() error {
	_, err := forumDatabase.ExecuteStatement(`DELETE FROM USER;`)
	return err
}

func DropUserTable() error {
	_, err := forumDatabase.ExecuteStatement(`DROP TABLE USER;`)
	return err
}

func GenerateUUID() string {
	return uuid.NewString()
}

func HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}
