package usecase

import (
	"context"
	"errors"

	"github.com/itsuki-yamada/gonew/workspace/domain"
	log "github.com/itsuki-yamada/gonew/workspace/log"
)

type FetchPingUsecase struct{}

func reportError(err *error) *error {
	if err != nil {
		log.Logger.Error("fetch_ping_usecase", "Exec", "error", err)
	}
	return err
}

func (ps *FetchPingUsecase) Exec(c context.Context, ping domain.Pinger) (msg domain.Message, e error) {
	defer reportError(&e)
	log.Logger.Info("fetch_ping_usecase", "Exec", "start")

	return ping.Pong(), errors.New("error")
}
