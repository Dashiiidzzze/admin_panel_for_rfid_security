package staff

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/Dashiiidzzze/admin_panel_for_rfid_security/internal/lib/api/response"
	"github.com/Dashiiidzzze/admin_panel_for_rfid_security/internal/storage/postgres"
	"github.com/go-chi/chi"
)

type StaffUpdator interface {
	UpdateStaff(int, postgres.CreateStaffRequest) error
}

// обновляет информацию о сотруднике по его ID
func Update(log *slog.Logger, storage StaffUpdator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.staff.Update"

		log := log.With(
			slog.String("op", op),
		)

		log.Info(op + "получен запрос на обновление карточки сотрудника")

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

		// Читаем тело запроса
		var req postgres.CreateStaffRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			response.SendError(w, r, log, err, "invalid request body")
			return
		}

		// Вызываем логику обновления
		err = storage.UpdateStaff(id, req)
		if err != nil {
			response.SendError(w, r, log, err, "failed to update staff")
			return
		}

		response.SendSuccess(w, r, nil)
	}
}
