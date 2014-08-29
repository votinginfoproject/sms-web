package main

import (
	"log"
	"net/http"
	"os"
	"runtime"
	"strconv"

	"github.com/votinginfoproject/sms-web/env"
	"github.com/votinginfoproject/sms-web/queue"
	"github.com/votinginfoproject/sms-web/routes"
)

func main() {
	env.Load()

	procs, err := strconv.Atoi(os.Getenv("PROCS"))
	if err != nil {
		log.Fatal("[ERROR] you must specify procs in the .env file")
	}
	runtime.GOMAXPROCS(procs)

	q := queue.New()
	log.Panic(http.ListenAndServe(":8080", routes.New(q)))
}
