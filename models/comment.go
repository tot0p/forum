package models

import "strings"

type Comment struct {
	Id        string `json:"id"`
	Owner     string `json:"owner"`
	Content   string `json:"content"`
	UpVotes   string `json:"upvotes"`
	DownVotes string `json:"downvotes"`
	Responses string `json:"responses"`
}

func (comment *Comment) ConvertUpVotes() []string {
	return strings.Split(comment.UpVotes, "#")
}

func (comment *Comment) ConvertDownVotes() []string {
	return strings.Split(comment.DownVotes, "#")
}

func (comment *Comment) ConvertResponses() []string {
	return strings.Split(comment.Responses, "#")
}

func (comment *Comment) ConvertSliceToString(a []string) string {
	return strings.Join(a, "#")
}
