package server

import "SeaMaple/dto/static"

type Router struct {
	logger static.ILogger
}

func (r *Router) SetUseCase(usecase interface{}) {
	panic("implement me")
}

func (r *Router) SetLogger(logger static.ILogger) {
	r.logger = logger
}
