package sms_test

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
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
	os.Setenv("TWILIO_SID", "test")
	routes := routes.New(fq, nil)
	server = httptest.NewServer(routes)
	log.SetOutput(ioutil.Discard)
}

func teardown() {
	os.Setenv("TWILIO_SID", "")
	server.Close()
}

func TestReceiveSuccess(t *testing.T) {
	fq := new(FakeQueue)
	fq.On("Connect").Return()
	fq.On("Enqueue", "+15555555555", "this is a test").Return()

	setup(fq)
	defer teardown()

	data := url.Values{}
	data.Set("Body", "this is a test")
	data.Set("From", "+15555555555")
	data.Set("AccountSid", "test")

	res, _ := http.PostForm(server.URL+"/", data)

	assert.Equal(t, "200 OK", res.Status)

	fq.Mock.AssertExpectations(t)
}

func TestReceiveError(t *testing.T) {
	fq := new(FakeQueue)
	fq.On("Connect").Return()

	setup(fq)
	defer teardown()

	data := url.Values{}
	data.Set("Body", "this is a test")
	data.Set("AccountSid", "test")

	res, _ := http.PostForm(server.URL+"/", data)

	assert.Equal(t, "500 Internal Server Error", res.Status)

	fq.Mock.AssertExpectations(t)
}

func TestReceiveForbiddenNoEnv(t *testing.T) {
	fq := new(FakeQueue)
	fq.On("Connect").Return()

	setup(fq)
	defer teardown()

	os.Setenv("TWILIO_SID", "")
	data := url.Values{}
	data.Set("Body", "this is a test")
	data.Set("From", "+15555555555")
	data.Set("AccountSid", "test")

	res, _ := http.PostForm(server.URL+"/", data)

	assert.Equal(t, "403 Forbidden", res.Status)

	fq.Mock.AssertExpectations(t)
}

func TestReceiveForbiddenNoFormData(t *testing.T) {
	fq := new(FakeQueue)
	fq.On("Connect").Return()

	setup(fq)
	defer teardown()

	data := url.Values{}
	data.Set("Body", "this is a test")
	data.Set("From", "+15555555555")

	res, _ := http.PostForm(server.URL+"/", data)

	assert.Equal(t, "403 Forbidden", res.Status)

	fq.Mock.AssertExpectations(t)
}
