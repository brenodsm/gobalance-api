package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/brenodsm/gobalance-api/internal/database"
	"github.com/brenodsm/gobalance-api/internal/response"
	"github.com/brenodsm/gobalance-api/internal/user/model"
	"github.com/brenodsm/gobalance-api/internal/user/repository"
	"github.com/brenodsm/gobalance-api/internal/user/service"
)

// HandlerCreateUser trata requisições HTTP para registrar um novo usuário.
func HandlerCreateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		response.WriteErrorJSON(w, http.StatusBadRequest, err)
		return
	}
	db, err := database.OpenConnection()
	if err != nil {
		response.WriteErrorJSON(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	userService := service.NewUserService()
	if err := userService.Register(&user); err != nil {
		response.WriteErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	repo := repository.NewUserRepository(db)
	id, err := repo.Create(user)
	if err != nil {
		if errors.Is(err, repository.ErrUserAlreadyRegister) {
			response.WriteErrorJSON(w, http.StatusInternalServerError, err)
			return
		}
		response.WriteErrorJSON(w, http.StatusConflict, err)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("/users/%d", id))
	response.WriteJSON(w, http.StatusCreated, id)
}
