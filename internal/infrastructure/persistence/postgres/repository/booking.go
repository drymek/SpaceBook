package repository

import (
	"database/sql"
	"errors"
	"time"

	"dryka.pl/SpaceBook/internal/domain/booking/model"
	"dryka.pl/SpaceBook/internal/domain/booking/repository"
)

var ErrPersistencePrepareError = errors.New("persistence: prepare error")
var ErrPersistenceCannotAddBooking = errors.New("persistence: cannot add booking")
var ErrPersistence = errors.New("persistence: error")

type bookingRepository struct {
	db *sql.DB
}

func (b bookingRepository) Find(ID string) (model.Booking, error) {
	stmt, err := b.db.Prepare(
		`SELECT custom_id, firstname, lastname, gender, birthday, launchpad_id, destination_id, launch_date FROM bookings WHERE custom_id = $1`,
	)
	if err != nil {
		return model.Booking{}, ErrPersistencePrepareError
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			panic(err)
		}
	}(stmt)
	booking := &model.Booking{}
	row := stmt.QueryRow(ID)
	var birthday time.Time
	var launchDate time.Time
	if err := row.Scan(&booking.ID, &booking.Firstname, &booking.Lastname, &booking.Gender, &birthday, &booking.LaunchpadID, &booking.DestinationID, &launchDate); err != nil {
		if err == sql.ErrNoRows {
			return model.Booking{}, repository.ErrBookingNotFound
		}
		return model.Booking{}, ErrPersistence
	}
	booking.Birthday = model.DayDate{Time: birthday}
	booking.LaunchDate = model.DayDate{Time: launchDate}

	return *booking, nil
}

func (b bookingRepository) Delete(ID string) error {
	stmt, err := b.db.Prepare(
		`DELETE FROM bookings WHERE custom_id = $1`,
	)
	if err != nil {
		return ErrPersistencePrepareError
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			panic(err)
		}
	}(stmt)

	_, err = stmt.Exec(ID)
	if err != nil {
		return ErrPersistence
	}

	return nil
}

func (b bookingRepository) Create(booking *model.Booking) error {
	stmt, err := b.db.Prepare(
		`INSERT INTO bookings
   		(custom_id, firstname, lastname, gender, birthday, launchpad_id, destination_id, launch_date)
   		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
	)
	if err != nil {
		return ErrPersistencePrepareError
	}

	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			panic(err)
		}
	}(stmt)

	_, err = stmt.Exec(
		booking.ID,
		booking.Firstname,
		booking.Lastname,
		booking.Gender,
		booking.Birthday.Time,
		booking.LaunchpadID,
		booking.DestinationID,
		booking.LaunchDate.Time,
	)

	if err != nil {
		return ErrPersistenceCannotAddBooking
	}

	return nil
}

func (b bookingRepository) List() ([]model.Booking, error) {
	//TODO implement me
	panic("implement me")
}

func NewBookingRepository(db *sql.DB) repository.BookingRepository {
	return &bookingRepository{
		db: db,
	}
}
