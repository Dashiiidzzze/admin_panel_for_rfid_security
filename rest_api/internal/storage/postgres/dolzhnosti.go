package postgres

import (
	"fmt"
)

// Вспомогательная структура
type Dolzhnost struct {
	ID        int    `json:"id"`
	Dolzhnost string `json:"dolzhnost"`
}

// Получает список должностей из БД
func (s *Storage) GetDolzhnosti() ([]Dolzhnost, error) {
	const op = "storage.postgres.GetDolzhnosti"

	const query = "SELECT id, dolzhnost FROM dolzhnosti"

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("%s: prepare statement: %w", op, err)
	}
	defer rows.Close()

	var result []Dolzhnost
	for rows.Next() {
		var d Dolzhnost
		if err := rows.Scan(&d.ID, &d.Dolzhnost); err != nil {
			return nil, fmt.Errorf("%s: execute statement: %w", op, err)
		}
		result = append(result, d)
	}

	return result, nil
}
