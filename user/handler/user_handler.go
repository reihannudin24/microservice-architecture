package handler

import (
	"Blast/internal/user/dto"
	"Blast/internal/user/model"
	"Blast/internal/user/repository"
	"Blast/internal/user/service"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

type UserHandler struct {
	UserService *service.UserService
}

func NewUserHandler(db *sql.DB) *UserHandler {
	repo := repository.NewUserRepository(db)
	service := service.NewUserService(repo)
	return &UserHandler{UserService: service}
}

func (h *UserHandler) HandleUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var req dto.RegisterEmailRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		user := model.User{
			Email: req.Email,
		}

		id, err := h.UserService.RegisterEmail(user)
		if err != nil {
			fmt.Printf("Failed to create user: %v\n", err)

			http.Error(w, "Failed to create user ", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"id": id,
		})
	}
}
