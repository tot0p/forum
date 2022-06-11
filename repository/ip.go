package repository

import (
	"fmt"
	"forum/models"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func CreateIPTable() error {
	_, err := forumDatabase.ExecuteStatement(`CREATE TABLE ip (
		id   INTEGER PRIMARY KEY AUTOINCREMENT,
		ip   STRING,
		date STRING
	);	
	`)
	return err
}

func GetAllIp() (*[]models.Ip, error) {
	rows, err := forumDatabase.QuerryData(`SELECT id,ip,date FROM ip;`)
	postTable := new([]models.Ip)
	for rows.Next() {
		post := new(models.Ip)
		rows.Scan(&post.Id, &post.Ip, &post.Date)
		*postTable = append(*postTable, *post)
		if err != nil {
			log.Fatal(err)
		}
	}
	return postTable, err
}

func AddConnection(Ip string) error {
	_, err := forumDatabase.ExecuteStatement(fmt.Sprintf(`INSERT INTO ip (
		ip,
		date
	)
	VALUES (
		'%s',
		'%s'
	);`, Ip, time.Now().Format("2006-01-02 15:04:05")))
	return err
}
