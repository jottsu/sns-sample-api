package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserSecret struct {
	UserID         uint64
	HashedPassword string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func CreateUserSecret(userID uint64, password string) (*UserSecret, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	userSecret := &UserSecret{
		UserID:         userID,
		HashedPassword: string(hash),
	}
	return userSecret, nil
}

func (userSecret *UserSecret) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(userSecret.HashedPassword), []byte(password))
}
