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

func (post *Post) ToBase64() string {
	return base64.StdEncoding.EncodeToString(post.Image)
}

func (post *Post) ConvertTags() []string {
	return strings.Split(post.Tags, "#")
}

func (post *Post) ConvertComments() []string {
	return strings.Split(post.Comments, "#")
}

func (post *Post) ConvertUpVotes() []string {
	return strings.Split(post.UpVotes, "#")
}

func (post *Post) ConvertDownVotes() []string {
	return strings.Split(post.DownVotes, "#")
}

func (post *Post) ConvertSliceToString(a []string) string {
	return strings.Join(a, "#")
}
