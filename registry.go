package main

import (
	"log"
	"sync"
)

type Registry struct {
	mu      sync.RWMutex
	clients map[*Client]struct{}
}

func NewRegistry() *Registry {
	return &Registry{clients: make(map[*Client]struct{})}
}

func (r *Registry) Add(c *Client) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.clients[c] = struct{}{}
}

func (r *Registry) Remove(c *Client) {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.clients, c)
}

func (r *Registry) Broadcast(msg string, except *Client) int {
	r.mu.RLock()
	defer r.mu.RUnlock()

	delivered := 0
	for c := range r.clients {
		if c == except {
			continue
		}
		if err := c.Send(msg); err != nil {
			log.Printf(sendFailedFormat, c.addr, err)
			continue
		}
		delivered++
	}
	return delivered
}
