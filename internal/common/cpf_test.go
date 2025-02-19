package common_test

import (
	"fmt"
	"testing"

	"ingressos/internal/common"
)

func TestCreateCPFWithInvalidCPFs(t *testing.T) {
	invalidCPFCases := []struct {
		Input string
		Error error
	}{
		{Input: "00000000000", Error: common.ErrInvalidCPF},
		{Input: "11111111111", Error: common.ErrInvalidCPF},
		{Input: "1234567890", Error: common.ErrInvalidCPF},
		{Input: "123456789012", Error: common.ErrInvalidCPF},
	}

	for _, test := range invalidCPFCases {
		t.Run(fmt.Sprintf("input %s should return error %v", test.Input, test.Error), func(t *testing.T) {
			_, err := common.CreateCPF(test.Input)

			if err == nil || err.Error() != test.Error.Error() {
				t.Fatalf("expected error %v but received %v", test.Error, err)
			}
		})
	}
}

func TestCreateCPFWithValidCPFs(t *testing.T) {
	validCPFCases := []struct {
		Input, Expected string
	}{
		{Input: "444.074.338-25", Expected: "44407433825"},
		{Input: "123.456.789-09", Expected: "12345678909"},
	}

	for _, test := range validCPFCases {
		t.Run(fmt.Sprintf("input %s should be valid", test.Input), func(t *testing.T) {
			cpf, err := common.CreateCPF(test.Input)

			if err != nil {
				t.Fatalf("expected no error but received %v", err)
			}

			if string(cpf) != test.Expected {
				t.Fatalf("expected CPF %s but received %s", test.Expected, cpf)
			}
		})
	}
}

func TestIsEq(t *testing.T) {
	rawCpf := "44407433825"
	cpf, _ := common.CreateCPF(rawCpf)
	sameCpf, _ := common.CreateCPF(rawCpf)

	if !cpf.IsEqual(sameCpf) {
		t.Fatalf("cpf %s should be equal to %s", cpf, sameCpf)
	}
}
