package staff

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/Dashiiidzzze/admin_panel_for_rfid_security/internal/lib/api/response"
	"github.com/Dashiiidzzze/admin_panel_for_rfid_security/internal/storage/postgres"
)

type StaffCreator interface {
	CreateStaff(postgres.CreateStaffRequest) error
}

// создание карточки сотрудника
func Create(log *slog.Logger, storage StaffCreator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.staff.post.Create"

		log := log.With(
			slog.String("op", op),
		)

		log.Info(op + " получен запрос на создание записи сотрудника")

		var req postgres.CreateStaffRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			response.SendError(w, r, log, err, "не удалось декодировать тело запроса")
			return
		}

		err := storage.CreateStaff(req)
		if err != nil {
			response.SendError(w, r, log, err, "не удалось создать сотрудника")
			return
		}

		response.SendSuccess(w, r, nil)
	}
}
