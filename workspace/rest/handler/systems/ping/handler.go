package handler

import (
	"context"

	"github.com/itsuki-yamada/gonew/workspace/domain"
)

type (
	PingHandler struct {
		usecase FetchPingUsecase
	}

	FetchPingUsecase interface {
		Exec(c context.Context, ping domain.Pinger) (domain.Message, error)
	}
)

func NewPingHandler(u FetchPingUsecase) *PingHandler {
	return &PingHandler{
		usecase: u,
	}
}
