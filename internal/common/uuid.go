package common

import googleUuid "github.com/google/uuid"

type UUID string

func (actual UUID) IsEqual(other UUID) bool {
	return actual == other
}

func CreateUUID() (UUID, error) {
	uuid, err := googleUuid.NewV7()
	if err != nil {
		return "", err
	}
	return UUID(uuid.String()), nil
}

func RestoreUUID(inputUUID string) (UUID, error) {
	err := googleUuid.Validate(inputUUID)

	if err != nil {
		return "", err
	}

	return UUID(inputUUID), nil
}
