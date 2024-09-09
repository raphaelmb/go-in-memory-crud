package types

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
	Biography string
}

func NewUser(id uuid.UUID, firstName, lastName, bio string) (User, error) {
	user := User{ID: id, FirstName: firstName, LastName: lastName, Biography: bio}
	if err := validateUser(user); err != nil {
		return User{}, err
	}
	return user, nil
}

func validateUser(user User) error {
	var (
		errs        []error
		minNameChar = 2
		maxNameChar = 20
		minBioChar  = 20
		maxBioChar  = 450
	)

	if err := validateName(user.FirstName, "first name", minNameChar, maxNameChar); err != nil {
		errs = append(errs, err)
	}
	if err := validateName(user.LastName, "last name", minNameChar, maxNameChar); err != nil {
		errs = append(errs, err)
	}
	if err := validateBio(user.Biography, minBioChar, maxBioChar); err != nil {
		errs = append(errs, err)
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	return nil

}

func validateName(name, fieldName string, minLen, maxLen int) error {
	if len(name) < minLen || len(name) > maxLen {
		return fmt.Errorf("%s length should be greater than %d and less than %d characters", fieldName, minLen, maxLen)
	}
	return nil
}

func validateBio(bio string, minLen, maxLen int) error {
	if len(bio) < minLen || len(bio) > maxLen {
		return fmt.Errorf("biography length should be greater than %d and less than %d characters", minLen, maxLen)
	}
	return nil
}
