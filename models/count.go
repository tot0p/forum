package models

type Count struct {
	Nb int
}

type AllCount struct {
	Session, Subject, Post, User int
}

type Vote struct {
	UpVote   int
	DownVote int
}
