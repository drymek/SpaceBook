package service

import "fmt"

func ErrInvalidBookingDate(err *error) error {
	if err == nil {
		return fmt.Errorf("invalid booking date")
	}

	return fmt.Errorf("invalid booking date: %w", *err)
}
