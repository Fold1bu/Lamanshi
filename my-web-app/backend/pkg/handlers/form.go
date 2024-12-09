

package handlers

import (
	"log"
	"my-go-backend/backend/pkg/models"
	"my-go-backend/backend/pkg/repository"
	"net/http"
	"time"
)

// Handler содержит зависимости для обработки запросов
type Handler struct {
	Repo *repository.UserRepository
}

// ServeForm отображает HTML-форму
func (h *Handler) ServeForm(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}


// SubmitForm обрабатывает отправку формы
func (h *Handler) SubmitForm(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		user := models.User{
			UserName:  r.FormValue("userName"),
			PhoneUser: r.FormValue("phoneUser"),
			Time:      r.FormValue("time"),
		}

		// Разбор даты
		dateStr := r.FormValue("date")
		parsedDate, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			log.Printf("Ошибка при разборе даты: %v\n", err) // Логируем ошибку
			http.Error(w, "Ошибка при разборе даты: "+err.Error(), http.StatusBadRequest)
			return
		}
		user.Date = parsedDate

		// Сохранение пользователя в базе данных
		err = h.Repo.SaveUser(user)
		if err != nil {
			log.Printf("Ошибка при сохранении данных пользователя: %v\n", err) // Логируем ошибку
			http.Error(w, "Ошибка при сохранении данных: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Перенаправление или успешный ответ
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
	}
}
