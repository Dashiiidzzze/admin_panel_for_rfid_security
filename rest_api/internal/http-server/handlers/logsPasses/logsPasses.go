package logsPasses

import (
	"log/slog"
	"net/http"

	"github.com/Dashiiidzzze/admin_panel_for_rfid_security/internal/lib/api/response"
	"github.com/Dashiiidzzze/admin_panel_for_rfid_security/internal/storage/postgres"
)

type Response struct {
	Logs []postgres.LogsPases `json:"logspasses"`
}

type LogsPasesGetter interface {
	GetLogsPases() ([]postgres.LogsPases, error) // интерфейс взаимодействия с бд
}

// TODO: сделать тесты через моки

// возвращает http.HandlerFunc, который обрабатывает GET-запрос на получение должностей
func New(log *slog.Logger, storage LogsPasesGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.LogsPases.get.New"

		log := log.With(
			slog.String("op", op),
		)

		log.Info("получен запрос на список логов проходов")

		// Получаем должности из БД
		logsPases, err := storage.GetLogsPases()
		if err != nil {
			response.SendError(w, r, log, err, "не удалось получить логи проходов")
			return
		}

		response.SendSuccess(w, r, Response{
			Logs: logsPases,
		})
	}
}
