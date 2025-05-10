package authhandler

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/Dashiiidzzze/admin_panel_for_rfid_security/internal/auth"
)

type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

// запрос аутентификации
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	login := os.Getenv("AIR_SAFETY_LOGIN")
	password := os.Getenv("AIR_SAFETY_PASS")
	if req.Login != login || req.Password != password {
		http.Error(w, "Неверный логин или пароль", http.StatusUnauthorized)
		return
	}

	token, err := auth.GenerateToken(req.Login)
	if err != nil {
		http.Error(w, "не удалось сгенерировать токен", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(LoginResponse{Token: token})
}
