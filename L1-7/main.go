package main

import (
	"fmt"
	"sync"
	"time"
)

type DataMap struct {
	data map[string]int
	mu   sync.Mutex
}

func NewDataMap() *DataMap {
	return &DataMap{
		data: map[string]int{},
	}

}

func main() {
	dataMap := NewDataMap()
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := fmt.Sprintf("key-%d", i%3)
			dataMap.Set(key, i)
			time.Sleep(30 * time.Millisecond)
		}(i)
	}
	wg.Wait()

	for i := 0; i < 3; i++ {
		key := fmt.Sprintf("key-%d", i)
		if value, ok := dataMap.Get(key); ok {
			fmt.Printf("%s: %d\n", key, value)
		}

	}

}

func (dm *DataMap) Set(key string, value int) {
	dm.mu.Lock()
	defer dm.mu.Unlock()
	dm.data[key] = value
}

func (dm *DataMap) Get(key string) (int, bool) {
	dm.mu.Lock()
	defer dm.mu.Unlock()
	value, ok := dm.data[key]
	return value, ok
}
