package models

import (
	"encoding/base64"
	"strings"
)

type Post struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	NSFW        int    `json:"nsfw"`
	Image       []byte `json:"image"`
	Tags        string `json:"tags"`
	UpVotes     string `json:"upvotes"`
	DownVotes   string `json:"downvotes"`
	PublishDate string `json:"publishdate"`
	Comments    string `json:"comments"`
	Owner       string `json:"owner"`
	Parent      string `json:"parent"`
}

//Function to convert Hex to Base64
func (post *Post) ToBase64() string {
	return base64.StdEncoding.EncodeToString(post.Image)
}

//model for tags in post
func (post *Post) ConvertTags() []string {
	return strings.Split(post.Tags, "#")[1:]
}

//model for comments in post
func (post *Post) ConvertComments() []string {
	return strings.Split(post.Comments, "#")
}

//model for upvotes in post
func (post *Post) ConvertUpVotes() []string {
	return strings.Split(post.UpVotes, "#")[1:]
}

//model for downvotes in post
func (post *Post) ConvertDownVotes() []string {
	return strings.Split(post.DownVotes, "#")[1:]
}

//Function to convert Slice to String
func (post *Post) ConvertSliceToString(a []string) string {
	return strings.Join(a, "#")
}
