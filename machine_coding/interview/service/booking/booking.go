package booking

import (
	"errors"
	"interview/models"
	"interview/service/owner"
	"interview/service/turf"
	"interview/service/user"
	"interview/stores"
)

type BookingService struct {
	UserService  *user.UserService
	OwnerService *owner.OwnerService
	Store        *stores.Strore
	TurfService  *turf.TurfService
}

func NewBookingService(Store *stores.Strore, UserService *user.UserService, OwnerService *owner.OwnerService, TurfService *turf.TurfService) *BookingService {
	return &BookingService{
		UserService:  UserService,
		OwnerService: OwnerService,
		Store:        Store,
		TurfService:  TurfService,
	}
}

func (bs *BookingService) AddBooking(userId, slotId int) (*models.Booking, error) {
	if userId <= 0 {
		return nil, errors.New("invalid user")
	}

	if slotId <= 0 {
		return nil, errors.New("invalid slot")
	}

	_, err := bs.TurfService.GetSlotById(slotId)
	if err != nil {
		return nil, err
	}

	booking := models.Booking{
		Id:       len(bs.Store.Bookings) + 1,
		UserId:   userId,
		SlotId:   slotId,
		IsBooked: true,
		Status:   "BOOKED",
	}

	bs.Store.Bookings = append(bs.Store.Bookings, booking)

	return &booking, nil
}

func (bs *BookingService) GetAllBooking(id int) ([]models.Booking, error) {
	return bs.Store.Bookings, nil
}

func (bs *BookingService) CancelBooking(bookingId int) error {
	for i, booking := range bs.Store.Bookings {
		if booking.Id == bookingId {
			bs.Store.Bookings[i].Status = "CANCELLED"
			for j, slot := range bs.Store.Slots {
				if slot.Id == booking.SlotId {
					bs.Store.Slots[j].IsBooked = false
					bs.Store.Slots[j].UserId = 0
					break
				}
			}
			return nil
		}
	}
	return errors.New("booking not found")
}
