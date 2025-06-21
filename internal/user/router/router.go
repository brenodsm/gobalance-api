package router

import (
	"github.com/brenodsm/gobalance-api/internal/user/handler"
	"github.com/go-chi/chi/v5"
)

// UserRoutes monta todas as rotas relacionadas a usuários no roteador fornecido.
func UserRoute(r chi.Router) {
	r.Route("/users", func(r chi.Router) {
		r.Post("/", handler.HandlerCreateUser)
	})
}
