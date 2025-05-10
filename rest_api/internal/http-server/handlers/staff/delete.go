package staff

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/Dashiiidzzze/admin_panel_for_rfid_security/internal/lib/api/response"
	"github.com/go-chi/chi"
)

type StaffDeleter interface {
	DeleteStaffByID(id int) error // интерфейс взаимодействия с бд
}

// TODO: сделать тесты через моки

// New возвращает http.HandlerFunc, который обрабатывает запрос на удаление должностей
func Delete(log *slog.Logger, storage StaffDeleter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.staff.get.Read"

		log := log.With(
			slog.String("op", op),
		)

		log.Info(op + " получен запрос на удаление сотрудника")

		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			response.SendError(w, r, log, err, "Неверный ID")
			return
		}

		// Получаем должности из БД через интерфейс
		err = storage.DeleteStaffByID(id)
		if err != nil {
			response.SendError(w, r, log, err, "не удалось удалить сотрудника")
			return
		}

		response.SendSuccess(w, r, nil)
	}
}
