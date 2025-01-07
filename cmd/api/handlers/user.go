package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"

	apierrors "app360/internal/errors"
	"app360/internal/store"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Handlers struct {
	User *UserHandler
}

func NewHandlers(store *store.Storage) *Handlers {
	return &Handlers{
		User: NewUserHandler(store),
	}
}

type UserHandler struct {
	store *store.Storage
}

func NewUserHandler(store *store.Storage) *UserHandler {
	return &UserHandler{store: store}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var params store.CreateUserParams
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		respondWithError(w, apierrors.NewAPIError(http.StatusBadRequest, "Formato de dados inválido"))
		return
	}

	if err := validateCreateUserParams(params); err != nil {
		respondWithError(w, apierrors.NewAPIError(http.StatusBadRequest, err.Error()))
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Erro ao gerar hash da senha: %v", err)
		respondWithError(w, apierrors.NewAPIError(http.StatusInternalServerError, "Erro ao processar senha"))
		return
	}

	user := &store.User{
		ID:       uuid.New().String(),
		Name:     params.Name,
		Email:    params.Email,
		Password: string(hashedPassword),
	}

	if err := h.store.User.Create(r.Context(), user); err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			respondWithError(w, apierrors.NewAPIError(http.StatusConflict, "Email já cadastrado"))
			return
		}
		log.Printf("Erro ao criar usuário: %v", err)
		respondWithError(w, apierrors.NewAPIError(http.StatusInternalServerError, "Erro ao criar usuário"))
		return
	}

	respondWithJSON(w, http.StatusCreated, user)
}

func validateCreateUserParams(params store.CreateUserParams) error {
	if params.Name == "" {
		return errors.New("nome é obrigatório")
	}
	if params.Email == "" {
		return errors.New("email é obrigatório")
	}
	if params.Password == "" {
		return errors.New("senha é obrigatória")
	}
	if len(params.Password) < 6 {
		return errors.New("senha deve ter no mínimo 6 caracteres")
	}
	return nil
}
