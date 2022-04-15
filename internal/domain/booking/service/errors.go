package service

import "fmt"

var ErrBookingDateConflict = fmt.Errorf("SpaceX already has a launch on this date")
var ErrBookingDateTimetable = fmt.Errorf("wrong destination for this date")

func ErrBookingService(err error) error {
	if err == nil {
		return fmt.Errorf("booking service error")
	}

	return fmt.Errorf("booking service error: %w", err)
}
