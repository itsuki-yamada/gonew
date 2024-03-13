package usecase

import (
	"context"

	"github.com/itsuki-yamada/gonew/workspace/domain"
)

type FetchPingUsecase struct{}

func (ps *FetchPingUsecase) Exec(c context.Context, ping domain.Pinger) domain.Message {
	return ping.Pong()
}
