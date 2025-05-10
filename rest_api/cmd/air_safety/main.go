package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Dashiiidzzze/admin_panel_for_rfid_security/internal/auth"
	"github.com/Dashiiidzzze/admin_panel_for_rfid_security/internal/config"
	authhandler "github.com/Dashiiidzzze/admin_panel_for_rfid_security/internal/http-server/handlers/auth"
	"github.com/Dashiiidzzze/admin_panel_for_rfid_security/internal/http-server/handlers/logsPasses"
	"github.com/Dashiiidzzze/admin_panel_for_rfid_security/internal/http-server/handlers/positions"
	"github.com/Dashiiidzzze/admin_panel_for_rfid_security/internal/http-server/handlers/staff"
	"github.com/Dashiiidzzze/admin_panel_for_rfid_security/internal/lib/logger/sl"
	"github.com/Dashiiidzzze/admin_panel_for_rfid_security/internal/storage/postgres"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

// чтобы не использовать текст в прямом виде
const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	// TODO: init config
	cfg := config.MustLoad()

	// TODO: init logger: slog
	log := setupLogger(cfg.Env)
	log.Info("starting air safety api", slog.String("env", cfg.Env)) // выведем окружение в котором находимся
	//log.Debug("Debug messages are enabled")

	// TODO: init storage: postgres
	storage, err := postgres.New(cfg.StoragePath)
	if err != nil {
		log.Error("failed to init storage", sl.Err(err))
		os.Exit(1)
	}

	// TODO: init router: chi (совместим с net/http), render
	router := chi.NewRouter()

	// Разрешаем CORS
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // разрешённый фронтенд
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300, // Cache preflight result for 5 min
	}))

	// Подключаем middleware
	// router.Use(middleware.RequestID) // добавляет к каждому запросу уникальный ID
	// router.Use(middleware.Recoverer) // Восстановление после panic

	// публичный маршрут
	router.Post("/login", authhandler.LoginHandler)

	// защищённые ручки
	router.Group(func(r chi.Router) {
		r.Use(auth.AuthMiddleware) // Мидлварь на все вложенные роуты

		// /items
		r.Route("/items", func(r chi.Router) {
			r.Get("/{id}", staff.ReadCard(log, storage))
			r.Get("/", staff.ReadAll(log, storage))
			r.Post("/", staff.Create(log, storage))
			r.Put("/{id}", staff.Update(log, storage))
			r.Delete("/{id}", staff.Delete(log, storage))
		})

		// Прочие защищённые ручки
		r.Get("/api/positions", positions.New(log, storage))
		r.Get("/logs", logsPasses.New(log, storage))
	})

	log.Info("starting server", slog.String("address", cfg.Address))

	// Канал для обработки системных сигналов
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Создаём HTTP-сервер с параметрами из конфигурации
	srv := &http.Server{
		Addr:         cfg.Address,                // Адрес, на котором запускается сервер
		Handler:      router,                     // Обработчик запросов — наш роутер
		ReadTimeout:  cfg.HTTPServer.Timeout,     // Тайм-аут чтения запроса
		WriteTimeout: cfg.HTTPServer.Timeout,     // Тайм-аут записи ответа
		IdleTimeout:  cfg.HTTPServer.IdleTimeout, // Тайм-аут ожидания новых соединений
	}

	// Запускаем сервер в отдельной горутине
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Error("server error", sl.Err(err))
		}
	}()

	log.Info("server started")

	<-done // Ждём сигнала завершения
	log.Info("stopping server")

	// Создаём контекст с тайм-аутом для graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Завершаем работу сервера корректно
	if err := srv.Shutdown(ctx); err != nil {
		log.Error("failed to stop server", sl.Err(err))
	}

	// Закрываем соединение с базой данных
	if err := storage.Close(); err != nil {
		log.Error("failed to close storage", sl.Err(err))
	}

	log.Info("server stopped")
}

// Настройка логгера в зависимости от окружения
func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal: // логгер для локальной разработки и дебага
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)

	case envDev: // логгер для dev окружения
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)

	case envProd: // логгер для продакшена
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
