package main

import (
	"log"
	"net/http"
	"os"
	"runtime"
	"strconv"

	"github.com/votinginfoproject/sms-web/env"
	"github.com/votinginfoproject/sms-web/logger"
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

	log.SetOutput(logger.New())

	q := queue.New()
	log.Print("[INFO] starting server on port 8080")
	log.Panic(http.ListenAndServe(":8080", routes.New(q)))
}
