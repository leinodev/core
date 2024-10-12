package request

import "net/http"

type HttpContext struct {
	R *http.Request
	W http.ResponseWriter
}
