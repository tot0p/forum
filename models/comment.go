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

//model for upvotes
func (comment *Comment) ConvertUpVotes() []string {
	return strings.Split(comment.UpVotes, "#")[1:]
}

//model for downvote
func (comment *Comment) ConvertDownVotes() []string {
	return strings.Split(comment.DownVotes, "#")[1:]
}

//Function to convert Slice to String
func (comment *Comment) ConvertSliceToString(a []string) string {
	return strings.Join(a, "#")
}
