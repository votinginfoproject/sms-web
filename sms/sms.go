package sms

import (
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"github.com/votinginfoproject/sms-web/queue"
)

var q queue.ExternalQueueService

func WireUp(eqs queue.ExternalQueueService) {
	q = eqs
	q.Connect()
}

func Receive(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	sid := req.Form["AccountSid"]

	if len(sid) == 0 || sid[0] != os.Getenv("TWILIO_SID") {
		res.WriteHeader(http.StatusForbidden)
		log.Printf("[FORBIDDEN] Method: %s - Path: %s - Host: %s, FormData: %s", req.Method, req.URL.RequestURI(), req.Host, req.Form)
	} else {
		number := req.Form["From"][0]
		message := req.Form["Body"][0]
		q.Enqueue(number, message)
	}
}
