package server

import "SeaMaple/app"

func init() {

}

type Server struct {
	router interface{}
	Run func()
	Stop func()
	OnPanic func()
}

func (s *Server) SetRouter(router app.IRouter) {
	s.router = router
}

func NewMRNGServer() app.IServer {
	return &Server{}
}

func (s *Server) InitServer () {
	
}

func (s *Server) GraceFullStop () {

}


