package postgres

import (
	"database/sql"
	"fmt"
	"strings"
)

// Структура для представления сотрудника с ФИО и ID
type StaffShort struct {
	ID   int    `json:"id"`
	Name string `json:"name"` // ФИО
}

// Получает id и ФИО сотрудников
func (s *Storage) GetStaffShort() ([]StaffShort, error) {
	const op = "storage.postgres.GetStaffShort"

	const query = `
		SELECT s.id, 
		       pd.familiya || ' ' || pd.imya || ' ' || COALESCE(pd.otchestvo, '') AS name
		FROM staff s
		JOIN personalnye_dannye pd ON s.id_pers_dannykh = pd.id
	`

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("%s: query failed: %w", op, err)
	}
	defer rows.Close()

	var result []StaffShort
	for rows.Next() {
		var sotr StaffShort
		if err := rows.Scan(&sotr.ID, &sotr.Name); err != nil {
			return nil, fmt.Errorf("%s: row scan failed: %w", op, err)
		}
		result = append(result, sotr)
	}

	return result, nil
}

// Удаляет сотрудника по ID
func (s *Storage) DeleteStaffByID(id int) error {
	const op = "storage.postgres.DeleteSotrudnikByID"

	const query = `DELETE FROM staff WHERE id = $1`

	_, err := s.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("%s: delete failed: %w", op, err)
	}

	return nil
}

// Структура для представления сотрудника
type Staff struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Position  string `json:"position"`
	Phone     string `json:"phone"`
	KeyNumber string `json:"keyNumber"`
	Zone      Zone   `json:"zone"`
}

type Zone struct {
	FlightZone   bool `json:"flightZone"`
	ClearZone    bool `json:"clearZone"`
	Runaway      bool `json:"runaway"`
	BaggageZone  bool `json:"baggageZone"`
	ControlTower bool `json:"controlTower"`
}

func (s *Storage) GetStaff(id int) ([]Staff, error) {
	const op = "storage.postgres.GetStaff"

	const query = `
		SELECT
			s.id AS staff_id,
			pd.familiya || ' ' || pd.imya || ' ' || COALESCE(pd.otchestvo, '') AS full_name,
			d.dolzhnost AS position,
			pd.nomer_telefona AS phone,
			s.nomer_klyucha AS key_number,
			-- Каждая зона в виде булевого признака: true, если есть, иначе false
			bool_or(CASE WHEN z.name = 'Зона A' THEN true ELSE false END) AS flight_zone,
			bool_or(CASE WHEN z.name = 'Зона B' THEN true ELSE false END) AS clear_zone,
			bool_or(CASE WHEN z.name = 'Зона C' THEN true ELSE false END) AS runaway,
			bool_or(CASE WHEN z.name = 'Зона D' THEN true ELSE false END) AS baggage_zone,
			bool_or(CASE WHEN z.name = 'Зона E' THEN true ELSE false END) AS control_tower
		FROM staff s
		JOIN personalnye_dannye pd ON s.id_pers_dannykh = pd.id
		LEFT JOIN dolzhnosti d ON s.id_dolzhnosti = d.id
		LEFT JOIN staff_zones sz ON sz.staff_id = s.id
		LEFT JOIN zony z ON sz.zona_id = z.id
		WHERE s.id = $1
		GROUP BY s.id, pd.familiya, pd.imya, pd.otchestvo, pd.nomer_telefona, d.dolzhnost, s.nomer_klyucha;
	`

	rows, err := s.db.Query(query, id)
	if err != nil {
		return nil, fmt.Errorf("%s: query failed: %w", op, err)
	}
	defer rows.Close()

	var result []Staff
	for rows.Next() {
		var sotr Staff
		var zone Zone

		err := rows.Scan(
			&sotr.ID,
			&sotr.Name,
			&sotr.Position,
			&sotr.Phone,
			&sotr.KeyNumber,
			&zone.FlightZone,
			&zone.ClearZone,
			&zone.Runaway,
			&zone.BaggageZone,
			&zone.ControlTower,
		)
		if err != nil {
			return nil, fmt.Errorf("%s: row scan failed: %w", op, err)
		}

		sotr.Zone = zone
		result = append(result, sotr)
	}

	return result, nil
}

// Структура запроса на создание сотрудника
type CreateStaffRequest struct {
	Name      string     `json:"name"`
	Position  string     `json:"position"`
	Phone     string     `json:"phone"`
	KeyNumber string     `json:"keyNumber"`
	Zone      ZoneFields `json:"zone"`
}

// Структура зон доступа
type ZoneFields struct {
	FlightZone   bool `json:"flightZone"`
	ClearZone    bool `json:"clearZone"`
	Runaway      bool `json:"runaway"`
	BaggageZone  bool `json:"baggageZone"`
	ControlTower bool `json:"controlTower"`
}

