package repository

import (
	"encoding/hex"
	"fmt"
	"log"
	"time"

	"forum/models"

	_ "github.com/mattn/go-sqlite3"
)

func CreateSubjectTable() error {
	_, err := forumDatabase.ExecuteStatement(`CREATE TABLE subject (
		id           INTEGER PRIMARY KEY AUTOINCREMENT,
		title        STRING  NOT NULL,
		description  STRING,
		nsfw         INT,
		image        STRING,
		tags         STRING  NOT NULL,
		upvotes      STRING,
		downvotes    STRING,
		publishDate  STRING  NOT NULL,
		lastPostDate STRING,
		allPosts     STRING,
		owner        STRING  REFERENCES subject (UUID) 
							 NOT NULL
	);
	`)
	return err
}

func CreateSubject(title, description, ownerUUID string, image []byte, nsfw int, tags []string) error {
	subject := models.Subject{
		Title:        title,
		Description:  description,
		NSFW:         nsfw,
		Image:        image,
		PublishDate:  time.Now().Format("2006-01-02 15:04:05"),
		LastPostDate: time.Now().Format("2006-01-02 15:04:05"),
		Owner:        ownerUUID,
	}
	subject.Tags = subject.ConvertSliceToString(tags)
	return InsertSubjectTable(subject)
}

func InsertSubjectTable(subject models.Subject) error {
	_, err := forumDatabase.ExecuteStatement(fmt.Sprintf(`INSERT INTO subject (
		title,
		description,
		nsfw,
		image,
		tags,
		upvotes,
		downvotes,
		publishDate,
		lastPostDate,
		allPosts,
		owner
	)
	VALUES (
		'%s',
		'%s',
		'%d',
		'%02x',
		'%s',
		'%s',
		'%s',
		'%s',
		'%s',
		'%s',
		'%s'
	);
`, subject.Title, subject.Description, subject.NSFW, subject.Image, subject.Tags, subject.UpVotes, subject.DownVotes, subject.PublishDate, subject.LastPostDate, subject.AllPosts, subject.Owner))
	return err
}

func GetSubject(searchColumn, searchValue string) (*models.Subject, error) {
	var str string
	rows, err := forumDatabase.QuerryData(fmt.Sprintf(`SELECT id,
	title,
	description,
	nsfw,
	image,
	tags,
	upvotes,
	downvotes,
	publishDate,
	lastPostDate,
	allPosts,
	owner
FROM subject WHERE
%s = '%s';`, searchColumn, searchValue))
	subject := new(models.Subject)
	for rows.Next() {
		rows.Scan(&subject.Id, &subject.Title, &subject.Description, &subject.NSFW, &str, &subject.Tags, &subject.UpVotes, &subject.DownVotes, &subject.PublishDate, &subject.LastPostDate, &subject.AllPosts, &subject.Owner)
		if err != nil {
			log.Fatal(err)
		}
	}
	subject.Image, err = hex.DecodeString(str)
	return subject, err
}

func GetAllSubject() (*[]models.Subject, error) {
	var str string
	rows, err := forumDatabase.QuerryData(`SELECT id,
	title,
	description,
	nsfw,
	image,
	tags,
	upvotes,
	downvotes,
	publishDate,
	lastPostDate,
	allPosts,
	owner
FROM subject;`)
	subjectTable := new([]models.Subject)
	for rows.Next() {
		subject := new(models.Subject)
		rows.Scan(&subject.Id, &subject.Title, &subject.Description, &subject.NSFW, &str, &subject.Tags, &subject.UpVotes, &subject.DownVotes, &subject.PublishDate, &subject.LastPostDate, &subject.AllPosts, &subject.Owner)
		subject.Image, err = hex.DecodeString(str)
		*subjectTable = append(*subjectTable, *subject)
		if err != nil {
			log.Fatal(err)
		}
	}
	return subjectTable, err
}

func PostSubject(subject models.Subject, searchColumn, searchValue string) error {
	_, err := forumDatabase.ExecuteStatement(fmt.Sprintf(`UPDATE subject
	SET title = '%s',
		description = '%s',
		nsfw = '%d',
		image = '%02x',
		tags = '%s',
		upvotes = '%s',
		downvotes = '%s',
		publishDate = '%s',
		lastPostDate = '%s',
		allPosts = '%s',
		owner = '%s'
  WHERE %s = '%s';`, subject.Title, subject.Description, subject.NSFW, subject.Image, subject.Tags, subject.UpVotes, subject.DownVotes, subject.PublishDate, subject.LastPostDate, subject.AllPosts, subject.Owner, searchColumn, searchValue))
	return err
}

func DeleteSubject(searchColumn, searchValue string) error {
	_, err := forumDatabase.ExecuteStatement(fmt.Sprintf(`DELETE FROM subject
	WHERE %s = '%s';`, searchColumn, searchValue))
	return err
}

func ResetSubjectTable() error {
	_, err := forumDatabase.ExecuteStatement(`DELETE FROM subject;`)
	return err
}

func DropSubjectTable() error {
	_, err := forumDatabase.ExecuteStatement(`DROP TABLE subject;`)
	return err
}

func GetSubjectLastUpdate() (*[]models.Subject, error) {
	// SELECT id,title,	description,nsfw,image,tags,upvotes,downvotes,publishDate,lastPostDate,	allPosts,owner FROM subject ORDER BY lastPostDate desc
	var str string
	rows, err := forumDatabase.QuerryData(`SELECT id,title,	description,nsfw,image,tags,upvotes,downvotes,publishDate,lastPostDate,	allPosts,owner FROM subject ORDER BY lastPostDate desc;`)
	subjectTable := new([]models.Subject)
	for rows.Next() {
		subject := new(models.Subject)
		rows.Scan(&subject.Id, &subject.Title, &subject.Description, &subject.NSFW, &str, &subject.Tags, &subject.UpVotes, &subject.DownVotes, &subject.PublishDate, &subject.LastPostDate, &subject.AllPosts, &subject.Owner)
		subject.Image, err = hex.DecodeString(str)
		*subjectTable = append(*subjectTable, *subject)
		if err != nil {
			log.Fatal(err)
		}
	}
	return subjectTable, err
}
