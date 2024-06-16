package main

import (
	"fmt"
	"reflect"
	"sync"
	"time"
)

type Event struct {
	Type    string
	Payload interface{}
}

type EventService interface {
	Publish(event Event)
	Subscribe(eventType string, handler func(Event))
	UnSubscribe(eventType string, handler func(Event))
}

type SimpleEventService struct {
	subscribers map[string][]func(Event)
	mu          sync.RWMutex
}

func NewSimpleEventService() *SimpleEventService {
	return &SimpleEventService{
		subscribers: make(map[string][]func(Event)),
	}
}

func (s *SimpleEventService) Publish(event Event) {
	s.mu.RLock()
	handlers := s.subscribers[event.Type]
	s.mu.RUnlock()

	for _, handler := range handlers {
		go handler(event)
	}
}

func (s *SimpleEventService) Subscribe(eventType string, handler func(Event)) {
	s.mu.Lock()
	s.subscribers[eventType] = append(s.subscribers[eventType], handler)
	s.mu.Unlock()
}

func (s *SimpleEventService) Unsubscribe(eventType string, handler func(Event)) {
	s.mu.Lock()
	handlers := s.subscribers[eventType]
	for i, h := range handlers {
		if reflect.ValueOf(h).Pointer() == reflect.ValueOf(handler).Pointer() {
			s.subscribers[eventType] = append(handlers[:i], handlers[i+1:]...)
			break
		}
	}
	s.mu.Unlock()
}

func main() {
	bus := NewSimpleEventService()

	// Define a subscriber
	handler := func(event Event) {
		fmt.Printf("Received event: %v\n", event)
	}

	// Subscribe to an event type
	bus.Subscribe("user.created", handler)

	// Publish an event
	event := Event{
		Type:    "user.created",
		Payload: map[string]string{"username": "johndoe"},
	}
	bus.Publish(event)

	// Unsubscribe from the event type
	bus.Unsubscribe("user.created", handler)
	time.Sleep(1 * time.Second)
}
