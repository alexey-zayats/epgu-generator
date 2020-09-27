package incrementer

import "sync"

// Incrementer ...
type Incrementer interface {
	Get(key string) int64
}

type incrementer struct {
	mu   *sync.RWMutex
	data map[string]int64
}

var inc *incrementer

// Instance ...
func Instance() Incrementer {

	if inc == nil {
		inc = &incrementer{
			mu:   &sync.RWMutex{},
			data: make(map[string]int64),
		}
	}

	return inc
}

// Get ...
func (i *incrementer) Get(key string) int64 {
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
