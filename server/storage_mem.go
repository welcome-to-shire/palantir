// In memory storage
package main

import (
	"errors"
	"time"

	"code.google.com/p/go-uuid/uuid"
	palantir "github.com/welcome-to-shire/palantir-go"
)

type MemStorage struct {
	notifications map[string]map[string]*palantir.Message
}

func MakeMemStorage() *MemStorage {
	s := new(MemStorage)
	s.notifications = make(map[string]map[string]*palantir.Message)

	return s
}

func (m *MemStorage) getSubjectInbox(subject string) map[string]*palantir.Message {
	box, found := m.notifications[subject]
	if !found {
		m.notifications[subject] = make(map[string]*palantir.Message)
		box = m.notifications[subject]
	}

	return box
}

func (m *MemStorage) Create(subject, title, content string) (string, error) {
	if _, found := m.notifications[subject]; !found {
		m.notifications[subject] = make(map[string]*palantir.Message)
	}

	ticket := generateTicket()
	m.notifications[subject][ticket] = &palantir.Message{
		title,
		content,
		palantir.ISO8601Time(time.Now()),
	}
	return ticket, nil
}

func (m *MemStorage) GetOne(subject string) (*palantir.Message, error) {
	box, found := m.notifications[subject]
	if !found {
		return nil, errors.New("no new notification")
	}
	if len(box) == 0 {
		return nil, errors.New("no new notification")
	}

	var ticket string
	for i := range box {
		ticket = i
		break
	}
	message := box[ticket]
	delete(m.notifications[subject], ticket)

	return message, nil
}

func generateTicket() string {
	return uuid.New()
}
