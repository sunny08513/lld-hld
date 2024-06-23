package user

import (
	"errors"
	"interview/models"
	"interview/stores"
)

type UserService struct {
	store *stores.Strore
}

func NewUserService(store *stores.Strore) *UserService {
	return &UserService{
		store: store,
	}
}

func (u *UserService) RegisterUser(user *models.User) (*models.User, error) {
	if user.Name == "" {
		return nil, errors.New("name should not empty")
	}
	if user.Email == "" {
		return nil, errors.New("password should not empty")
	}

	user.Id = len(u.store.Users) + 1
	u.store.Users = append(u.store.Users, *user)
	return user, nil
}

func (u *UserService) GetUser() ([]models.User, error) {
	return u.store.Users, nil
}

func (u *UserService) GetuserById(id int) (*models.User, error) {
	for _, user := range u.store.Users {
		if user.Id == id {
			return &user, nil
		}
	}

	return nil, errors.New("user not found")
}
