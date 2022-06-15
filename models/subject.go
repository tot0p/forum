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

//Method to convert Hex to Base64
func (subject *Subject) ToBase64() string {
	return base64.StdEncoding.EncodeToString(subject.Image)
}

//Model for tags in subject
func (subject *Subject) ConvertTags() []string {
	return strings.Split(subject.Tags, "#")[1:]
}

func (subject *Subject) HaveTag(tag string) bool {
	for _, i := range subject.ConvertTags() {
		if tag == strings.ToLower(i) {
			return true
		}
	}
	return false
}

//Model for upvotes in subject
func (subject *Subject) ConvertUpVotes() []string {
	return strings.Split(subject.UpVotes, "#")[1:]
}

//Model for downvotes in subject
func (subject *Subject) ConvertDownVotes() []string {
	return strings.Split(subject.DownVotes, "#")[1:]
}

//Method to convert slice to string
func (subject *Subject) ConvertSliceToString(a []string) string {
	return strings.Join(a, "#")
}

func (subject *Subject) GetVote() Vote {
	return Vote{UpVote: strings.Count(subject.UpVotes, "#"), DownVote: strings.Count(subject.DownVotes, "#")}
}

func (subject *Subject) IsLike(UUID string) bool {
	return strings.Contains(subject.UpVotes, UUID)
}

func (subject *Subject) IsHate(UUID string) bool {
	return strings.Contains(subject.DownVotes, UUID)
}
