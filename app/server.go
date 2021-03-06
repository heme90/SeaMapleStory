package app

import "SeaMaple/dto/static"

type IServer interface {
	SetRouter(router IRouter)
}

type IRouter interface {
	SetUseCase(usecase interface{})
	SetLogger(logger static.ILogger)
}