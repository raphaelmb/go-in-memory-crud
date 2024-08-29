package db

import (
	"errors"
	"sync"

	"github.com/google/uuid"
	"github.com/raphaelmb/go-in-memory-crud/types"
)

type Database struct {
	Db map[uuid.UUID]types.User
	mu *sync.Mutex
}

func NewDB() Database {
	return Database{
		Db: make(map[uuid.UUID]types.User),
		mu: &sync.Mutex{},
	}
}

func (d *Database) InsertUser(id uuid.UUID, user types.User) {
	d.mu.Lock()
	defer d.mu.Unlock()

	d.Db[id] = user
}

func (d *Database) FindAllUsers() []types.User {
	users := []types.User{}
	for _, user := range d.Db {
		users = append(users, user)
	}
	return users
}

func (d *Database) FindUserByID(id uuid.UUID) (types.User, error) {
	user, ok := d.Db[id]
	if !ok {
		return types.User{}, errors.New("the user with the specified ID does not exist")
	}

	return user, nil
}

func (d *Database) DeleteUser(id uuid.UUID) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	if _, ok := d.Db[id]; !ok {
		return errors.New("the user with the specified ID does not exist")
	}

	delete(d.Db, id)

	return nil
}

func (d *Database) UpdateUser(id uuid.UUID, user types.User) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	if _, ok := d.Db[id]; !ok {
		return errors.New("user with given id not found")
	}

	d.Db[id] = user

	return nil
}
