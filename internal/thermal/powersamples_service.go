package thermal

import "sync"

type PowerSamplesHistory struct {
	mu     sync.RWMutex
	values []int
}

func (h *PowerSamplesHistory) Add(value int) {
	h.mu.Lock()
	h.values = append(h.values, value)
	h.mu.Unlock()
}
func (h *PowerSamplesHistory) Values() []int {
	h.mu.RLock()
	defer h.mu.RUnlock()
	out := make([]int, len(h.values))
	copy(out, h.values)
	return out
}
