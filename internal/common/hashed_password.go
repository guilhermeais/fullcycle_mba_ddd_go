package common

import (
	"golang.org/x/crypto/bcrypt"
)

type HashedPassword string

func CreateHashedPassword(plain string) (HashedPassword, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return HashedPassword(hashed), nil
}

func (hashed HashedPassword) Compare(plain string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))

	return err == nil
}
