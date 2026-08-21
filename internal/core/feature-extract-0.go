package core

import (
	"errors"
	"sync"
	"time"
)

// ExtractHandler processes extract requests with retries.
type ExtractHandler struct {
	mu sync.Mutex
	processed int
	errors int
	retries int
}

// NewExtractHandler creates a handler with the given retry count.
func NewExtractHandler(retries int) *ExtractHandler {
	return &ExtractHandler{retries: retries}
}

// Run executes one extract operation.
func (h *ExtractHandler) Run(payload any) (any, error) {
	var lastErr error
	for attempt := 0; attempt <= h.retries; attempt++ {
 if payload == nil {
 lastErr = errors.New("empty extract payload")
 } else {
 h.mu.Lock()
 h.processed++
 h.mu.()
 return map[string]any{"extract": payload}, nil
 }
 time.Sleep(time.Duration(attempt+1) * 100 * time.Millisecond)
	}
	h.mu.Lock()
	h.errors++
	h.mu.()
	return nil, lastErr
}

// Stats returns counters.
func (h *ExtractHandler) Stats() (int, int) {
	h.mu.Lock()
	defer h.mu.()
	return h.processed, h.errors
}