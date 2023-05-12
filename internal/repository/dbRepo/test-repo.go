package dbrepo

import (
	"errors"
	"time"

	"github.com/optimus1508/bookings/internal/models"
)

func (m *testDBRepo) AllUsers() bool {
	return true
}

// insertreservation inserts reservation into the database
func (m *testDBRepo) InsertReservation(res models.Reservation) (int, error) {
	return 1, nil
}

// InsertRoomRestriction inserts a room restriction into the database
func (m *testDBRepo) InsertRoomRestriction(r models.RoomRestriction) error {
	return nil
}

// SearchAvailabilityByDatesByRoomID returns true if there is availability for roomID and false when there isn't
func (m *testDBRepo) SearchAvailabilityByDatesByRoomID(start, end time.Time, roomID int) (bool, error) {
	return false, nil
}

// SearchAvailabilityForAllRooms returns a slice of available rooms if any, for given date range
func (m *testDBRepo) SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error) {
	var rooms []models.Room
	return rooms, nil
}

// GetRoomByID gets a room by its id
func (m *testDBRepo) GetRoomByID(id int) (models.Room, error) {
	var room models.Room
	if id > 2 {
		return room, errors.New("some error")
	}

	return room, nil
}

func (m *testDBRepo) GetUserByID(id int) (models.User, error) {
	var u models.User

	return u, nil
}

func (m *testDBRepo) UpdateUser(u models.User) error {
	return nil
}

func (m *testDBRepo) Authenticate(email, testPassword string) (int, string, error) {
	return 1, "", nil
}

func (m *testDBRepo) AllReservations() ([]models.Reservation, error) {
	var reservations []models.Reservation

	return reservations, nil
}

func (m *testDBRepo) AllNewReservations() ([]models.Reservation, error) {
	var reservations []models.Reservation

	return reservations, nil
}

func (m *testDBRepo) GetReservationByID(id int) (models.Reservation, error) {
	var res models.Reservation

	return res, nil
}

func (m *testDBRepo) UpdateReservation(u models.Reservation) error {
	return nil
}

// Deletereservation Deletes one reservation by Id
func (m *testDBRepo) DeleteReservation(id int) error {
	return nil
}

func (m *testDBRepo) UpdateProcessesForReservation(id, processed int) error {
	return nil
}
