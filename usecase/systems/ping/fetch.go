package ping

import (
	"context"

	"github.com/itsuki-yamada/gonew/domain"
)

func (ps *PingUsecase) FetchPing(c context.Context, ping domain.Pinger) domain.Message {
	return ping.Pong()
}
