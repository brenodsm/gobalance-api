package router

import "github.com/go-chi/chi/v5"

func RunServer() (r *chi.Mux) {
	r = chi.NewRouter()
	ConfigRoute(r)
	return r
}
