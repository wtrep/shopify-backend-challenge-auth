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

func (u *User) goodPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword(u.PwHash, []byte(password))
	if err != nil {
		return false
	}
	return true
}
