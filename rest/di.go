package main

import (
	handler "github.com/itsuki-yamada/gonew/rest/handler/systems/ping"
	usecase "github.com/itsuki-yamada/gonew/usecase/systems/ping"
)

type (
	Container struct {
		PingHandler *handler.PingHandler
	}
)

func NewDI() *Container {
	return &Container{
		PingHandler: handler.NewPingHandler(&usecase.PingUsecase{}),
	}
}
