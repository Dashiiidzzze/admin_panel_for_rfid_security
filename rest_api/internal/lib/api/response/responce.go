package response

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/go-chi/render"
)

// стандартная структура ответа с ошибкой
type ErrorResponse struct {
	Error      string `json:"error"`       // Сообщение для клиента
	StatusCode int    `json:"status_code"` // HTTP-код ошибки
}

// универсальная функция для отправки ошибок
func SendError(w http.ResponseWriter, r *http.Request, log *slog.Logger, err error, customMsg ...string) {
	var statusCode int
	msg := "Internal server error"

	// Определяем статус-код по типу ошибки
	switch {
	case errors.Is(err, ErrNotFound):
		statusCode = http.StatusNotFound
		msg = "Not found"
	case errors.Is(err, ErrUnauthorized):
		statusCode = http.StatusUnauthorized
		msg = "Unauthorized"
	case errors.Is(err, ErrBadRequest):
		statusCode = http.StatusBadRequest
		msg = "Bad request"
	default:
		statusCode = http.StatusInternalServerError
	}

	// Переопределяем сообщение, если передано кастомное
	if len(customMsg) > 0 {
		msg = customMsg[0]
	}

	// Логируем ошибку
	log.Error(msg,
		slog.Int("status", statusCode),
		slog.String("err", err.Error()),
	)

	// Отправляем ответ
	render.Status(r, statusCode)
	render.JSON(w, r, ErrorResponse{
		Error:      msg,
		StatusCode: statusCode,
	})
}

// Предопределенные ошибки
var (
	ErrNotFound     = errors.New("not found")
	ErrUnauthorized = errors.New("unauthorized")
	ErrBadRequest   = errors.New("bad request")
)

// универсальная функция для успешных ответов
func SendSuccess(w http.ResponseWriter, r *http.Request, data interface{}) {
	render.Status(r, http.StatusOK)
	render.JSON(w, r, data)
}
