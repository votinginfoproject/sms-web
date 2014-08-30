package main

import . "github.com/visionmedia/go-debug"
import "time"
import "fmt"
import "os"

var debug = Debug("single")

func main() {
	fmt.Printf("pid: %d\n", os.Getpid())

	for {
		debug("sending mail")
		debug("send email to %s", "tobi@segment.io")
		debug("send email to %s", "loki@segment.io")
		debug("send email to %s", "jane@segment.io")
		time.Sleep(500 * time.Millisecond)
	}
}
