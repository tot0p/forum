package repository

import (
	"fmt"
	"forum/models"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func CreateCommentTable() error {
	_, err := forumDatabase.ExecuteStatement(`CREATE TABLE comment (
		id        INTEGER PRIMARY KEY AUTOINCREMENT,
		owner     STRING  REFERENCES user (UUID),
		content   STRING  NOT NULL,
		upvotes   STRING,
		downvotes STRING,
		responses STRING
	);	
	`)
	return err
}

func CreateComment(content, ownerUUID string) error {
	comment := models.Comment{
		Content: content,
		Owner:   ownerUUID,
	}
	return InsertCommentTable(comment)
}

func InsertCommentTable(comment models.Comment) error {
	_, err := forumDatabase.ExecuteStatement(fmt.Sprintf(`INSERT INTO comment (
		owner,
		content,
		upvotes,
		downvotes,
		responses
	)
	VALUES (
		'%s',
		'%s',
		'%s',
		'%s',
		'%s'
	);

`, comment.Owner, comment.Content, comment.UpVotes, comment.DownVotes, comment.Responses))
	return err
}

func GetComment(searchColumn, searchValue string) (*models.Comment, error) {
	rows, err := forumDatabase.QuerryData(fmt.Sprintf(`SELECT id,
	owner,
	content,
	upvotes,
	downvotes,
	responses
FROM comment WHERE
%s = '%s';`, searchColumn, searchValue))
	comment := new(models.Comment)
	for rows.Next() {
		rows.Scan(&comment.Id, &comment.Owner, &comment.Content, &comment.UpVotes, &comment.DownVotes, &comment.Responses)
		if err != nil {
			log.Fatal(err)
		}
	}
	return comment, err
}

func GetAllComment() (*[]models.Comment, error) {
	rows, err := forumDatabase.QuerryData(`SELECT id,
	owner,
	content,
	upvotes,
	downvotes,
	responses
FROM comment;
`)
	commentTable := new([]models.Comment)
	for rows.Next() {
		comment := new(models.Comment)
		rows.Scan(&comment.Id, &comment.Owner, &comment.Content, &comment.UpVotes, &comment.DownVotes, &comment.Responses)
		*commentTable = append(*commentTable, *comment)
		if err != nil {
			log.Fatal(err)
		}
	}
	return commentTable, err
}

func PostComment(comment models.Comment, searchColumn, searchValue string) error {
	_, err := forumDatabase.ExecuteStatement(fmt.Sprintf(`UPDATE comment
	SET	owner = '%s',
		content = '%s',
		upvotes = '%s',
		downvotes = '%s',
		responses = '%s'
  WHERE %s = '%s';`, comment.Owner, comment.Content, comment.UpVotes, comment.DownVotes, comment.Responses, searchColumn, searchValue))
	return err
}

func DeleteComment(searchColumn, searchValue string) error {
	_, err := forumDatabase.ExecuteStatement(fmt.Sprintf(`DELETE FROM comment
	WHERE %s = '%s';`, searchColumn, searchValue))
	return err
}

func ResetCommentTable() error {
	_, err := forumDatabase.ExecuteStatement(`DELETE FROM comment;`)
	return err
}

func DropCommentTable() error {
	_, err := forumDatabase.ExecuteStatement(`DROP TABLE comment;`)
	return err
}
