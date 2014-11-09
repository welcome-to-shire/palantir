// Storage interface
package main

type Storage interface {
	// Create a message.
	Create(subject, title, content string) (string, error)

	// Get and unqueue a message.
	GetOne(subject string) (*Message, error)
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
