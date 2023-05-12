package repository

import (
	"time"

	"github.com/optimus1508/bookings/internal/models"
)

type DatabaseRepo interface {
	AllUsers() bool

	InsertReservation(res models.Reservation) (int, error)
	InsertRoomRestriction(r models.RoomRestriction) error
	SearchAvailabilityByDatesByRoomID(start, end time.Time, roomID int) (bool, error)
	SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error)
	GetRoomByID(id int) (models.Room, error)

	GetUserByID(id int) (models.User, error)
	UpdateUser(u models.User) error
	Authenticate(email, testPassword string) (int, string, error)

	AllReservations() ([]models.Reservation, error)
	AllNewReservations() ([]models.Reservation, error)
	AllRooms() ([]models.Room, error)
	GetReservationByID(id int) (models.Reservation, error)

	UpdateReservation(u models.Reservation) error
	DeleteReservation(id int) error
	UpdateProcessesForReservation(id, processed int) error
}
