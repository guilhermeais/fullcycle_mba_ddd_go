package common

import (
	"fmt"
	"time"
)

type Birthday struct {
	date    time.Time
	isValid bool
	clock   Clock
}

func CreateBirthday(date time.Time, c Clock) (Birthday, error) {
	age := getYearOfBirth(date, c.Now())

	if age <= 0 {
		return Birthday{isValid: false, clock: c}, MakeErrInvalidBirthday(age)
	}

	return Birthday{isValid: true, date: date, clock: c}, nil
}

func MakeErrInvalidBirthday(age int) error {
	return fmt.Errorf("%w: A idade %d é inválida. Não pode ser menor ou igual a 0", ErrValidation, age)
}

func (b Birthday) GetAge() int {
	return getYearOfBirth(b.date, b.clock.Now())
}

func getYearOfBirth(birthdate time.Time, now time.Time) int {
	age := now.Year() - birthdate.Year()
	if now.YearDay() != birthdate.YearDay() {
		age--
	}

	return age
}
