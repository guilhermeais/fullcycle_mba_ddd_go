package common

import (
	"fmt"
	"regexp"
)

type Email string

const validateEmailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

func CreateEmail(address string) (Email, error) {
	re := regexp.MustCompile(validateEmailRegex)
	if !re.MatchString(address) {
		return "", MakeErrEmailValidation(address)
	}

	return Email(address), nil
}

func MakeErrEmailValidation(invalidAddress string) error {
	return fmt.Errorf("%w: o email %q não é valido", ErrValidation, invalidAddress)
}

func (actual Email) IsEqual(other Email) bool {
	return actual == other
}
