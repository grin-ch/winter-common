package trace_util

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

func Pprof(port int) error {
	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
