package api

import (
	"encoding/json"
	"fmt"
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

		data, _ := json.Marshal(users)

		w.Write(data)
	}
}

func handleInsert(db db.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body utils.UserDTO
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			// TODO
			return
		}

		id, err := uuid.NewV7()
		if err != nil {
			// TODO
			return
		}

		user, err := types.NewUser(id, body.FirstName, body.LastName, body.Biography)
		if err != nil {
			// TODO
			fmt.Println(err)
			return
		}

		db.InsertUser(id, user)

		w.WriteHeader(http.StatusCreated)
	}

}

func handleFindByID(db db.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func handleDelete(db db.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func handleUpdate(db db.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
