package common_test

import (
	"ingressos/internal/common"
	"testing"
)

func TestCreateHashedPassword(t *testing.T) {
	plain := "testing123"
	hashed, err := common.CreateHashedPassword(plain)
	if err != nil {
		t.Fatalf("does not expect the error %v", err)
	}

	comparingCorrectPass := hashed.Compare(plain)
	if !comparingCorrectPass {
		t.Fatal("returned false when comparing the correct password, expected true")
	}

	comparingIncorrectPass := hashed.Compare("invalid-pass")
	if comparingIncorrectPass {
		t.Fatal("returned true when comparing the incorrect password, expected false")
	}
}
