package main

import (
	"fmt"
	"interview/models"
	"interview/service/booking"
	"interview/service/owner"
	"interview/service/turf"
	"interview/service/user"
	"interview/stores"
	"time"
)

func main() {
	storeObj := stores.NewStore()
	userObj := user.NewUserService(storeObj)
	ownerObj := owner.NewOwnerService(storeObj)
	turfObj := turf.NewTurfService(storeObj, userObj, ownerObj)
	bookingObj := booking.NewBookingService(storeObj, userObj, ownerObj, turfObj)

	//Add use
	user, err := userObj.RegisterUser(&models.User{
		Name:  "Sunny Kumar",
		Email: "sunnyk.nitjsr@gmail.com",
	})
	fmt.Println(user, err)

	owner, err := ownerObj.Registerowner(&models.Owner{
		User: models.User{
			Name:  "Rahul Kumar",
			Email: "example@gmail.com",
		},
	})
	fmt.Println(owner, err)

	//Add turf
	turf, err := turfObj.AddTurf(1, "Green Field")
	fmt.Println(turf, err)

	//Add slot
	startTime := time.Date(2024, time.July, int(time.Sunday), 7, 0, 0, 0, time.Local)
	endTime := time.Date(2024, time.July, int(time.Sunday), 8, 0, 0, 0, time.Local)
	slot, err := turfObj.AddSlot(&models.Slot{
		TurfId:    1,
		StartTime: &startTime,
		EndTime:   &endTime,
	})

	fmt.Println(slot, err)

	//Add booking

	booking, err := bookingObj.AddBooking(1, 1)
	fmt.Println(booking, err)

	//Cancel Booking
	err = bookingObj.CancelBooking(1)
	fmt.Println(err)
	bookings, err := bookingObj.GetAllBooking(1)
	fmt.Println(bookings, err)
}
