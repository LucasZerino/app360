package errors

func (e *APIError) Error() string {
	return e.Message
}
