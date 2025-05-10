package postgres

import (
	"fmt"
	"time"
)

// Структура для логов проходов
type LogsPases struct {
	Date time.Time `json:"time"`
	Zone string    `json:"zone"`
	FIO  string    `json:"name"`
}

// Получает список логов проходов
func (s *Storage) GetLogsPases() ([]LogsPases, error) {
	const op = "storage.postgres.GetLogsPases"

	const query = `
		SELECT 
			lp.vremya,
			z.name,
			pd.familiya || ' ' || pd.imya || ' ' || COALESCE(pd.otchestvo, '')
		FROM logs_passes lp
		JOIN zony z ON lp.zone_id = z.id
		JOIN staff s ON lp.staff_id = s.id
		JOIN personalnye_dannye pd ON s.id_pers_dannykh = pd.id
		ORDER BY lp.vremya DESC
	`

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("%s: query error: %w", op, err)
	}
	defer rows.Close()

	var result []LogsPases
	for rows.Next() {
		var lp LogsPases
		if err := rows.Scan(&lp.Date, &lp.Zone, &lp.FIO); err != nil {
			return nil, fmt.Errorf("%s: scan error: %w", op, err)
		}
		result = append(result, lp)
	}

	return result, nil
}
