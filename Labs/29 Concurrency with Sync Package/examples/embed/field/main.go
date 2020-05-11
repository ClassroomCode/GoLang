package main

import "sync"

// section: embed
type Hits struct {
	mu sync.Mutex // methods are no longer promoted and only accessible within this package now.
	n  int
}

func (h *Hits) Inc() {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.n++
}

// section: embed
