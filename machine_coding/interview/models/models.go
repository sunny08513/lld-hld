package models

import "time"

type User struct {
	Id    int
	Name  string
	Email string
}

type Owner struct {
	User
}

type Turf struct {
	Id      int
	OwnerId int
	Name    string
	Slots   []Slot
}

type Slot struct {
	Id        int
	TurfId    int
	StartTime *time.Time
	EndTime   *time.Time
	IsBooked  bool
	UserId    int
}

type Booking struct {
	Id       int
	UserId   int
	SlotId   int
	IsBooked bool
	Status   string // BOOKED, CANCELLED
}
