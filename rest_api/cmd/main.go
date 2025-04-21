package main

import (
	"log"
	"os"

	"github.com/Dashiiidzzze/admin_panel_for_rfid_security/internal/repo"
)

func main() {
	log.SetOutput(os.Stdout) // Логи будут отправляться в Docker logs
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbIPport := os.Getenv("DB_IP_PORT")

	// Строка подключения к базе данных
	connString := "postgres://" + dbUser + ":" + dbPassword + "@" + dbIPport + "/" + dbName
	log.Printf("Строка подключения к базе данных: %s", connString)
	// Инициализация базы данных
	repo.InitDB(connString)
	defer repo.CloseDB() // Закрытие соединения при завершении программы
}
