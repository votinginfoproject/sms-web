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

	els.logger.Send(loggly.Message{"message": message[:len(message)-1]})

	if os.Getenv("ENVIRONMENT") == "development" {
		fmt.Print(message)
	}

	return len(p), nil
}
