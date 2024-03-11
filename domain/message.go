package domain

type Message string

func newMessage(msg string) Message {
	if msg == "" {
		msg = "pong"
	}
	return Message(msg)
}
