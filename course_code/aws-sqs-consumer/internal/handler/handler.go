package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
)

type Config struct {

}

type Handler struct {
	config Config
}

func NewHandler(c Config) *Handler {
	return &Handler{config: c}
}

type MessageType string

const (
	CreateCard MessageType = "CreateCard"
	DeleteCard MessageType = "DeleteCard"
)

type Message struct {
	Type MessageType
	Data map[string]interface{}
}

func (h *Handler) Handle(ctx context.Context, m string) error {
	log.Printf("Received a new message: '%v'", m)

	var message Message

	if err := json.Unmarshal([]byte(m), &message); err != nil {
		return err
	}

	switch (message.Type) {
	case CreateCard:
		log.Print("Creating Card")
	case DeleteCard:
		log.Print("Deleting Card")
	default:
		return fmt.Errorf("Invalid Message Type: %v", message.Type)
	}

	return nil
}