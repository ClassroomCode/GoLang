package main

import "sync"

// section: embed
type Hits struct {
	sync.Mutex // embedding will promote all public methods
	n          int
}

func (h *Hits) Inc() {
	h.Lock()
	defer h.Unlock()
	h.n++
}

// section: embed
