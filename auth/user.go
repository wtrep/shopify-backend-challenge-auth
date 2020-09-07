package auth

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string
	PwHash   []byte
}

const (
	defaultCost = 10
)

// Create a new User object by hashing the provided password
func NewUser(username string, password string) (*User, error) {
	pwHash, err := bcrypt.GenerateFromPassword([]byte(password), defaultCost)
	if err != nil {
		return nil, err
	}
	return &User{
		Username: username,
		PwHash:   pwHash,
	}, nil
}

// Check that the password match the stored hash
func (u *User) goodPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword(u.PwHash, []byte(password))
	if err != nil {
		return false
	}
	return true
}
