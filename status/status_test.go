package status_test

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/votinginfoproject/sms-web/routes"
)

var (
	server *httptest.Server
)

func setup() {
	routes := routes.New()
	server = httptest.NewServer(routes)
	log.SetOutput(ioutil.Discard)
}

func teardown() {
	server.Close()
}

func TestGetStatus(t *testing.T) {
	setup()
	defer teardown()

	res, _ := http.Get(server.URL + "/status")

	body, _ := ioutil.ReadAll(res.Body)
	assert.Equal(t, "OK\n", string(body))

	assert.Equal(t, "text/html", res.Header["Content-Type"][0])

	assert.Equal(t, "200 OK", res.Status)
}
