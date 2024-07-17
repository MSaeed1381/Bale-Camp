package main

import (
	"sync"
	"time"
)

type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Inc() {
	c.mu.Lock()
	c.value++
	c.mu.Unlock()
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func Solution(d time.Duration, message string, ch ...chan string) (numberOfAccesses int) {
	counter := &Counter{} // safe concurrent counter
	wg := new(sync.WaitGroup)
	wg.Add(len(ch)) // start len(ch) (number of channels) goroutines

	for _, c := range ch {
		go func(channel chan string) {
			defer wg.Done()
			select {
			case channel <- message: // every goroutine write in a channel
				counter.Inc()
				return
			case <-time.After(d * time.Second): // timeout
				return
			}
		}(c) // variable shadow
	}

	wg.Wait()
	return counter.Value()
}
