package repository

import "github.com/rocky_bgta/booking-with-go/internal/models"

type DatabaseRepo interface {
	AllUser() bool

	InsertReservation(res models.Reservation) (int, error)
	InsertRoomRestriction(r models.RoomRestriction) error
}
