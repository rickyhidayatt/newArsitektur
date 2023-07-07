package liberror

import "errors"

// Errors
var (
	ErrWrongOTP        = errors.New("wrong one time password")
	ErrWrongCredential = errors.New("invalid user credentials")
	ErrUserExisted     = errors.New("user the same email or phone number already exists")
	ErrNotFound        = errors.New("not found")
	ErrNoInput         = errors.New("no input")
	ErrInvalidRequest  = errors.New("invalid request")
	ErrServerError     = errors.New("internal server error")
	ErrUnauthorized    = errors.New("unauthorized")
)

// session types
const (
	LoginSessionType = "login"
	OTPSessionType   = "otp"
)
