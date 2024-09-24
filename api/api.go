package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
	"github.com/raphaelmb/go-in-memory-crud/internal/database"
	"github.com/raphaelmb/go-in-memory-crud/types"
	"github.com/raphaelmb/go-in-memory-crud/utils"
)

func NewHandler(db database.Database) http.Handler {
	r := chi.NewMux()

	r.Use(middleware.Recoverer, middleware.Logger, middleware.RequestID)

	r.Route("/api/users", func(r chi.Router) {
		r.Post("/", handleInsert(db))
		r.Get("/", handleFindAll(db))
		r.Get("/{id}", handleFindByID(db))
		r.Delete("/{id}", handleDelete(db))
		r.Put("/{id}", handleUpdate(db))
	})

	return r
}

func handleInsert(db database.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body UserInputDTO
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			sendJSON(w, Response{Error: "unable to process body"}, http.StatusUnprocessableEntity)
			return
		}

		err := validateBody(body)
		if err != nil {
			sendJSON(w, Response{Error: err.Error()}, http.StatusBadRequest)
			return
		}

		id, err := uuid.NewV7()
		if err != nil {
			sendJSON(w, Response{Error: "something went wrong"}, http.StatusInternalServerError)
			return
		}

		user, err := types.NewUser(id, body.FirstName, body.LastName, body.Biography)
		if err != nil {
			sendJSON(w, Response{Error: utils.FormatErrors(err.Error())}, http.StatusBadRequest)
			return
		}

		user, err = db.Insert(id, user)
		if err != nil {
			sendJSON(w, Response{Error: "There was an error while saving the user to the database"}, http.StatusInternalServerError)
			return
		}

		data := toUserOutputDTO(user)

		sendJSON(w, Response{Data: data}, http.StatusCreated)
	}

}

func handleFindAll(db database.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := db.FindAll()
		if err != nil {
			sendJSON(w, Response{Error: "The users information could not be retrieved"}, http.StatusInternalServerError)
			return
		}

		usersDTO := toUserOutputDTOList(users)

		sendJSON(w, Response{Data: usersDTO}, http.StatusOK)
	}
}

func handleFindByID(db database.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := checkUUID(r, "id")
		if err != nil {
			sendJSON(w, Response{Error: err.Error()}, http.StatusBadRequest)
			return
		}

		user, err := db.FindById(id)
		if err != nil {
			if errors.Is(err, database.ErrUserIDNotExists) {
				sendJSON(w, Response{Error: err.Error()}, http.StatusNotFound)
				return
			}
			sendJSON(w, Response{Error: "The user information could not be retrieved"}, http.StatusInternalServerError)
			return
		}

		data := toUserOutputDTO(user)

		sendJSON(w, Response{Data: data}, http.StatusOK)
	}
}

func handleDelete(db database.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := checkUUID(r, "id")
		if err != nil {
			sendJSON(w, Response{Error: err.Error()}, http.StatusBadRequest)
			return
		}

		err = db.Delete(id)
		if err != nil {
			if errors.Is(err, database.ErrUserIDNotExists) {
				sendJSON(w, Response{Error: err.Error()}, http.StatusNotFound)
				return
			}
			sendJSON(w, Response{Error: "The user could not be removed"}, http.StatusInternalServerError)
			return
		}

		sendJSON(w, Response{}, http.StatusNoContent)
	}
}

func handleUpdate(db database.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := checkUUID(r, "id")
		if err != nil {
			sendJSON(w, Response{Error: err.Error()}, http.StatusBadRequest)
			return
		}

		var body UserInputDTO
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			sendJSON(w, Response{Error: "unable to process body"}, http.StatusUnprocessableEntity)
			return
		}

		err = validateBody(body)
		if err != nil {
			sendJSON(w, Response{Error: err.Error()}, http.StatusBadRequest)
			return
		}

		user, err := types.NewUser(id, body.FirstName, body.LastName, body.Biography)
		if err != nil {
			sendJSON(w, Response{Error: utils.FormatErrors(err.Error())}, http.StatusBadRequest)
			return
		}

		user, err = db.Update(id, user)
		if err != nil {
			if errors.Is(err, database.ErrUserIDNotExists) {
				sendJSON(w, Response{Error: err.Error()}, http.StatusNotFound)
				return
			}
			sendJSON(w, Response{Error: "The user information could not be modified"}, http.StatusInternalServerError)
			return
		}

		data := toUserOutputDTO(user)

		sendJSON(w, Response{Data: data}, http.StatusOK)
	}
}
