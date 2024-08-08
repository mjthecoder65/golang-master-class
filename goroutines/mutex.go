package goroutines

import (
	"fmt"
	"sync"
)

var mu sync.Mutex

type Container struct {
	counters map[string]int
}

// Increment the counters.
func (con *Container) Increment(name string) {
	mu.Lock()
	defer mu.Unlock()
	con.counters[name]++
}

/*
	Handling Race Condition is very important.
*/

func LearnMutex() {

	cont := Container{
		counters: map[string]int{"a": 0, "b": 0},
	}

	doIncrement := func(name string, n int, wg *sync.WaitGroup) {
		for i := 0; i < n; i++ {
			cont.Increment(name)
		}
		wg.Done()
	}

	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(2)
		go doIncrement("a", 10, &wg)
		go doIncrement("b", 10, &wg)
	}

	wg.Wait() // Wait for all goroutines to finish and then print results.
	fmt.Println(cont.counters)
}
