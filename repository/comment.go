package repository

import (
	"fmt"
	"forum/models"
	"log"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

//Function for sending a request to the database to create the comment table
func CreateCommentTable() error {
	_, err := forumDatabase.ExecuteStatement(`CREATE TABLE comment (
		id        INTEGER PRIMARY KEY AUTOINCREMENT,
		owner     STRING  REFERENCES user (UUID),
		content   STRING  NOT NULL,
		upvotes   STRING,
		downvotes STRING
	);	
	`)
	return err
}

//Function to simplify the request
func CreateComment(content, ownerUUID, parentId string) error {
	comment := models.Comment{
		Content:     content,
		Owner:       ownerUUID,
		Parent:      parentId,
		PublishDate: time.Now().Format("2006-01-02 15:04:05.000000"),
	}
	return InsertCommentTable(comment)
}

//Function to send the request
func InsertCommentTable(comment models.Comment) error {
	_, err := forumDatabase.ExecuteStatement(fmt.Sprintf(`INSERT INTO comment (
		owner,
		content,
		upvotes,
		downvotes,
		publishDate,
		parent
	)
	VALUES (
		'%s',
		'%s',
		'%s',
		'%s',
		'%s',
		'%s'
	);

`, comment.Owner, strings.Replace(comment.Content, "'", "''", -1), comment.UpVotes, comment.DownVotes, comment.PublishDate, comment.Parent))
	if err != nil {
		return err
	}
	postToUpdate, err := GetPost("id", comment.Parent)
	if err != nil {
		return err
	}
	postNew, err := GetComment("publishDate", comment.PublishDate)
	if err != nil {
		return err
	}
	comments := postToUpdate.ConvertComments()
	comments = append(comments, postNew.Id)
	postToUpdate.Comments = postToUpdate.ConvertSliceToString(comments)
	err = PostPost(*postToUpdate, "id", postToUpdate.Id)
	if err != nil {
		return err
	}
	return err
}

//Function to get a comment in the database
func GetComment(searchColumn, searchValue string) (*models.Comment, error) {
	rows, err := forumDatabase.QuerryData(fmt.Sprintf(`SELECT id,
	owner,
	content,
	upvotes,
	downvotes,
	publishDate,
	parent
FROM comment WHERE
%s = '%s';`, searchColumn, strings.Replace(searchValue, "'", "''", -1)))
	comment := new(models.Comment)
	for rows.Next() {
		rows.Scan(&comment.Id, &comment.Owner, &comment.Content, &comment.UpVotes, &comment.DownVotes, &comment.PublishDate, &comment.Parent)
		if err != nil {
			log.Fatal(err)
		}
	}
	return comment, err
}

//Function to get all the comments from the database
func GetAllComment() (*[]models.Comment, error) {
	rows, err := forumDatabase.QuerryData(`SELECT id,
	owner,
	content,
	upvotes,
	downvotes,
	publishDate,
	parent
FROM comment;
`)
	commentTable := new([]models.Comment)
	for rows.Next() {
		comment := new(models.Comment)
		rows.Scan(&comment.Id, &comment.Owner, &comment.Content, &comment.UpVotes, &comment.DownVotes, &comment.PublishDate, &comment.Parent)
		*commentTable = append(*commentTable, *comment)
		if err != nil {
			log.Fatal(err)
		}
	}
	return commentTable, err
}

//Function to post a comment in the database
func PostComment(comment models.Comment, searchColumn, searchValue string) error {
	_, err := forumDatabase.ExecuteStatement(fmt.Sprintf(`UPDATE comment
	SET	owner = '%s',
		content = '%s',
		upvotes = '%s',
		downvotes = '%s',
		publishDate = '%s',
		parent = '%s'
  WHERE %s = '%s';`, comment.Owner, strings.Replace(comment.Content, "'", "''", -1), comment.UpVotes, comment.DownVotes, comment.PublishDate, comment.Parent, searchColumn, searchValue))
	return err
}

//Function to delete a comment from the database
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

func GetCommentsByPostId(id string) (*[]models.Comment, error) {
	rows, err := forumDatabase.QuerryData(fmt.Sprintf(`SELECT id,
	owner,
	content,
	upvotes,
	downvotes,
	publishDate,
	parent
FROM comment where parent = %s;
`, id))
	commentTable := new([]models.Comment)
	for rows.Next() {
		comment := new(models.Comment)
		rows.Scan(&comment.Id, &comment.Owner, &comment.Content, &comment.UpVotes, &comment.DownVotes, &comment.PublishDate, &comment.Parent)
		*commentTable = append(*commentTable, *comment)
		if err != nil {
			log.Fatal(err)
		}
	}
	return commentTable, err
}
