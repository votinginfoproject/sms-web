package main

import (
	"net/http"

	"github.com/votinginfoproject/sms-web/env"
	"github.com/votinginfoproject/sms-web/queue"
	"github.com/votinginfoproject/sms-web/routes"
)

func main() {
	env.Load()

	q := queue.New()
	panic(http.ListenAndServe(":8080", routes.New(q)))
}
