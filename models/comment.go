package models

import "strings"

type Comment struct {
	Id          string `json:"id"`
	Owner       string `json:"owner"`
	Content     string `json:"content"`
	UpVotes     string `json:"upvotes"`
	DownVotes   string `json:"downvotes"`
	Parent      string `json:"parent"`
	PublishDate string `json:"publishDate"`
}

//Models for comment

//Model for upvotes
func (comment *Comment) ConvertUpVotes() []string {
	return strings.Split(comment.UpVotes, "#")[1:]
}

//Model for downvotes
func (comment *Comment) ConvertDownVotes() []string {
	return strings.Split(comment.DownVotes, "#")[1:]
}

//Method to convert slice to string
func (comment *Comment) ConvertSliceToString(a []string) string {
	return strings.Join(a, "#")
}

func (comment *Comment) GetVote() Vote {
	return Vote{UpVote: strings.Count(comment.UpVotes, "#"), DownVote: strings.Count(comment.DownVotes, "#")}
}

func (comment *Comment) IsLike(UUID string) bool {
	return strings.Contains(comment.UpVotes, UUID)
}

func (comment *Comment) IsHate(UUID string) bool {
	return strings.Contains(comment.DownVotes, UUID)
}
