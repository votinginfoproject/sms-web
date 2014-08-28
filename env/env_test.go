package env

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setup() {
	os.Remove(".env")
	ioutil.WriteFile(".env", []byte("ENVIRONMENT=DEVELOPMENT\n"), 0644)
}

func teardown() {
	os.Remove(".env")
}

func TestLoadEnvFile(t *testing.T) {
	setup()
	defer teardown()

	Load()
	assert.Equal(t, "DEVELOPMENT", os.Getenv("ENVIRONMENT"))
}
