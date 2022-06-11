package repository

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type DatabaseManager struct {
	IsInit   bool
	Database *sql.DB
}

var forumDatabase = new(DatabaseManager)

//Initialize and connect to the specified database
func (Db *DatabaseManager) Init(DbFile string) {
	if !Db.IsInit {
		var err error
		Db.Database, err = sql.Open("sqlite3", DbFile)
		if err != nil {
			log.Fatal(err)
		}
		Db.IsInit = true
	}
}

func InitializeDatabase(DbFile string) {
	forumDatabase.Init(DbFile)
}

func (Db *DatabaseManager) ExecuteStatement(stmt string) (sql.Result, error) {
	statement, err := Db.Database.Prepare(stmt)
	if err != nil {
		log.Fatal(err.Error())
	}
	return statement.Exec()
}

func (Db *DatabaseManager) QuerryData(stmt string) (*sql.Rows, error) {
	return Db.Database.Query(stmt)
}
