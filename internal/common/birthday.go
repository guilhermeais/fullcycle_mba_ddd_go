package common

import (
	"fmt"
	"time"
)

type Birthday struct {
	time.Time
	isValid bool
	clock   Clock
}

const BirthdateLayout = "2006-01-02"

func CreateBirthday(date time.Time, c Clock) (Birthday, error) {
	age := getYearOfBirth(date, c.Now())

	if age <= 0 {
		return Birthday{isValid: false, clock: c}, MakeErrInvalidBirthday(age)
	}

	return Birthday{isValid: true, Time: date, clock: c}, nil
}

func MakeErrInvalidBirthday(age int) error {
	return fmt.Errorf("%w: A idade %d é inválida. Não pode ser menor ou igual a 0", ErrValidation, age)
}

func (b Birthday) GetAge() int {
	return getYearOfBirth(b.Time, b.clock.Now())
}

func getYearOfBirth(birthdate time.Time, now time.Time) int {
	age := now.Year() - birthdate.Year()
	if now.YearDay() < birthdate.YearDay() {
		age--
	}

	return age
}
