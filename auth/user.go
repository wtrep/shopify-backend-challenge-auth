package auth

import (
	"github.com/wtrep/shopify-backend-challenge-auth/common"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string
	PwHash   []byte
}

const (
	defaultCost       = 10
	minPasswordLength = 8
	maxPasswordLength = 32
	minUserLength     = 4
	maxUserLength     = 24
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
func (u *User) isGoodPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword(u.PwHash, []byte(password))
	if err != nil {
		return false
	}
	return true
}

// Check if the length of the username and password is correct
func verifyCredentials(username string, password string) *common.ErrorResponseError {
	if len(username) < minUserLength {
		return &common.UserTooShortError
	}
	if len(username) > maxUserLength {
		return &common.UserTooLongError
	}
	if len(password) < minPasswordLength {
		return &common.PasswordTooShortError
	}
	if len(password) > maxPasswordLength {
		return &common.PasswordTooLongError
	}
	return nil
}
