package usecase

import (
	"context"

	"github.com/itsuki-yamada/gonew/domain"
)

func (ps *PingUsecase) FetchPing(c context.Context) string {
	p := domain.Ping{
		Message: "pong",
	}
	return p.Pong()
}
