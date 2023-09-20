package api

import "net/http"

type Server struct {
	h *http.Handler
}