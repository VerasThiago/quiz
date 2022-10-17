package utils

import "github.com/google/uuid"

func IsUUID(id string) bool {
	_, err := uuid.Parse(id)
	return err == nil
}
