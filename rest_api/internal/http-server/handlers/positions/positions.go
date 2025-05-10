package positions

import (
	"log/slog"
	"net/http"

	"github.com/Dashiiidzzze/admin_panel_for_rfid_security/internal/lib/api/response"
	"github.com/Dashiiidzzze/admin_panel_for_rfid_security/internal/storage/postgres"
)

type Response struct {
	Dolzhnosti []postgres.Dolzhnost `json:"dolzhnosti"`
}

type PositionsGetter interface {
	GetDolzhnosti() ([]postgres.Dolzhnost, error) // интерфейс взаимодействия с бд
}

// TODO: сделать тесты через моки

// New возвращает http.HandlerFunc, который обрабатывает GET-запрос на получение должностей
func New(log *slog.Logger, storage PositionsGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.positions.get.New"

		log := log.With(
			slog.String("op", op),
		)

		log.Info("получен запрос на список должностей")

		// Получаем должности из БД через интерфейс
		dolzhnosti, err := storage.GetDolzhnosti()
		if err != nil {
			response.SendError(w, r, log, err, "не удалось получить должности")
			return
		}

		response.SendSuccess(w, r, Response{
			Dolzhnosti: dolzhnosti,
		})
	}
}
