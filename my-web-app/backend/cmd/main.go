

package main

import (
	"context"
	"log"
	"my-go-backend/backend/pkg/api"
	"my-go-backend/backend/pkg/repository"
	"net/http"
	"path/filepath"

	"github.com/jackc/pgx/v4/pgxpool"
)

const connStr = "postgres://postgres:1258@localhost:5432/myTestSite"

func main() {
	// Подключаемся к базе данных
	dbPool, err := pgxpool.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}
	defer dbPool.Close()

	// Настраиваем маршруты API
	userRepo := repository.NewUserRepository(dbPool)
	api.SetupRoutes(userRepo)

	// Указываем путь к статическим файлам (HTML, CSS, JS)
	staticDir := filepath.Join("..", "..", "frontend")
	http.Handle("/", http.FileServer(http.Dir(staticDir)))

	// Запускаем сервер
	log.Println("Сервер запущен на порту 8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Ошибка запуска сервера: ", err)
	}
}
