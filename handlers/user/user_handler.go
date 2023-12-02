package user_handler

import (
	"encoding/json"
	user_service "go-testing/service/user"
	"net/http"
)

type UserHandler struct {
	service *user_service.UserService
}

func NewUserHandler(service *user_service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	user, err := h.service.GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}
