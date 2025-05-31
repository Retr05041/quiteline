package backend

import (
	"log"
)

type BackendEvent struct {
	Type string
	Data interface{}
}

type Backend struct {
	Events chan BackendEvent
}

func NewBackend() *Backend {
	log.Println("BACKEND: CREATED")
	return &Backend{
		Events: make(chan BackendEvent),
	}
}

// Send an event to the adapters to update the UI
func (b *Backend) emitEvent(eventType string, data interface{}) {
	go func() {
		b.Events <- BackendEvent{Type: eventType, Data: data}
	}()
}

func (b *Backend) SendMessage(newMsg string) {
	log.Printf("BACKEND: SendMessage was called: %s\n", newMsg)
	b.emitEvent("LOG", "Thanks for the message!")
}