// CreateStaff сохраняет нового сотрудника и его зоны доступа
func (s *Storage) CreateStaff(req CreateStaffRequest) error {
	const op = "storage.postgres.CreateStaff"
	tx, err := s.db.Begin()
	if err != nil {
		return fmt.Errorf("%s: begin tx: %w", op, err)
	}
	defer tx.Rollback()

	// Разделим ФИО
	parts := strings.Fields(req.Name)
	if len(parts) < 2 {
		return fmt.Errorf("%s: имя должно содержать хотя бы фамилию и имя", op)
	}

	lastName := parts[0]
	firstName := parts[1]
	var middleName sql.NullString
	if len(parts) > 2 {
		middleName = sql.NullString{String: parts[2], Valid: true}
	} else {
		middleName = sql.NullString{Valid: false}
	}

	var personID int
	err = tx.QueryRow(`
		INSERT INTO personalnye_dannye (familiya, imya, otchestvo, nomer_telefona)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`, lastName, firstName, middleName, req.Phone).Scan(&personID)
	if err != nil {
		return fmt.Errorf("%s: insert personal data: %w", op, err)
	}

	// 2. Получение или вставка должности
	var positionID int
	err = tx.QueryRow(`SELECT id FROM dolzhnosti WHERE dolzhnost = $1`, req.Position).Scan(&positionID)
	if err == sql.ErrNoRows {
		err = tx.QueryRow(`INSERT INTO dolzhnosti (dolzhnost) VALUES ($1) RETURNING id`, req.Position).Scan(&positionID)
	}
	if err != nil {
		return fmt.Errorf("%s: get/insert position: %w", op, err)
	}

	// 3. Вставка в staff
	var staffID int
	err = tx.QueryRow(`
		INSERT INTO staff (id_pers_dannykh, id_dolzhnosti, nomer_klyucha)
		VALUES ($1, $2, $3)
		RETURNING id
	`, personID, positionID, req.KeyNumber).Scan(&staffID)
	if err != nil {
		return fmt.Errorf("%s: insert staff: %w", op, err)
	}

	// 4. Привязка зон
	// Сопоставление ключей из фронта с названиями зон в БД
	zoneMap := map[string]struct {
		Value bool
		Name  string
	}{
		"flightZone":   {req.Zone.FlightZone, "Зона A"},
		"clearZone":    {req.Zone.ClearZone, "Зона B"},
		"runaway":      {req.Zone.Runaway, "Зона C"},
		"baggageZone":  {req.Zone.BaggageZone, "Зона D"},
		"controlTower": {req.Zone.ControlTower, "Зона E"},
	}

	// Добавляем доступные зоны сотруднику
	for _, zone := range zoneMap {
		if !zone.Value {
			continue
		}

		var zonaID int
		err := tx.QueryRow(`SELECT id FROM zony WHERE name = $1`, zone.Name).Scan(&zonaID)
		if err != nil {
			fmt.Println("Ошибка поиска зоны:", zone.Name, err)
			continue
		}

		_, err = tx.Exec(`INSERT INTO staff_zones (staff_id, zona_id) VALUES ($1, $2)`, staffID, zonaID)
		if err != nil {
			fmt.Println("Ошибка вставки в staff_zones:", err)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("%s: commit failed: %w", op, err)
	}
	return nil
}

// обновляет информацию о сотруднике и его зоны доступа
func (s *Storage) UpdateStaff(staffID int, req CreateStaffRequest) error {
	const op = "storage.postgres.UpdateStaff"
	tx, err := s.db.Begin()
	if err != nil {
		return fmt.Errorf("%s: begin tx: %w", op, err)
	}
	defer tx.Rollback()

	// 1. Обновление personalnye_dannye
	parts := strings.Fields(req.Name)
	if len(parts) < 2 {
		return fmt.Errorf("%s: имя должно содержать хотя бы фамилию и имя", op)
	}
	lastName := parts[0]
	firstName := parts[1]
	var middleName sql.NullString
	if len(parts) > 2 {
		middleName = sql.NullString{String: parts[2], Valid: true}
	} else {
		middleName = sql.NullString{Valid: false}
	}

	_, err = tx.Exec(`
		UPDATE personalnye_dannye
		SET familiya = $1, imya = $2, otchestvo = $3, nomer_telefona = $4
		WHERE id = (SELECT id_pers_dannykh FROM staff WHERE id = $5)
	`, lastName, firstName, middleName, req.Phone, staffID)
	if err != nil {
		return fmt.Errorf("%s: update personal data: %w", op, err)
	}

	// 2. Получение/вставка и обновление должности
	var positionID int
	err = tx.QueryRow(`SELECT id FROM dolzhnosti WHERE dolzhnost = $1`, req.Position).Scan(&positionID)
	if err == sql.ErrNoRows {
		err = tx.QueryRow(`INSERT INTO dolzhnosti (dolzhnost) VALUES ($1) RETURNING id`, req.Position).Scan(&positionID)
	}
	if err != nil {
		return fmt.Errorf("%s: get/insert position: %w", op, err)
	}

	_, err = tx.Exec(`UPDATE staff SET id_dolzhnosti = $1, nomer_klyucha = $2 WHERE id = $3`, positionID, req.KeyNumber, staffID)
	if err != nil {
		return fmt.Errorf("%s: update staff: %w", op, err)
	}

	// 3. Очистка старых зон доступа
	_, err = tx.Exec(`DELETE FROM staff_zones WHERE staff_id = $1`, staffID)
	if err != nil {
		return fmt.Errorf("%s: clear zones: %w", op, err)
	}

	// 4. Добавление новых зон
	zones := map[string]bool{
		"Зона A": req.Zone.FlightZone,
		"Зона B": req.Zone.ClearZone,
		"Зона C": req.Zone.Runaway,
		"Зона D": req.Zone.BaggageZone,
		"Зона E": req.Zone.ControlTower,
	}
	for name, active := range zones {
		if !active {
			continue
		}
		var zoneID int
		err = tx.QueryRow(`SELECT id FROM zony WHERE name = $1`, name).Scan(&zoneID)
		if err != nil {
			return fmt.Errorf("%s: find zone %s: %w", op, name, err)
		}
		_, err = tx.Exec(`INSERT INTO staff_zones (staff_id, zona_id) VALUES ($1, $2)`, staffID, zoneID)
		if err != nil {
			return fmt.Errorf("%s: bind zone: %w", op, err)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("%s: commit: %w", op, err)
	}
	return nil
}
