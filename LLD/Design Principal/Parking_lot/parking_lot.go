package parkinglot

import (
	"fmt"
	"time"

	"container/list"
)

/*
Requirement
- Parking should have multiple floors
- Parking a vehicle and assign spots to vehicle
- Unparking a vehicle and free slot
- Check the sport for vehicle
- Handling diffrent type of vehicle
*/

/*Use case
- Parking Vehicle ans assigning spot to the vehicle
- Unparking vehicle and calculate fee
- Handling different type of vehicle(COMPACT,BIG, TWO_WHEELER, HANDICAPPED)
- Spot Avability to check
*/

/*Class
ParkingLot
ParkingFloor
ParkingSlot
Vehicle
Ticket
*/

type VehicleType int

const (
	COMPACT VehicleType = iota
	BIG
	TWO_WHEELER
	HANDICAPPED
)

type Vehicle struct {
	RegistrationNo string
	VehicleType    VehicleType
}

// constructor to create new vehicle
func NewVehicle(RegistrationNo string, VehicleType VehicleType) *Vehicle {
	return &Vehicle{
		RegistrationNo: RegistrationNo,
		VehicleType:    VehicleType,
	}
}

type ParkingSpot struct {
	id         int
	VType      VehicleType
	IsOccupied bool
}

func (p *ParkingSpot) OccupySpot() {
	p.IsOccupied = true
}

func (p *ParkingSpot) FreeSpot() {
	p.IsOccupied = false
}

type ParkingFloor struct {
	id           int
	ParkingFloor []ParkingFloor
}

type Ticket struct {
	id         int
	VehicleRNo string
	SlotId     int
	FloorId    int
	EntryTime  time.Time
	ExitTime   time.Time
	Fee        float64
}

func (t *Ticket) IssuedTicket(VehicleRNo string, SlotId, FloorId int, EntryTime time.Time) *Ticket {
	return &Ticket{
		VehicleRNo: VehicleRNo,
		SlotId:     SlotId,
		FloorId:    FloorId,
		EntryTime:  EntryTime,
	}
}

func (t *Ticket) Markpaid(id int, ExitTime time.Time, fee float64) *Ticket {
	//Update the ticket of particular id
	return &Ticket{
		ExitTime: time.Now(),
		Fee:      fee,
	}
}

type parkingLot struct {
	Id            int
	ParkingFloors []ParkingFloor
	l             *list.List
}

func (p *parkingLot) GetAvailableSpot(vType VehicleType) *ParkingSpot {
	//Search algo and get parking spot for given particular vehicle type
	return nil
}

func (p *parkingLot) BookTicket(v Vehicle) *Ticket {
	//Search algo and get parking spot for given particular vehicle type

	slot := p.GetAvailableSpot(v.VehicleType)
	fmt.Println(slot)
	//get slot and then book ticket
	return nil
}

func (p *parkingLot) ProcessPayment(t *Ticket, fee float64) {
	t.Markpaid(1, time.Now(), 50.5)
	//process payment
}
