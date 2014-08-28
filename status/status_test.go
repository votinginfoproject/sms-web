package status_test

import (
	"fmt"
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

func TestGet(t *testing.T) {
	setup()
	defer teardown()

	res, err := http.Get(server.URL + "/status")

	if err != nil {
		fmt.Println(err)
	}

	body, _ := ioutil.ReadAll(res.Body)
	assert.Equal(t, "OK\n", string(body))

	assert.Equal(t, "text/html", res.Header["Content-Type"][0])

	assert.Equal(t, "200 OK", res.Status)
}
