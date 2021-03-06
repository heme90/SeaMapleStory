package protocol

import (
	"SeaMaple/dto"
	"SeaMaple/dto/static"
	"net/http"
)

//에러 처리 룰은 다음과 같음
//유즈케이스에서 핸들링 가능한 경우에는 그냥 처리
//에러가 발생한 경우 에러 객체를 올리고 올려서 라우터에 내장된 logger로 찍음
type Router struct {
	router *http.Server
	usecase dto.Usecase
	logger static.ILogger
}

func NewRouter(http *http.Server, logger static.ILogger) *Router {
	return &Router{
		router:  http,
		logger:  logger,
	}
}

func (r *Router) Router() *http.Server {
	return r.router
}

func (r *Router) SetRouter(router *http.Server) {
	r.router = router
}

func (r *Router) RegisterUseCase(u dto.Usecase) {

}

func (r *Router) SetLogger(logger static.ILogger) {
	r.logger = logger
}

func (r *Router) GetLogger() static.ILogger {
	return r.logger
}

func (r *Router) Run() {
	if err := r.router.ListenAndServe();err != nil {
		r.router.Shutdown(nil)
	}
}