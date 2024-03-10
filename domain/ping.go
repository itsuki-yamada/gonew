package domain

type (
	Ping struct {
		Message string
	}

	Pinger interface {
		Pong() string
	}
)

func (p *Ping) Pong() string {
	return p.Message
}
