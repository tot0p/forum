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

//Method to convert Hex to Base64
func (post *Post) ToBase64() string {
	return base64.StdEncoding.EncodeToString(post.Image)
}

//Model for tags in post
func (post *Post) ConvertTags() []string {
	return strings.Split(post.Tags, "#")[1:]
}

//Model for comments in post
func (post *Post) ConvertComments() []string {
	return strings.Split(post.Comments, "#")
}

//Model for upvotes in post
func (post *Post) ConvertUpVotes() []string {
	return strings.Split(post.UpVotes, "#")[1:]
}

//Model for downvotes in post
func (post *Post) ConvertDownVotes() []string {
	return strings.Split(post.DownVotes, "#")[1:]
}

//Method to convert slice to string
func (post *Post) ConvertSliceToString(a []string) string {
	return strings.Join(a, "#")
}

func (post *Post) GetVote() Vote {
	return Vote{UpVote: strings.Count(post.UpVotes, "#"), DownVote: strings.Count(post.DownVotes, "#")}
}

func (post *Post) IsLike(UUID string) bool {
	return strings.Contains(post.UpVotes, UUID)
}

func (post *Post) IsHate(UUID string) bool {
	return strings.Contains(post.DownVotes, UUID)
}

func (post *Post) HaveTag(tag string) bool {
	for _, i := range post.ConvertTags() {
		if tag == strings.ToLower(i) {
			return true
		}
	}
	return false
}
