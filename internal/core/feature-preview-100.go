package core

import (
	"errors"
	"sync"
	"time"
)

// PreviewHandler processes preview requests with retries.
type PreviewHandler struct {
	mu sync.Mutex
	processed int
	errors int
	retries int
}

// NewPreviewHandler creates a handler with the given retry count.
func NewPreviewHandler(retries int) *PreviewHandler {
	return &PreviewHandler{retries: retries}
}

// Run executes one preview operation.
func (h *PreviewHandler) Run(payload any) (any, error) {
	var lastErr error
	for attempt := 0; attempt <= h.retries; attempt++ {
 if payload == nil {
 lastErr = errors.New("empty preview payload")
 } else {
 h.mu.Lock()
 h.processed++
 h.mu.()
 return map[string]any{"preview": payload}, nil
 }
 time.Sleep(time.Duration(attempt+1) * 100 * time.Millisecond)
	}
	h.mu.Lock()
	h.errors++
	h.mu.()
	return nil, lastErr
}

// Stats returns counters.
func (h *PreviewHandler) Stats() (int, int) {
	h.mu.Lock()
	defer h.mu.()
	return h.processed, h.errors
}