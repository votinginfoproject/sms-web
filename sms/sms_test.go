package sms_test

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/votinginfoproject/sms-web/routes"
)

type FakeQueue struct {
	mock.Mock
}

func (fq *FakeQueue) Enqueue(number string, message string) {
	fq.Mock.Called(number, message)
}

func (fq *FakeQueue) Connect() {
	fq.Mock.Called()
}

var (
	server *httptest.Server
)

func setup(fq *FakeQueue) {
	routes := routes.New(fq)
	server = httptest.NewServer(routes)
	log.SetOutput(ioutil.Discard)
}

func teardown() {
	server.Close()
}

func TestReceive(t *testing.T) {
	fq := new(FakeQueue)
	fq.On("Connect").Return()
	fq.On("Enqueue", "+15555555555", "this is a test").Return()

	setup(fq)
	defer teardown()

	data := url.Values{}
	data.Set("Body", "this is a test")
	data.Set("From", "+15555555555")

	res, _ := http.PostForm(server.URL+"/", data)

	assert.Equal(t, "200 OK", res.Status)

	fq.Mock.AssertExpectations(t)
}
