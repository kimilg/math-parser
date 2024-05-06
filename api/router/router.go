package router

import (
	"github.com/go-chi/chi/v5"
	"math-parser/api/resource/health"
	"math-parser/api/resource/math"
	"math-parser/api/router/requestlog"
	"net/http"
)

func New(repository *math.Repository) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/livez", health.Read)

	r.Route("/v1", func(r chi.Router) {
		
		mathAPI := math.New(repository)
		r.Method(http.MethodPost, "/parse", requestlog.NewHandler(mathAPI.Parse))
	})

	return r
}
