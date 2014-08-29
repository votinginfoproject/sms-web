package status

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Get(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	res.Header().Set("Content-Type", "text/plain")
	fmt.Fprintln(res, "OK")
}
