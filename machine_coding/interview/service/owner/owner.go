package owner

import (
	"errors"
	"interview/models"
	"interview/stores"
)

type OwnerService struct {
	store *stores.Strore
}

func NewOwnerService(store *stores.Strore) *OwnerService {
	return &OwnerService{
		store: store,
	}
}

func (u *OwnerService) Registerowner(owner *models.Owner) (*models.Owner, error) {
	if owner.Name == "" {
		return nil, errors.New("name should not empty")
	}
	if owner.Email == "" {
		return nil, errors.New("password should not empty")
	}

	owner.Id = len(u.store.Owners) + 1
	u.store.Owners = append(u.store.Owners, *owner)
	return owner, nil
}

func (u *OwnerService) Getowner() ([]models.Owner, error) {
	return u.store.Owners, nil
}

func (u *OwnerService) GetOwnerById(id int) (*models.Owner, error) {
	for _, owner := range u.store.Owners {
		if owner.Id == id {
			return &owner, nil
		}
	}

	return nil, errors.New("owner not found")
}
