package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/raphaelmb/go-in-memory-crud/types"
)

type Response struct {
	Data  any    `json:"data,omitempty"`
	Error string `json:"message,omitempty"`
}

func sendJSON(w http.ResponseWriter, resp Response, status int) {
	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(resp)
	if err != nil {
		slog.Error("failed to marshal json data", "error", err)
		sendJSON(w, Response{Error: "something went wrong"}, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)
	if _, err := w.Write(data); err != nil {
		slog.Error("failed to write response to client", "error", err)
		return
	}
}

type UserInputDTO struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Biography string `json:"biography"`
}

type UserOutputDTO struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Biography string `json:"biography"`
}

func toUserOutputDTO(user types.User) UserOutputDTO {
	return UserOutputDTO{
		ID:        user.ID.String(),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Biography: user.Biography,
	}
}

func toUserOutputDTOList(users []types.User) []UserOutputDTO {
	result := []UserOutputDTO{}
	for _, user := range users {
		u := UserOutputDTO{ID: user.ID.String(), FirstName: user.FirstName, LastName: user.LastName, Biography: user.Biography}
		result = append(result, u)
	}
	return result
}

func checkUUID(r *http.Request, id string) (uuid.UUID, error) {
	idStr := chi.URLParam(r, id)
	uid, err := uuid.Parse(idStr)
	if err != nil {
		return uuid.Nil, fmt.Errorf("invalid UUID for parameter '%s': %s", id, idStr)
	}
	return uid, nil
}

func validateBody(user UserInputDTO) error {
	if user.FirstName == "" || user.LastName == "" || user.Biography == "" {
		return errors.New("please provide FirstName LastName and bio for the user")
	}
	return nil
}
