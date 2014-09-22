package logger

import (
	"fmt"
	"os"

	"github.com/segmentio/go-loggly"
)

type ExternalLoggerService struct {
	logger *loggly.Client
}

func New() *ExternalLoggerService {
	logger := loggly.New(os.Getenv("LOGGLY_TOKEN"))
	logger.Tag("sms-web")
	logger.Tag(os.Getenv("ENVIRONMENT"))

	return &ExternalLoggerService{logger}
}

func (els *ExternalLoggerService) Write(p []byte) (n int, err error) {
	message := string(p)
	fmt.Print(message)

	if os.Getenv("ENVIRONMENT") != "development" {
		els.logger.Send(loggly.Message{"message": message[:len(message)-1]})
	}

	return len(p), nil
}
