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
	AllPosts     string `json:"allposts"`
	Owner        string `json:"owner"`
}

func (subject *Subject) ToBase64() string {
	return base64.StdEncoding.EncodeToString(subject.Image)
}

func (subject *Subject) ConvertTagsTo() []string {
	return strings.Split(subject.Tags, "#")[1:]
}

func (subject *Subject) ConvertTags() []string {
	return strings.Split(subject.Tags, "#")[1:]
}

func (subject *Subject) ConvertUpVotes() []string {
	return strings.Split(subject.UpVotes, "#")[1:]
}

func (subject *Subject) ConvertDownVotes() []string {
	return strings.Split(subject.DownVotes, "#")[1:]
}

func (subject *Subject) ConvertAllPosts() []string {
	return strings.Split(subject.AllPosts, "#")[1:]
}

func (subject *Subject) ConvertSliceToString(a []string) string {
	return strings.Join(a, "#")
}
