package models

import (
	"encoding/base64"
	"strings"
)

type Subject struct {
	Id           string `json:"id"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	NSFW         int    `json:"nsfw"`
	Image        []byte `json:"image"`
	Tags         string `json:"tags"`
	UpVotes      string `json:"upvotes"`
	DownVotes    string `json:"downvotes"`
	PublishDate  string `json:"publishdate"`
	LastPostDate string `json:"lastpostdate"`
	Owner        string `json:"owner"`
}

//Function to convert Hex to Base64
func (subject *Subject) ToBase64() string {
	return base64.StdEncoding.EncodeToString(subject.Image)
}

//model for tags in subject
func (subject *Subject) ConvertTags() []string {
	return strings.Split(subject.Tags, "#")[1:]
}

//model for upvotes in subject
func (subject *Subject) ConvertUpVotes() []string {
	return strings.Split(subject.UpVotes, "#")[1:]
}

//model for down in subject
func (subject *Subject) ConvertDownVotes() []string {
	return strings.Split(subject.DownVotes, "#")[1:]
}

//Function to convert Slice to String
func (subject *Subject) ConvertSliceToString(a []string) string {
	return strings.Join(a, "#")
}
