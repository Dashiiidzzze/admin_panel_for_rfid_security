package auth

import (
	"net/http"
	"strings"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Получаем заголовок Authorization: Bearer <токен>
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		// Извлекаем сам токен
		token := strings.TrimPrefix(authHeader, "Bearer ")

		// Проверяем токен
		_, err := ValidateToken(token)
		if err != nil {
			http.Error(w, "недействительный или просроченный токен", http.StatusUnauthorized)
			return
		}

		// Продолжаем выполнение запроса
		next.ServeHTTP(w, r)
	})
}
