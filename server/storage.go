// Storage interface
package main

import (
	palantir "github.com/welcome-to-shire/palantir-go"
)

type Storage interface {
	// Create a message.
	Create(subject, title, content string) (string, error)

	// Get and unqueue a message.
	GetOne(subject string) (*palantir.Message, error)
}

var instance Storage

func ensureHaveStorageInstance() {
	if instance == nil {
		instance = MakeMemStorage()
	}
}

func GetStorageInstance() Storage {
	ensureHaveStorageInstance()

	return instance
}
