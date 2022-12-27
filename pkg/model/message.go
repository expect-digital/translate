package model

import "golang.org/x/text/language"

type Message struct {
	ID      string
	Message string
	Fuzzy   bool
}

type Messages struct {
	Language language.Tag
	Messages []Message
	Labels   map[string]string
}
