package stores

import (
	"interview/models"
)

type Strore struct {
	Users    []models.User
	Owners   []models.Owner
	Turfs    []models.Turf
	Slots    []models.Slot
	Bookings []models.Booking
}

func NewStore() *Strore {
	return &Strore{
		Users:    []models.User{},
		Owners:   []models.Owner{},
		Turfs:    []models.Turf{},
		Slots:    []models.Slot{},
		Bookings: []models.Booking{},
	}
}
