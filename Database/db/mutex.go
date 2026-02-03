package db

import (
	"os"
	"sync"
)

func (d *Driver) getOrCreateMutex(collection string) *sync.Mutex {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	if m, ok := d.mutexes[collection]; ok {
		return m
	}

	m := &sync.Mutex{}
	d.mutexes[collection] = m
	return m
}

func stat(path string) (os.FileInfo, error) {
	if fi, err := os.Stat(path); err == nil {
		return fi, nil
	}
	return os.Stat(path + ".json")
}
