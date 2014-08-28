package routes

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/votinginfoproject/sms-web/status"
)

type Server struct {
	handler http.Handler
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.handler.ServeHTTP(w, r)
}

func New() *Server {
	routes := httprouter.New()

	routes.GET("/status", status.Get)

	return &Server{routes}
}
