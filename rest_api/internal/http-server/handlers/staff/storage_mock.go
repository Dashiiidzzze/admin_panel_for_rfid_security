// staff/storage_mock.go
package staff

import (
	"github.com/Dashiiidzzze/admin_panel_for_rfid_security/internal/storage/postgres"
)

type MockStaffStorage struct {
	CreateStaffFunc     func(postgres.CreateStaffRequest) error
	GetStaffShortFunc   func() ([]postgres.StaffShort, error)
	GetStaffFunc        func(int) ([]postgres.Staff, error)
	UpdateStaffFunc     func(int, postgres.CreateStaffRequest) error
	DeleteStaffByIDFunc func(int) error
}

func (m *MockStaffStorage) CreateStaff(req postgres.CreateStaffRequest) error {
	return m.CreateStaffFunc(req)
}

func (m *MockStaffStorage) GetStaffShort() ([]postgres.StaffShort, error) {
	return m.GetStaffShortFunc()
}

func (m *MockStaffStorage) GetStaff(id int) ([]postgres.Staff, error) {
	return m.GetStaffFunc(id)
}

func (m *MockStaffStorage) UpdateStaff(id int, req postgres.CreateStaffRequest) error {
	return m.UpdateStaffFunc(id, req)
}

func (m *MockStaffStorage) DeleteStaffByID(id int) error {
	return m.DeleteStaffByIDFunc(id)
}
