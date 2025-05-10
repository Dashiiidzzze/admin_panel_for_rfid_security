package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq" // драйвер для работы с PostgreSQL
)

type Storage struct {
	db *sql.DB
}

func New(schemaPath string) (*Storage, error) {
	const op = "storage.postgres.New" // название функции для вывода ошибки можно использовать op вместо operation

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Формируем строку подключения
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)
	fmt.Println(dsn)
	// Открываем подключение к базе данных
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("%s: sql.Open: %w", op, err)
	}

	for i := 0; i < 10; i++ {
		err = db.Ping()
		if err == nil {
			break
		}
		log.Printf("waiting for database... (%d/10)", i+1)
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		return nil, fmt.Errorf("%s: db.Ping: %w", op, err)
	}

	// Читаем SQL-файл со схемой
	schemaBytes, err := os.ReadFile(schemaPath)
	if err != nil {
		return nil, fmt.Errorf("%s: read schema file: %w", op, err)
	}

	// Выполняем SQL-скрипт
	if _, err := db.Exec(string(schemaBytes)); err != nil {
		return nil, fmt.Errorf("%s: executing schema: %w", op, err)
	}

	// Возвращаем инициализированный Storage
	return &Storage{
		db: db,
	}, nil
}

// закрывает подключение к базе данных
func (s *Storage) Close() error {
	if s.db != nil {
		return s.db.Close() // закрываем соединение, если оно открыто
	}
	return nil // если соединение не было установлено, возвращаем nil
}
