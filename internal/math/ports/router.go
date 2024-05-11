package ports

import (
	"github.com/go-chi/chi/v5"
	"math-parser/internal/common/health"
)

func New(httpServer HttpServer) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/livez", health.Read)

	r.Route("/v1", func(r chi.Router) {
		r.Post("/parse", httpServer.Parse)
	})

	return r
}
