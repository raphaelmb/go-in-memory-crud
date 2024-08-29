package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
	"github.com/raphaelmb/go-in-memory-crud/internal/db"
	"github.com/raphaelmb/go-in-memory-crud/types"
	"github.com/raphaelmb/go-in-memory-crud/utils"
)

func NewHandler(db db.Database) http.Handler {
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

func handleFindAll(db db.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users := db.FindAllUsers()

		usersDTO := toUserOutputDTOList(users)

		data, err := json.Marshal(usersDTO)
		if err != nil {
			sendJSON(w, Response{Error: "something went wrong"}, http.StatusInternalServerError)
			return
		}

		sendJSON(w, Response{Data: data}, http.StatusOK)
	}
}

func handleInsert(db db.Database) http.HandlerFunc {
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
			sendJSON(w, Response{Error: utils.CleanErrors(err.Error())}, http.StatusBadRequest)
			return
		}

		db.InsertUser(id, user)

		data := toUserOutputDTO(user)

		sendJSON(w, Response{Data: data}, http.StatusCreated)
	}

}

func handleFindByID(db db.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := checkUUID(r, "id")
		if err != nil {
			sendJSON(w, Response{Error: err.Error()}, http.StatusBadRequest)
			return
		}

		user, err := db.FindUserByID(id)
		if err != nil {
			sendJSON(w, Response{Error: err.Error()}, http.StatusNotFound)
			return
		}

		data := toUserOutputDTO(user)

		sendJSON(w, Response{Data: data}, http.StatusOK)
	}
}

func handleDelete(db db.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := checkUUID(r, "id")
		if err != nil {
			sendJSON(w, Response{Error: err.Error()}, http.StatusBadRequest)
			return
		}

		err = db.DeleteUser(id)
		if err != nil {
			sendJSON(w, Response{Error: err.Error()}, http.StatusNotFound)
			return
		}
	}
}

func handleUpdate(db db.Database) http.HandlerFunc {
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
			sendJSON(w, Response{Error: utils.CleanErrors(err.Error())}, http.StatusBadRequest)
			return
		}

		err = db.UpdateUser(id, user)
		if err != nil {
			sendJSON(w, Response{Error: err.Error()}, http.StatusNotFound)
			return
		}

		data := toUserOutputDTO(user)

		sendJSON(w, Response{Data: data}, http.StatusOK)
	}
}
