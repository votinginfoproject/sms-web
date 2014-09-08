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
	"github.com/yvasiyarov/gorelic"
)

func main() {
	env.Load()

	host, _ := os.Hostname()

	var agent *gorelic.Agent
	if os.Getenv("ENVIRONMENT") == "production" {
		agent = gorelic.NewAgent()
		agent.NewrelicName = "sms-web" + "-" + host
		agent.NewrelicLicense = os.Getenv("NEWRELIC_TOKEN")
		agent.CollectHTTPStat = true
		agent.NewrelicPollInterval = 15
		agent.Run()
	}

	procs, err := strconv.Atoi(os.Getenv("PROCS"))
	if err != nil {
		log.Fatal("[ERROR] you must specify procs in the .env file")
	}
	runtime.GOMAXPROCS(procs)

	log.SetOutput(logger.New())

	q := queue.New()
	log.Print("[INFO] starting server on port 8080")
	log.Panic(http.ListenAndServe(":8080", routes.New(q, agent)))
}
