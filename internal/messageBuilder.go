package internal

import "github.com/ttacon/chalk"

type Message struct {
	message string
}

func NewMessage(color chalk.Color, message string) *Message {
	msg := chalk.Cyan.Color("[LUN] ") + color.Color(message)
	return &Message{message: msg}
}

func (msg *Message) ThenColor(color chalk.Color, message string) *Message {
	msg.message += " " + color.Color(message)
	return msg
}

func (msg *Message) ThenStyle(style chalk.TextStyle, message string) *Message {
	msg.message += " " + style.TextStyle(message)
	return msg
}

func (msg *Message) ThenColorStyle(color chalk.Color, style chalk.TextStyle, message string) *Message {
	msg.message += " " + color.Color(style.TextStyle(message))
	return msg
}

func (msg *Message) String() string {
	return msg.message
}