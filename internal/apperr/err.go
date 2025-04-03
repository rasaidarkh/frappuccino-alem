package apperr

type statusError struct {
	error
	Message string
	Status  int
}

func (e statusError) Unwrap() error   { return e.error }
func (e statusError) HTTPStatus() int { return e.Status }

func WithHTTPStatus(err error, msg string, status int) error {
	return statusError{
		error:   err,
		Message: msg,
		Status:  status,
	}
}
