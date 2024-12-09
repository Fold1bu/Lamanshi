package repository

import (
	"context"
	"fmt"
	"log"
	"my-go-backend/backend/pkg/models"

	"github.com/jackc/pgx/v4/pgxpool"
)

// UserRepository определяет методы для работы с пользователями
type UserRepository struct {
	db *pgxpool.Pool
}

// NewUserRepository создает новый экземпляр UserRepository
func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

// SaveUser сохраняет пользователя в базе данных
func (repo *UserRepository) SaveUser(user models.User) error {
	datetime := user.Date.Format("2006-01-02 15:04:05")
	_, err := repo.db.Exec(context.Background(), "INSERT INTO users (user_name, user_phone, time_slot, date_time) VALUES ($1, $2, $3, $4)", user.UserName, user.PhoneUser, user.Time, datetime)
	if err != nil {
		log.Printf("SQL ошибка: %v\n", err)                           // Логируем ошибку SQL
		return fmt.Errorf("не удалось выполнить SQL-запрос: %w", err) // Возвращаем более детальное сообщение
	}
	return nil
}
