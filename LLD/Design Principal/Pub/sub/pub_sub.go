package main

import (
	"fmt"
	"sync"
)

type Agent struct {
	mu     *sync.Mutex
	subs   map[string][]chan string
	quit   chan struct{}
	closed bool
}

func NewAgent() *Agent {
	return &Agent{
		mu:   &sync.Mutex{},
		subs: make(map[string][]chan string),
		quit: make(chan struct{}),
	}
}

func (b *Agent) Publish(topic string, msg string) {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.closed {
		return
	}

	for _, ch := range b.subs[topic] {
		ch <- msg
	}
}

func (b *Agent) Subscribe(topic string) <-chan string {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.closed {
		return nil
	}

	ch := make(chan string)
	b.subs[topic] = append(b.subs[topic], ch)
	return ch
}

func (b *Agent) Close() {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.closed {
		return
	}

	b.closed = true
	close(b.quit)

	for _, ch := range b.subs {
		for _, sub := range ch {
			close(sub)
		}
	}
}

func main() {
	// Create a new agent
	agent := NewAgent()

	// Subscribe to a topic
	sub := agent.Subscribe("foo")

	// Publish a message to the topic
	go agent.Publish("foo", "hello world")

	// Print the message
	fmt.Println(<-sub)

	// Close the agent
	agent.Close()
}
