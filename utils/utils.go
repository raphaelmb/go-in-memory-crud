package utils

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/raphaelmb/go-in-memory-crud/types"
)

// TODO:
// send json

type UserInputDTO struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Biography string `json:"biography"`
}

type UserOutputDTO struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Biography string `json:"biography"`
}

func ToUserOutputDTOList(users []types.User) []UserOutputDTO {
	result := []UserOutputDTO{}
	for _, user := range users {
		u := UserOutputDTO{ID: user.ID.String(), FirstName: user.FirstName, LastName: user.LastName, Biography: user.Biography}
		result = append(result, u)
	}
	return result
}

type Response struct {
	Data  any      `json:"data,omitempty"`
	Error []string `json:"error,omitempty"`
}

func CheckUUID(r *http.Request, id string) (uuid.UUID, error) {
	idStr := chi.URLParam(r, id)
	uid, err := uuid.Parse(idStr)
	if err != nil {
		return uuid.Nil, err
	}
	return uid, nil
}
