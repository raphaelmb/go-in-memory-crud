package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/raphaelmb/go-in-memory-crud/api"
	"github.com/raphaelmb/go-in-memory-crud/internal/database"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	db := database.NewDB()
	handler := api.NewHandler(db)

	s := http.Server{
		Addr:         ":8080",
		Handler:      handler,
		IdleTimeout:  10 * time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("Server running on port 8080...")
	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
