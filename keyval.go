package main

import "sync"

type KV struct {
	mu   sync.RWMutex //  is a read-write mutex, a type of lock that 
	// allows multiple goroutines to hold the lock for reading (RLock) 
	// simultaneously, but only one goroutine to hold the lock for writing (Lock)
	data map[string][]byte
}

func NewKV() *KV {
	return &KV{
		data: map[string][]byte{},
	}
}

func (kv *KV) Set(key, val []byte) error {
	kv.mu.Lock()
	defer kv.mu.Unlock()
	kv.data[string(key)] = []byte(val)
	return nil
}

func (kv *KV) Get(key []byte) ([]byte, bool) {
	kv.mu.RLock()
	defer kv.mu.RUnlock()
	val, ok := kv.data[string(key)]
	return val, ok
}
