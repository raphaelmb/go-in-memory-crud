package utils

// TODO:
// send json

type UserDTO struct {
	ID        string `json:"id,omitempty"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Biography string `json:"biography"`
}

type Response struct {
	Data  any      `json:"data,omitempty"`
	Error []string `json:"error,omitempty"`
}
