package queue

import (
	"encoding/json"
	"os"

	"github.com/crowdmob/goamz/aws"
	"github.com/crowdmob/goamz/sqs"
)

type ExternalQueueService interface {
	Enqueue(number string, message string)
	Connect()
}

type SQS struct {
	q *sqs.Queue
}

type Data struct {
	Number  string `json:"number"`
	Message string `json:"message"`
}

func New() *SQS {
	return &SQS{nil}
}

func (s *SQS) Connect() {
	accessKey := os.Getenv("ACCESS_KEY_ID")
	secretKey := os.Getenv("SECRET_ACCESS_KEY")

	auth := aws.Auth{AccessKey: accessKey, SecretKey: secretKey}
	sqs := sqs.New(auth, aws.USEast)

	queue, err := sqs.GetQueue("vip-sms-development")
	if err != nil {
		panic(err)
	}

	s.q = queue
}

func (s *SQS) Enqueue(number string, message string) {
	data, err := json.Marshal(Data{number, message})

	if err != nil {
		panic(err)
	}

	s.q.SendMessage(string(data))
}
