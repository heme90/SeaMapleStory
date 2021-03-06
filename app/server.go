package app

type IServer interface {
	SetRouter(router IRouter)
}

type IRouter interface {
	SetUseCase(usecase interface{})
	SetLogger(logger interface{})
}