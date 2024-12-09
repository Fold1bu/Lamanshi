package models

import "time"

// User представляет модель пользователя
type User struct {
	UserName  string
	PhoneUser string
	Time      string    // Проверьте, этот тип данных соответствует тому, что вы ожидаете
	Date      time.Time // Дата, которую мы получаем из формы
}
