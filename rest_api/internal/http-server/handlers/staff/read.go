package staff

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/Dashiiidzzze/admin_panel_for_rfid_security/internal/lib/api/response"
	"github.com/Dashiiidzzze/admin_panel_for_rfid_security/internal/storage/postgres"
	"github.com/go-chi/chi"
)

type GetResponseAll struct {
	Staff []postgres.StaffShort `json:"staff"` // Список должностей
}

type StaffShortGetter interface {
	GetStaffShort() ([]postgres.StaffShort, error) // интерфейс взаимодействия с бд
}

// TODO: сделать тесты через моки

// New возвращает http.HandlerFunc, который обрабатывает GET-запрос на получение должностей
func ReadAll(log *slog.Logger, storage StaffShortGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.staff.get.Read" // Название операции для логов

		// Обогащаем логгер доп. информацией: имя операции и ID запроса
		log := log.With(
			slog.String("op", op),
		)

		// Пишем в лог, что пришёл запрос
		log.Info("получен запрос на список сотрудников")

		// Получаем должности из БД через интерфейс
		staff, err := storage.GetStaffShort()
		if err != nil {
			response.SendError(w, r, log, err, "не удалось получить список сотрудников")
			return
		}

		response.SendSuccess(w, r, GetResponseAll{
			Staff: staff, // список должностей
		})
	}
}

type GetResponseCard struct {
	//response.Response                  // Встраиваем общую структуру ответа (например, статус)
	Staff []postgres.Staff `json:"staff"` // Список должностей
}

type StaffGetter interface {
	GetStaff(id int) ([]postgres.Staff, error) // интерфейс взаимодействия с бд
}

// New возвращает http.HandlerFunc, который обрабатывает GET-запрос на получение карточки сотрудника
func ReadCard(log *slog.Logger, storage StaffGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.staff.get.Read"

		log := log.With(
			slog.String("op", op),
		)

		log.Info(op + " получен запрос на карточку сотрудника")

		// Получаем id из URL
		idStr := chi.URLParam(r, "id")
		if idStr == "" {
			response.SendError(w, r, log, nil, "не передан id сотрудника")
			return
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			response.SendError(w, r, log, err, "неверный формат id")
			return
		}

		staff, err := storage.GetStaff(id)
		if err != nil {
			response.SendError(w, r, log, err, "не удалось получить данные сотрудника")
			return
		}

		if len(staff) == 0 {
			response.SendError(w, r, log, err, "сотрудник не найден")
			return
		}

		response.SendSuccess(w, r, GetResponseCard{
			Staff: staff,
		})
	}
}
