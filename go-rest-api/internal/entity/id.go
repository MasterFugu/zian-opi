package entity

import "github.com/google/uuid"

// GenerateID generates a unique ID that can be used as an identtestier for an entity.
func GenerateID() string {
	return uuid.New().String()
}
