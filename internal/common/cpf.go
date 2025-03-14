package common

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type CPF string

func (actual CPF) IsEqual(other CPF) bool {
	return actual == other
}

var knownInvalidCPFs = [...]string{
	"00000000000",
	"11111111111",
	"22222222222",
	"33333333333",
	"44444444444",
	"55555555555",
	"66666666666",
	"77777777777",
	"88888888888",
	"99999999999",
}

func CreateCPF(rawInput string) (CPF, error) {
	inputOnlyDigits := extractDigits(rawInput)

	if len(inputOnlyDigits) != 11 {
		return "", ErrInvalidCPF
	}

	if isKnownInvalidCPF(inputOnlyDigits) {
		return "", ErrInvalidCPF
	}

	if !isValidCPF(inputOnlyDigits) {
		return "", ErrInvalidCPF
	}

	return CPF(inputOnlyDigits), nil
}

func extractDigits(input string) string {
	re := regexp.MustCompile(`\d`)
	digits := re.FindAllString(input, -1)
	return strings.Join(digits, "")
}

func isKnownInvalidCPF(cpf string) bool {
	for _, knownInvalidCPF := range knownInvalidCPFs {
		if cpf == knownInvalidCPF {
			return true
		}
	}
	return false
}

// isValidCPF validates the CPF using the verification digits.
func isValidCPF(cpf string) bool {
	expectedFirstVerifier, err := calculateVerifierDigit(cpf, 9, 10)
	if err != nil {
		return false
	}
	expectedSecondVerifier, err := calculateVerifierDigit(cpf, 10, 11)
	if err != nil {
		return false
	}

	receivedSecondVeifier, err := strconv.Atoi(string(cpf[10]))
	if err != nil {
		return false
	}

	return expectedFirstVerifier == int(cpf[9]-'0') && expectedSecondVerifier == receivedSecondVeifier
}

func calculateVerifierDigit(cpf string, length int, weight int) (int, error) {
	sum := 0
	for i := 0; i < length; i++ {
		num, err := strconv.Atoi(string(cpf[i]))
		if err != nil {
			return 0, err
		}
		sum += num * (weight - i)
	}
	verifier := (sum * 10) % 11
	if verifier == 10 {
		verifier = 0
	}
	return verifier, nil
}

var ErrInvalidCPF = fmt.Errorf("%w: CPF Inválido", ErrValidation)
