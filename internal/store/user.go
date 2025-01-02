package store

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID              string  `json:"id"`                // UUID único
	Name            string  `json:"name"`              // Nome do usuário
	Email           string  `json:"email"`             // E-mail para login e comunicação
	Password        string  `json:"password"`          // Senha do usuário (será hasheada)
	ImagePath       *string `json:"image_path"`        // Foto de perfil (opcional)
	LastLoginAt     *string `json:"last_login_at"`     // Última data/hora de login
	IsVerified      bool    `json:"is_verified"`       // Usuário já confirmou o e-mail?
	AuthToken       *string `json:"auth_token"`        // Token de autenticação (para Magic Link)
	AuthTokenExpiry *string `json:"auth_token_expiry"` // Validade do token de autenticação
	CreatedAt       string  `json:"created_at"`
	UpdatedAt       string  `json:"updated_at"`
	DeletedAt       *string `json:"deleted_at"` // Soft delete
}

type UserStore struct {
	db *sql.DB
}

func (s *UserStore) Create(ctx context.Context, user *User) error {
	// Gera um novo UUID para o usuário
	user.ID = uuid.New().String()

	// Define valores padrão
	now := time.Now().UTC().Format(time.RFC3339)
	user.CreatedAt = now
	user.UpdatedAt = now
	user.IsVerified = false

	query := `
		INSERT INTO users (
			id, name, email, password, created_at, updated_at, is_verified
		) VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id`

	err := s.db.QueryRowContext(
		ctx,
		query,
		user.ID,
		user.Name,
		user.Email,
		user.Password, // Nota: A senha deve ser hasheada antes de chegar aqui
		user.CreatedAt,
		user.UpdatedAt,
		user.IsVerified,
	).Scan(&user.ID)

	if err != nil {
		return fmt.Errorf("erro ao criar usuário: %w", err)
	}

	return nil
}
