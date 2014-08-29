package routes

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/votinginfoproject/sms-web/queue"
	"github.com/votinginfoproject/sms-web/sms"
	"github.com/votinginfoproject/sms-web/status"
)

type Server struct {
	handler http.Handler
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.handler.ServeHTTP(w, r)
}

func New(q queue.ExternalQueueService) *Server {
	routes := httprouter.New()

	routes.GET("/status", status.Get)

	if q != nil {
		sms.WireUp(q)
		q.Connect()
	}
	routes.POST("/", sms.Receive)

	return &Server{routes}
}
