package liberror

type Error struct {
	Err        error
	StatusCode int
	Message    string
}

func New(err error, statusCode int, message string) *Error {
	e := &Error{
		Err:        err,
		StatusCode: statusCode,
		Message:    message,
	}

	return e
}

func (e *Error) Error() string {
	return e.Err.Error()
}
