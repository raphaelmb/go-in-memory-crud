package database

import (
	"errors"
	"sync"

	"github.com/google/uuid"
	"github.com/raphaelmb/go-in-memory-crud/types"
)

var (
	ErrUserIDNotExists = errors.New("the user with the specified ID does not exist")
)

type Database struct {
	DB map[uuid.UUID]types.User
	mu *sync.Mutex
}

func NewDB() Database {
	return Database{
		DB: make(map[uuid.UUID]types.User),
		mu: &sync.Mutex{},
	}
}

func (d *Database) Insert(id uuid.UUID, user types.User) (types.User, error) {
	d.mu.Lock()
	defer d.mu.Unlock()

	d.DB[id] = user

	return d.DB[id], nil
}

func (d *Database) FindAll() ([]types.User, error) {
	users := []types.User{}
	for _, user := range d.DB {
		users = append(users, user)
	}

	return users, nil
}

func (d *Database) FindById(id uuid.UUID) (types.User, error) {
	user, ok := d.DB[id]
	if !ok {
		return types.User{}, ErrUserIDNotExists
	}

	return user, nil
}

func (d *Database) Delete(id uuid.UUID) error {
	if _, ok := d.DB[id]; !ok {
		return ErrUserIDNotExists
	}

	d.mu.Lock()
	defer d.mu.Unlock()

	delete(d.DB, id)

	return nil
}

func (d *Database) Update(id uuid.UUID, user types.User) (types.User, error) {
	if _, ok := d.DB[id]; !ok {
		return types.User{}, ErrUserIDNotExists
	}

	d.mu.Lock()
	defer d.mu.Unlock()

	d.DB[id] = user

	return d.DB[id], nil
}
