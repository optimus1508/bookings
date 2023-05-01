package repository

import "github.com/optimus1508/bookings/internal/models"

type DatabaseRepo interface {
	AllUsers() bool

	InsertReservation(res models.Reservation) error
}
