package common_test

import (
	"ingressos/internal/common"
	"testing"

	"github.com/google/uuid"
)

func TestCreateUUID(t *testing.T) {
	u, err := common.CreateUUID()
	if err != nil {
		t.Fatalf("non expected error: %v", err)

	}

	err = uuid.Validate(string(u))

	if err != nil {
		t.Fatalf("should create a valid uuid: %v", err)
	}
}

func TestRestoreUUID(t *testing.T) {
	t.Run("should return error for invalid uuid", func(t *testing.T) {
		input := "invalid uuid"
		_, err := common.RestoreUUID(input)

		if err == nil {
			t.Fatalf("uuid %s should be invalid: %v", input, err)
		}
	})

	t.Run("should return UUID if UUID is valid", func(t *testing.T) {
		input := "94d884bd-11ba-4a87-80e1-2732c5164bc6"
		u, err := common.RestoreUUID(input)

		if err != nil {
			t.Fatalf("uuid %s should be valid: %v", input, err)
		}

		if string(u) != input {
			t.Fatalf("uuid shoud be equal to input uuid when comparing primitives")
		}
	})
}
