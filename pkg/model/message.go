package model

type Message struct {
	ID      string
	Message string
	Fuzzy   bool
}

type Messages struct {
	Language string
	Messages []Message
}
