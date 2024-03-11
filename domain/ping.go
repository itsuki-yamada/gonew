package domain

type (
	Pinger interface {
		Pong() Message
	}
	ping struct {
		message Message
	}
)

func NewPinger(msg string) Pinger {
	return &ping{
		message: newMessage(msg),
	}
}

func (p *ping) Pong() Message {
	return p.message
}
