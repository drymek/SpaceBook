package httpx

import (
	"net/http"
)

type HttpErrorWrapper struct {
	Message string `json:"message"`
	Code    int    `json:"-"`
	Err     error  `json:"-"`
}

func (err HttpErrorWrapper) Error() string {
	if err.Err != nil {
		return err.Err.Error()
	}
	return err.Message
}

func (err HttpErrorWrapper) Unwrap() error {
	return err.Err
}

func (err HttpErrorWrapper) StatusCode() int {
	return err.Code
}

func NewErrorWrapper(code int, err error, message string) error {
	return HttpErrorWrapper{
		Message: message,
		Code:    code,
		Err:     err,
	}
}

func NewBadRequest(err error) error {
	return NewErrorWrapper(http.StatusBadRequest, err, "Bad Request")
}

func NewInternalServerError(err error) error {
	return NewErrorWrapper(http.StatusInternalServerError, err, "Internal Server Error")
}
