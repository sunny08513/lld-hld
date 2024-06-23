package turf

import (
	"errors"
	"interview/models"
	"interview/service/owner"
	"interview/service/user"
	"interview/stores"
)

type TurfService struct {
	UserService  *user.UserService
	OwnerService *owner.OwnerService
	Store        *stores.Strore
}

func NewTurfService(Store *stores.Strore, UserService *user.UserService, OwnerService *owner.OwnerService) *TurfService {
	return &TurfService{
		UserService:  UserService,
		OwnerService: OwnerService,
		Store:        Store,
	}
}

func (ts *TurfService) AddTurf(ownerId int, name string) (*models.Turf, error) {
	if name == "" {
		return nil, errors.New("name should not empty")
	}
	if ownerId <= 0 {
		return nil, errors.New("ownerId is ivalid")
	}

	_, err := ts.OwnerService.GetOwnerById(ownerId)
	if err != nil {
		return nil, errors.New("ownerId is ivalid")
	}

	turf := models.Turf{
		Id:      len(ts.Store.Turfs) + 1,
		Name:    name,
		OwnerId: ownerId,
	}
	ts.Store.Turfs = append(ts.Store.Turfs, turf)
	return &turf, nil
}

func (ts *TurfService) GetTurfById(id int) (*models.Turf, error) {
	if id <= 0 {
		return nil, errors.New("id not valid")
	}

	for _, turf := range ts.Store.Turfs {
		if turf.Id == id {
			return &turf, nil
		}
	}

	return nil, errors.New("turf not available")
}

func (ts *TurfService) AddSlot(slot *models.Slot) (*models.Slot, error) {
	_, err := ts.GetTurfById(slot.TurfId)
	if err != nil {
		return nil, err
	}

	if slot.StartTime == nil {
		return nil, errors.New("invalid start time")
	}

	if slot.EndTime == nil {
		return nil, errors.New("invalid end time")
	}

	slot.Id = len(ts.Store.Slots) + 1
	ts.Store.Slots = append(ts.Store.Slots, *slot)

	return slot, nil
}

func (u *TurfService) GetSlotById(id int) (*models.Slot, error) {
	for _, slot := range u.Store.Slots {
		if slot.Id == id {
			return &slot, nil
		}
	}

	return nil, errors.New("slot not found")
}
