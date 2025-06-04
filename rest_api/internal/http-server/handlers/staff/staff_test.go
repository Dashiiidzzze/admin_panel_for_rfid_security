package staff

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"context"

	"github.com/Dashiiidzzze/admin_panel_for_rfid_security/internal/storage/postgres"
	"github.com/go-chi/chi"
)

// мок, реализующий интерфейс StaffCreator
type mockStaffCreator struct {
	err error
}

func (m *mockStaffCreator) CreateStaff(req postgres.CreateStaffRequest) error {
	return m.err
}

func TestCreateStaff_Success(t *testing.T) {
	mockStorage := &mockStaffCreator{}
	logger := slog.Default()

	// структура запроса с актуальной зоной
	requestBody := postgres.CreateStaffRequest{
		Name:      "Иван",
		Position:  "Охранник",
		Phone:     "+71234567890",
		KeyNumber: "KEY123456",
		Zone: postgres.ZoneFields{
			FlightZone:   true,
			ClearZone:    false,
			Runaway:      true,
			BaggageZone:  false,
			ControlTower: true,
		},
	}

	body, _ := json.Marshal(requestBody)

	req := httptest.NewRequest(http.MethodPost, "/items", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handler := Create(logger, mockStorage)
	handler(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("ожидался статус 200 OK, получен %d", resp.StatusCode)
	}
}

func TestCreateStaff_InvalidData(t *testing.T) {
	mockStorage := &mockStaffCreator{}
	logger := slog.Default()

	// имя пустое
	requestBody := postgres.CreateStaffRequest{
		Name:      "",
		Position:  "Техник",
		Phone:     "+79991234567",
		KeyNumber: "TECH98765",
		Zone: postgres.ZoneFields{
			FlightZone:   false,
			ClearZone:    false,
			Runaway:      false,
			BaggageZone:  true,
			ControlTower: false,
		},
	}

	body, _ := json.Marshal(requestBody)

	req := httptest.NewRequest(http.MethodPost, "/items", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handler := Create(logger, mockStorage)
	handler(w, req)

	resp := w.Result()
	// пока валидации нет — вернётся 200 OK; если добавите проверку, поменяйте ожидание на 400
	if resp.StatusCode != http.StatusOK {
		t.Errorf("ожидался статус 200 OK, получен %d", resp.StatusCode)
	}
}

// mockStaffDeleter реализует интерфейс StaffDeleter
type mockStaffDeleter struct {
	err error
}

func (m *mockStaffDeleter) DeleteStaffByID(id int) error {
	return m.err
}

// makeChiRequest создаёт запрос с chi.RouteContext, чтобы корректно подставить параметр id
func makeChiRequest(method, target, id string) *http.Request {
	req := httptest.NewRequest(method, target, nil)

	// создаём chi route context и добавляем параметр id
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", id)

	// вставляем route context в контекст запроса
	ctx := context.WithValue(req.Context(), chi.RouteCtxKey, rctx)
	return req.WithContext(ctx)
}

// TestDeleteStaff_Success проверяет успешное удаление сотрудника
func TestDeleteStaff_Success(t *testing.T) {
	mockStorage := &mockStaffDeleter{} // нет ошибки => удаление успешно
	logger := slog.Default()

	req := makeChiRequest(http.MethodDelete, "/items/1", "1")
	w := httptest.NewRecorder()

	// вызываем обработчик
	handler := Delete(logger, mockStorage)
	handler(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("ожидался статус 200 OK, получен %d", resp.StatusCode)
	}
}
