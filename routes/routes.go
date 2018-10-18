package routes

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/votinginfoproject/sms-web/queue"
	"github.com/votinginfoproject/sms-web/sms"
	"github.com/votinginfoproject/sms-web/status"
	"github.com/yvasiyarov/gorelic"
)

type Server struct {
	handler http.Handler
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	log.Printf("[INFO] [REQUEST] Method: %s - Path: %s - Host: %s - Remote: %s - FormData: %s", r.Method, r.URL.RequestURI(), r.Host, r.RemoteAddr, r.Form)

	s.handler.ServeHTTP(w, r)

	log.Printf("[INFO] [REQUEST COMPLETED] Method: %s - Path: %s - Host: %s - Remote: %s - FormData: %s", r.Method, r.URL.RequestURI(), r.Host, r.RemoteAddr, r.Form)
}

func New(q queue.ExternalQueueService, agent *gorelic.Agent) *Server {
	routes := httprouter.New()

	routes.PanicHandler = func(res http.ResponseWriter, req *http.Request, _ interface{}) {
		res.WriteHeader(http.StatusInternalServerError)
		log.Print("[ERROR] : ", req)
	}

	routes.GET("/status", status.Get)

	if q != nil {
		sms.WireUp(q)
	}
	routes.POST("/", sms.Receive)

	if agent != nil {
		return &Server{agent.WrapHTTPHandler(routes)}
	} else {
		return &Server{routes}
	}
}
