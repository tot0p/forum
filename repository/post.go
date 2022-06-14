package repository

import (
	"encoding/hex"
	"fmt"
	"forum/models"
	"forum/tools/session"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func CreatePostTable() error {
	_, err := forumDatabase.ExecuteStatement(`CREATE TABLE post (
		id          INTEGER PRIMARY KEY AUTOINCREMENT,
		title       STRING  NOT NULL,
		description STRING,
		image       STRING,
		tags        STRING  NOT NULL,
		comments    STRING,
		nsfw        INTEGER,
		publishDate STRING  NOT NULL,
		upvotes     STRING,
		downvotes   STRING,
		owner       STRING  REFERENCES user (UUID),
		parent      STRING  REFERENCES subject (id) 
	);
	`)
	return err
}

func CreatePost(title, description, ownerSID, parentPost string, image []byte, nsfw int, tags []string) error {
	sess, err := session.GlobalSessions.Provider.SessionRead(ownerSID)
	if err != nil {
		return err
	}
	SID, err := sess.Get("UUID")
	if err != nil {
		return err
	}
	post := models.Post{
		Title:       title,
		Description: description,
		NSFW:        nsfw,
		Image:       image,
		PublishDate: time.Now().Format("2006-01-02 15:04:05"),
		Owner:       SID.(string),
		Parent:      parentPost,
	}
	post.Tags = post.ConvertSliceToString(tags)
	return InsertPostTable(post)
}

func InsertPostTable(post models.Post) error {
	_, err := forumDatabase.ExecuteStatement(fmt.Sprintf(`INSERT INTO post (
		title,
		description,
		image,
		tags,
		comments,
		nsfw,
		publishDate,
		upvotes,
		downvotes,
		owner,
		parent
	)
	VALUES (
		'%s',
		'%s',
		'%02x',
		'%s',
		'%s',
		'%d',
		'%s',
		'%s',
		'%s',
		'%s',
		'%s'
	);
`, post.Title, post.Description, post.Image, post.Tags, post.Comments, post.NSFW, post.PublishDate, post.UpVotes, post.DownVotes, post.Owner, post.Parent))
	if err != nil {
		return err
	}
	subjToUpdate, err := GetSubject("id", post.Parent)
	if err != nil {
		return err
	}
	timeNow := time.Now().Format("2006-01-02 15:04:05")
	subjToUpdate.LastPostDate = timeNow
	err = PostSubject(*subjToUpdate, "id", subjToUpdate.Id)
	if err != nil {
		return err
	}
	return err
}

func GetPost(searchColumn, searchValue string) (*models.Post, error) {
	var str string
	rows, err := forumDatabase.QuerryData(fmt.Sprintf(`SELECT id,
	title,
	description,
	image,
	tags,
	comments,
	nsfw,
	publishDate,
	upvotes,
	downvotes,
	owner,
	parent
FROM post WHERE
%s = '%s';`, searchColumn, searchValue))
	post := new(models.Post)
	for rows.Next() {
		rows.Scan(&post.Id, &post.Title, &post.Description, &str, &post.Tags, &post.Comments, &post.NSFW, &post.PublishDate, &post.UpVotes, &post.DownVotes, &post.Owner, &post.Parent)
		if err != nil {
			log.Fatal(err)
		}
	}
	post.Image, err = hex.DecodeString(str)
	return post, err
}

func GetAllPost() (*[]models.Post, error) {
	var str string
	rows, err := forumDatabase.QuerryData(`SELECT id,
	title,
	description,
	image,
	tags,
	comments,
	nsfw,
	publishDate,
	upvotes,
	downvotes,
	owner,
	parent
FROM post;`)
	postTable := new([]models.Post)
	for rows.Next() {
		post := new(models.Post)
		rows.Scan(&post.Id, &post.Title, &post.Description, &str, &post.Tags, &post.Comments, &post.NSFW, &post.PublishDate, &post.UpVotes, &post.DownVotes, &post.Owner, &post.Parent)
		post.Image, err = hex.DecodeString(str)
		*postTable = append(*postTable, *post)
		if err != nil {
			log.Fatal(err)
		}
	}
	return postTable, err
}

func PostPost(post models.Post, searchColumn, searchValue string) error {
	_, err := forumDatabase.ExecuteStatement(fmt.Sprintf(`UPDATE post
	SET title = '%s',
		description = '%s',
		image = '%02x',
		tags = '%s',
		comments = '%s',
		nsfw = '%d',
		publishDate = '%s',
		upvotes = '%s',
		downvotes = '%s',
		owner = '%s',
		parent = '%s'
  WHERE %s = '%s';`, post.Title, post.Description, post.Image, post.Tags, post.Comments, post.NSFW, post.PublishDate, post.UpVotes, post.DownVotes, post.Owner, post.Parent, searchColumn, searchValue))
	return err
}

func DeletePost(searchColumn, searchValue string) error {
	_, err := forumDatabase.ExecuteStatement(fmt.Sprintf(`DELETE FROM post
	WHERE %s = '%s';`, searchColumn, searchValue))
	return err
}

func ResetPostTable() error {
	_, err := forumDatabase.ExecuteStatement(`DELETE FROM post;`)
	return err
}

func DropPostTable() error {
	_, err := forumDatabase.ExecuteStatement(`DROP TABLE post;`)
	return err
}

func GetLastPost() (*[]models.Post, error) {
	//SELECT id,title,description,image,tags,comments,nsfw,publishDate,	upvotes,	downvotes,	owner,	parent FROM post ORDER BY publishDate desc
	var str string
	rows, err := forumDatabase.QuerryData(`SELECT id,title,description,image,tags,comments,nsfw,publishDate,	upvotes,	downvotes,	owner,	parent FROM post ORDER BY publishDate desc;`)
	postTable := new([]models.Post)
	for rows.Next() {
		post := new(models.Post)
		rows.Scan(&post.Id, &post.Title, &post.Description, &str, &post.Tags, &post.Comments, &post.NSFW, &post.PublishDate, &post.UpVotes, &post.DownVotes, &post.Owner, &post.Parent)
		post.Image, err = hex.DecodeString(str)
		*postTable = append(*postTable, *post)
		if err != nil {
			log.Fatal(err)
		}
	}
	return postTable, err
}

func GetPostsBySubjectId(id string) (*[]models.Post, error) {
	var str string
	rows, err := forumDatabase.QuerryData(fmt.Sprintf(`SELECT id,title,description,image,tags,comments,nsfw,publishDate,	upvotes,	downvotes,	owner,	parent FROM post where parent = %s;`, id))
	postTable := new([]models.Post)
	for rows.Next() {
		post := new(models.Post)
		rows.Scan(&post.Id, &post.Title, &post.Description, &str, &post.Tags, &post.Comments, &post.NSFW, &post.PublishDate, &post.UpVotes, &post.DownVotes, &post.Owner, &post.Parent)
		post.Image, err = hex.DecodeString(str)
		*postTable = append(*postTable, *post)
		if err != nil {
			log.Fatal(err)
		}
	}
	return postTable, err
}

func GetPostsByUserId(id string) (*[]models.Post, error) {
	var str string
	rows, err := forumDatabase.QuerryData(fmt.Sprintf(`SELECT id,title,description,image,tags,comments,nsfw,publishDate,	upvotes,	downvotes,	owner,	parent FROM post where owner = "%s";`, id))
	postTable := new([]models.Post)
	for rows.Next() {
		post := new(models.Post)
		rows.Scan(&post.Id, &post.Title, &post.Description, &str, &post.Tags, &post.Comments, &post.NSFW, &post.PublishDate, &post.UpVotes, &post.DownVotes, &post.Owner, &post.Parent)
		post.Image, err = hex.DecodeString(str)
		*postTable = append(*postTable, *post)
		if err != nil {
			log.Fatal(err)
		}
	}
	return postTable, err
}
