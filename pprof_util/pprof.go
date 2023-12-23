package pprof_util

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

func Pprof(port int) error {
	if port <= 0 {
		return nil
	}
	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
