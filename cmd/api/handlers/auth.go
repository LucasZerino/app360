package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"app360/internal/auth"
	"app360/internal/errors"
	"app360/internal/store"

	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	store *store.Storage
	jwt   *auth.JWTManager
}

func NewAuthHandler(store *store.Storage, jwt *auth.JWTManager) *AuthHandler {
	return &AuthHandler{
		store: store,
		jwt:   jwt,
	}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var params store.LoginParams
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		respondWithError(w, errors.NewAPIError(http.StatusBadRequest, "Formato de dados inválido"))
		return
	}

	user, err := h.store.User.GetByEmail(r.Context(), params.Email)
	if err != nil {
		respondWithError(w, errors.NewAPIError(http.StatusUnauthorized, "Credenciais inválidas"))
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(params.Password)); err != nil {
		respondWithError(w, errors.NewAPIError(http.StatusUnauthorized, "Credenciais inválidas"))
		return
	}

	token, err := h.jwt.GenerateToken(user.ID, user.Email)
	if err != nil {
		respondWithError(w, errors.NewAPIError(http.StatusInternalServerError, "Erro ao gerar token"))
		return
	}

	// Configurar o cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "auth_token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteStrictMode,
		Expires:  time.Now().Add(24 * time.Hour),
	})

	respondWithJSON(w, http.StatusOK, map[string]interface{}{
		"message": "Login realizado com sucesso",
		"token":   token,
		"user": map[string]interface{}{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		},
	})
}

func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	// Remover o cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "auth_token",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		Expires:  time.Now().Add(-1 * time.Hour),
	})

	respondWithJSON(w, http.StatusOK, map[string]string{
		"message": "Logout realizado com sucesso",
	})
}
