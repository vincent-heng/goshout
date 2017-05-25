package main

import (
	"time"
)

type Message struct {
	Nickname	string		`json:"nickname"`
	Text		string		`json:"text"`
	Timestamp	time.Time	`json:"timestamp"`
}

func isMessageValid(message Message) bool {
	return len(message.Nickname) > 0 && len(message.Text) > 0
}
