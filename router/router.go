package router

import (
	"go-testing/handlers"
	"net/http"
)

func SetupRoutes(userHandler *handlers.UserHandler) {
	http.HandleFunc("/users", userHandler.GetUsers)
}
