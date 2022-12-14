package model

import "golang.org/x/text/language"

type Message struct {
	ID      string
	Message string
	Fuzzy   bool
}

type Messages struct {
	TranslationID string
	Labels        map[string]string
	Language      language.Tag
	Messages      []Message
}
