package main

import (
	"go.uber.org/zap"
	"net/http"
)

// EchoHandler is an http.Handler that copies its request body
// back to the response.
type EchoHandler struct {
	log *zap.Logger
}

// NewEchoHandler builds a new EchoHandler.
func NewEchoHandler(log *zap.Logger) *EchoHandler {
	return &EchoHandler{log: log}
}

// 实现接口 implement http.Handler (ServeHTTP method)
// ServeHTTP handles an HTTP request to the /echo endpoint.
func (h *EchoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//if _, err := io.Copy(w, r.Body); err != nil {
	//	//fmt.Fprintln(os.Stderr, "Failed to handle request:", err)
	//	h.log.Warn("Failed to handle request:", zap.Error(err))
	//}

	w.Write([]byte("hello echo success"))

}

func (*EchoHandler) Pattern() string {
	return "/echo"
}
