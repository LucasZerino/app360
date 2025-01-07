package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"app360/internal/store"
	"app360/internal/testutil"
)

func TestCreateUser(t *testing.T) {
	tests := []struct {
		name           string
		input          store.CreateUserParams
		expectedStatus int
		expectedError  string
	}{
		{
			name: "sucesso_criar_usuario",
			input: store.CreateUserParams{
				Name:     "John Doe",
				Email:    "john@example.com",
				Password: "123456",
			},
			expectedStatus: http.StatusCreated,
		},
		{
			name: "erro_email_duplicado",
			input: store.CreateUserParams{
				Name:     "Jane Doe",
				Email:    "existing@example.com",
				Password: "123456",
			},
			expectedStatus: http.StatusConflict,
			expectedError:  "Email já cadastrado",
		},
		{
			name: "erro_senha_curta",
			input: store.CreateUserParams{
				Name:     "Jane Doe",
				Email:    "jane@example.com",
				Password: "123",
			},
			expectedStatus: http.StatusBadRequest,
			expectedError:  "senha deve ter no mínimo 6 caracteres",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			db := testutil.NewTestDB(t)
			defer db.Close()

			store := store.NewStorage(db)
			handler := NewUserHandler(store)

			// Criar request
			body, _ := json.Marshal(tt.input)
			req := httptest.NewRequest(http.MethodPost, "/api/users", bytes.NewBuffer(body))
			rec := httptest.NewRecorder()

			// Executar handler
			handler.CreateUser(rec, req)

			// Verificar resultado
			if rec.Code != tt.expectedStatus {
				t.Errorf("expected status %d; got %d", tt.expectedStatus, rec.Code)
			}

			if tt.expectedError != "" {
				var response map[string]string
				json.NewDecoder(rec.Body).Decode(&response)
				if response["message"] != tt.expectedError {
					t.Errorf("expected error message %q; got %q", tt.expectedError, response["message"])
				}
			}
		})
	}
}
