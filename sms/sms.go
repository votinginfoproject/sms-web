package sms

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/votinginfoproject/sms-web/queue"
)

var q queue.ExternalQueueService

func WireUp(eqs queue.ExternalQueueService) {
	q = eqs
	q.Connect()
}

func Receive(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	err := req.ParseForm()
	if err != nil {
		log.Fatal("Failed to parse form data from twilio")
		return
	}

	number := req.Form["From"][0]
	message := req.Form["Body"][0]
	q.Enqueue(number, message)
}
