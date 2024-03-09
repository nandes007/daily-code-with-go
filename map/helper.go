package main

import "sync"

func getAllKeys(sample map[string]string) []string {
	var keys []string
	for k := range sample {
		keys = append(keys, k)
	}
	return keys
}

var (
	allData = make(map[string]string)
	rwm     sync.RWMutex
)

func get(key string) string {
	rwm.RLock()
	defer rwm.RUnlock()
	return allData[key]
}

func set(key string, value string) {
	rwm.Lock()
	defer rwm.Unlock()
	allData[key] = value
}
