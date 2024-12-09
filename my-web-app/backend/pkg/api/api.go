package api

import (
	"my-go-backend/backend/pkg/handlers"
	"my-go-backend/backend/pkg/repository"
	"net/http"
)

// SetupRoutes настраивает маршруты для API
func SetupRoutes(repo *repository.UserRepository) {
	h := handlers.Handler{Repo: repo}
	http.HandleFunc("/submit", h.SubmitForm)
}
