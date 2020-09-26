package artefact

import "sync"

// Incrementer ...
type Incrementer struct {
	mu   *sync.RWMutex
	data map[string]int64
}

// NewIncrementer ...
func NewIncrementer() *Incrementer {
	return &Incrementer{
		mu:   &sync.RWMutex{},
		data: make(map[string]int64),
	}
}

// Get ...
func (i *Incrementer) Get(key string) int64 {
	i.mu.Lock()
	defer i.mu.Unlock()

	value, ok := i.data[key]
	if ok == false {
		value = 1
	} else {
		value++
	}

	i.data[key] = value

	return value
}
