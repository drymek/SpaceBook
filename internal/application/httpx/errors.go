package httpx

import (
	"fmt"
)

func ErrBadRequest(err *error) error {
	if *err == nil {
		return fmt.Errorf("bad request")
	}

	return fmt.Errorf("bad request: %w", *err)
}
