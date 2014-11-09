// Message type.
package main

import (
	"fmt"
	"time"
)

type MessageTime time.Time

type Message struct {
	Title     string      `json:"title"`
	Content   string      `json:"content"`
	CreatedAt MessageTime `json:"created_at"`
}

func (t MessageTime) MarshalJSON() ([]byte, error) {
	timeString := fmt.Sprintf(
		"\"%s\"",
		time.Time(t).Format(time.RFC3339Nano),
	)
	return []byte(timeString), nil
}
