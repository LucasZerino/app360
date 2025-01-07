package errors

import "errors"

var (
	ErrDuplicateEmail = errors.New("email já cadastrado")
	ErrInvalidInput   = errors.New("dados de entrada inválidos")
	ErrInternalServer = errors.New("erro interno do servidor")
)

type APIError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func NewAPIError(status int, message string) *APIError {
	return &APIError{
		Status:  status,
		Message: message,
	}
}
