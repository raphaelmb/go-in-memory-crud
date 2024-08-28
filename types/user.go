package types

import (
	"errors"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
	Biography string
}

func NewUser(id uuid.UUID, firstName, lastName, bio string) (User, error) {
	var errs []error
	if len(firstName) <= 2 || len(firstName) >= 20 {
		errs = append(errs, errors.New("first name length should be greater than 2 and less than 20 characters"))
	}
	if len(lastName) <= 2 || len(lastName) >= 20 {
		errs = append(errs, errors.New("last name length should be greater than 2 and less than 20 characters"))
	}
	if len(bio) <= 20 || len(bio) >= 450 {
		errs = append(errs, errors.New("biography length should be greater than 20 and less than 450 characters"))
	}

	if errs != nil {
		return User{}, errors.Join(errs...)
	}

	return User{ID: id, FirstName: firstName, LastName: lastName, Biography: bio}, nil
}
