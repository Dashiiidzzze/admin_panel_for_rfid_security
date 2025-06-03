// staff/staff_test.go
package staff

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"log/slog"

	"github.com/Dashiiidzzze/admin_panel_for_rfid_security/internal/storage/postgres"
	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
)

func TestStaffHandlers(t *testing.T) {
	log := slog.Default()

	t.Run("Create Staff", func(t *testing.T) {
		tests := []struct {
			name           string
			requestBody    string
			mockBehavior   func(*MockStaffStorage)
			expectedStatus int
			expectedBody   string
		}{
			{
				name:        "Success",
				requestBody: `{"name":"John Doe","position":"Developer"}`,
				mockBehavior: func(m *MockStaffStorage) {
					m.CreateStaffFunc = func(req postgres.CreateStaffRequest) error {
						assert.Equal(t, "John Doe", req.Name)
						assert.Equal(t, "Developer", req.Position)
						return nil
					}
				},
				expectedStatus: http.StatusOK,
				expectedBody:   `{"status":"OK","data":null}`,
			},
			{
				name:        "Invalid JSON",
				requestBody: `{"name":"John Doe"`,
				mockBehavior: func(m *MockStaffStorage) {
					m.CreateStaffFunc = func(req postgres.CreateStaffRequest) error {
						return nil
					}
				},
				expectedStatus: http.StatusBadRequest,
				expectedBody:   `{"status":"error","error":"не удалось декодировать тело запроса"}`,
			},
			{
				name:        "Storage Error",
				requestBody: `{"name":"John Doe","position":"Developer"}`,
				mockBehavior: func(m *MockStaffStorage) {
					m.CreateStaffFunc = func(req postgres.CreateStaffRequest) error {
						return errors.New("storage error")
					}
				},
				expectedStatus: http.StatusInternalServerError,
				expectedBody:   `{"status":"error","error":"не удалось создать сотрудника"}`,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				mockStorage := &MockStaffStorage{}
				tt.mockBehavior(mockStorage)

				handler := Create(log, mockStorage)

				req := httptest.NewRequest("POST", "/items", bytes.NewBufferString(tt.requestBody))
				req.Header.Set("Content-Type", "application/json")

				rr := httptest.NewRecorder()
				handler.ServeHTTP(rr, req)

				assert.Equal(t, tt.expectedStatus, rr.Code)
				assert.JSONEq(t, tt.expectedBody, rr.Body.String())
			})
		}
	})

	t.Run("Read All Staff", func(t *testing.T) {
		mockStorage := &MockStaffStorage{
			GetStaffShortFunc: func() ([]postgres.StaffShort, error) {
				return []postgres.StaffShort{
					{ID: 1, Name: "John Doe"},
					{ID: 2, Name: "Jane Smith"},
				}, nil
			},
		}

		handler := ReadAll(log, mockStorage)

		req := httptest.NewRequest("GET", "/items", nil)
		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		expectedBody := `{
			"status": "OK",
			"data": {
				"staff": [
					{"id": 1, "name": "John Doe"},
					{"id": 2, "name": "Jane Smith"}
				]
			}
		}`
		assert.JSONEq(t, expectedBody, rr.Body.String())
	})

	t.Run("Read Staff Card", func(t *testing.T) {
		tests := []struct {
			name           string
			id             string
			mockBehavior   func(*MockStaffStorage)
			expectedStatus int
			expectedBody   string
		}{
			{
				name: "Success",
				id:   "1",
				mockBehavior: func(m *MockStaffStorage) {
					m.GetStaffFunc = func(id int) ([]postgres.Staff, error) {
						assert.Equal(t, 1, id)
						return []postgres.Staff{
							{ID: 1, Name: "John Doe", Position: "Developer"},
						}, nil
					}
				},
				expectedStatus: http.StatusOK,
				expectedBody: `{
					"status": "OK",
					"data": {
						"staff": [
							{"id": 1, "name": "John Doe", "position": "Developer"}
						]
					}
				}`,
			},
			{
				name:           "Invalid ID",
				id:             "abc",
				mockBehavior:   func(m *MockStaffStorage) {},
				expectedStatus: http.StatusBadRequest,
				expectedBody:   `{"status":"error","error":"неверный формат id"}`,
			},
			{
				name: "Not Found",
				id:   "999",
				mockBehavior: func(m *MockStaffStorage) {
					m.GetStaffFunc = func(id int) ([]postgres.Staff, error) {
						return nil, nil
					}
				},
				expectedStatus: http.StatusNotFound,
				expectedBody:   `{"status":"error","error":"сотрудник не найден"}`,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				mockStorage := &MockStaffStorage{}
				tt.mockBehavior(mockStorage)

				handler := ReadCard(log, mockStorage)

				r := chi.NewRouter()
				r.Get("/items/{id}", handler)

				req := httptest.NewRequest("GET", "/items/"+tt.id, nil)
				rr := httptest.NewRecorder()
				r.ServeHTTP(rr, req)

				assert.Equal(t, tt.expectedStatus, rr.Code)
				assert.JSONEq(t, tt.expectedBody, rr.Body.String())
			})
		}
	})

	t.Run("Update Staff", func(t *testing.T) {
		tests := []struct {
			name           string
			id             string
			requestBody    string
			mockBehavior   func(*MockStaffStorage)
			expectedStatus int
			expectedBody   string
		}{
			{
				name:        "Success",
				id:          "1",
				requestBody: `{"name":"John Updated","position":"Senior Developer"}`,
				mockBehavior: func(m *MockStaffStorage) {
					m.UpdateStaffFunc = func(id int, req postgres.CreateStaffRequest) error {
						assert.Equal(t, 1, id)
						assert.Equal(t, "John Updated", req.Name)
						assert.Equal(t, "Senior Developer", req.Position)
						return nil
					}
				},
				expectedStatus: http.StatusOK,
				expectedBody:   `{"status":"OK","data":null}`,
			},
			{
				name:           "Invalid ID",
				id:             "abc",
				requestBody:    `{"name":"John Updated"}`,
				mockBehavior:   func(m *MockStaffStorage) {},
				expectedStatus: http.StatusBadRequest,
				expectedBody:   `{"status":"error","error":"неверный формат id"}`,
			},
			{
				name:           "Invalid JSON",
				id:             "1",
				requestBody:    `{"name":"John Updated"`,
				mockBehavior:   func(m *MockStaffStorage) {},
				expectedStatus: http.StatusBadRequest,
				expectedBody:   `{"status":"error","error":"invalid request body"}`,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				mockStorage := &MockStaffStorage{}
				tt.mockBehavior(mockStorage)

				handler := Update(log, mockStorage)

				r := chi.NewRouter()
				r.Put("/items/{id}", handler)

				req := httptest.NewRequest("PUT", "/items/"+tt.id, bytes.NewBufferString(tt.requestBody))
				req.Header.Set("Content-Type", "application/json")

				rr := httptest.NewRecorder()
				r.ServeHTTP(rr, req)

				assert.Equal(t, tt.expectedStatus, rr.Code)
				assert.JSONEq(t, tt.expectedBody, rr.Body.String())
			})
		}
	})

	t.Run("Delete Staff", func(t *testing.T) {
		tests := []struct {
			name           string
			id             string
			mockBehavior   func(*MockStaffStorage)
			expectedStatus int
			expectedBody   string
		}{
			{
				name: "Success",
				id:   "1",
				mockBehavior: func(m *MockStaffStorage) {
					m.DeleteStaffByIDFunc = func(id int) error {
						assert.Equal(t, 1, id)
						return nil
					}
				},
				expectedStatus: http.StatusOK,
				expectedBody:   `{"status":"OK","data":null}`,
			},
			{
				name:           "Invalid ID",
				id:             "abc",
				mockBehavior:   func(m *MockStaffStorage) {},
				expectedStatus: http.StatusBadRequest,
				expectedBody:   `{"status":"error","error":"Неверный ID"}`,
			},
			{
				name: "Storage Error",
				id:   "1",
				mockBehavior: func(m *MockStaffStorage) {
					m.DeleteStaffByIDFunc = func(id int) error {
						return errors.New("storage error")
					}
				},
				expectedStatus: http.StatusInternalServerError,
				expectedBody:   `{"status":"error","error":"не удалось удалить сотрудника"}`,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				mockStorage := &MockStaffStorage{}
				tt.mockBehavior(mockStorage)

				handler := Delete(log, mockStorage)

				r := chi.NewRouter()
				r.Delete("/items/{id}", handler)

				req := httptest.NewRequest("DELETE", "/items/"+tt.id, nil)
				rr := httptest.NewRecorder()
				r.ServeHTTP(rr, req)

				assert.Equal(t, tt.expectedStatus, rr.Code)
				assert.JSONEq(t, tt.expectedBody, rr.Body.String())
			})
		}
	})
}
