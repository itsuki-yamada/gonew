package usecase

import (
	"context"

	"github.com/itsuki-yamada/gonew/workspace/domain"
	log "github.com/itsuki-yamada/gonew/workspace/log"
)

type FetchPingUsecase struct{}

func (ps *FetchPingUsecase) Exec(c context.Context, ping domain.Pinger) domain.Message {
	log.Logger.Info("fetch_ping_usecase", "Exec", "start");
	return ping.Pong()
}
