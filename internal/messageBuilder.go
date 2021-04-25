package internal

import (
	"github.com/fatih/color"
)

type Message struct {
	message string
}

func NewMessage(message string, col ...color.Attribute) *Message {
	msg := color.CyanString("[LUN] ") + color.New(col...).Sprint(message)
	return &Message{message: msg}
}

func (msg *Message) Then(message string, col ...color.Attribute) *Message {
	msg.message += " " + color.New(col...).Sprint(message)
	return msg
}

func (msg *Message) String() string {
	return msg.message
}