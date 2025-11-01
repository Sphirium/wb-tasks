package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	count int
	mu    sync.Mutex
}

func (c *Counter) Increment() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

func (c *Counter) GetInfo() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func main() {
	resCount := Counter{
		mu: sync.Mutex{},
	}

	var wg sync.WaitGroup

	for range 20 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			resCount.Increment()

		}()
	}

	wg.Wait()
	fmt.Printf("Было запущено %v горутин", resCount.GetInfo())

}
