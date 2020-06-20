package errs

import (
	goerrors "errors"
	"fmt"
)

var (
	ErrInvalidContext    = goerrors.New("invalid context")
	ErrWellsFarGoContext = goerrors.New("wells-far-go context error")
)

type AppErr struct {
	Err     error
	Message string
	Code    int
}

func NewAppErr(err error, message string, code int) *AppErr {
	return &AppErr{
		Err:     err,
		Message: message,
		Code:    code,
	}
}

func (e *AppErr) Unwrap() error {
	return e.Err
}

func (e *AppErr) Error() string {
	return fmt.Sprintf("code: %d | err: %v | message: %s", e.Code, e.Err, e.Message)
}
