package requestlog

import "net/http"

type Handler struct {
	handler http.Handler
}

func NewHandler(h http.HandlerFunc) *Handler {
	return &Handler{
		handler: h,
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.handler.ServeHTTP(w, r)
}