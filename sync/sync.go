package sync

import (
	"fmt"
	"sync"
	"time"
)

/*
Provides basic synchronization primitives such:

 1. Mutual exclusion locks (mutexes)
    var mu sync.Mutex: Used to provide mutual exclusion. It ensures that only one
    goroutine can access a critical part of code at time.

    Important method:

    mu.Lock(): Acquires the lock. If the lock is already held by another goroutine,
    the current goroutine will block until the lock is available. Ensure only one goroutines
	can access shared resources in a go program.

    mu.Unlock(): Releases the lock. it is important to always unlock after a lock
    is acquired, even if an error occurs.

    var rwmu sync.RWMutex: Allows multiple readers or one writer but not both at the same time.

    rwmu.RLock(): Acquire read locks, allowing other readers but blocking writers.
    rwmu.RUnlock(): Release the read lock.
    rwmu.Lock(): Acquire write lock, blocking other readers and writers.
    rwmu.Unlock(): Release the writer lock.

 2. WaitGroups
    A WaitGroup waits for a collection of goroutines to finish executing.
    It's particularly useful when you need to wait for multiple goroutines to complete
    before proceeding.
    Key Methods:

    Add(delta int):
    Adds to the WaitGroup Counter. If the counter,
    becomes zeros, all goroutines blocked on wait are released.

    Done(): Decrements the Waitgroup counter by one. It usually called with defer statement in the goroutine.
    Wait(): blocks the caller goroutines until the WaitGroup count is zero. This is typically the main goroutine,
    but it doesn't have to.

 3. once (for one time initialization)
*/

func Worker(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	fmt.Printf("INFO: Worker[%d] Starting...", id)
	time.Sleep(time.Second)
	fmt.Printf("INFO: Worker[%d] Finished", id)
}

func LearnSyncPackage() {

	var wg *sync.WaitGroup // Waits for a collection of goroutines to finish.
	const NUMBER_OF_WORKERS int = 10

	for id := 0; id < NUMBER_OF_WORKERS; id++ {
		wg.Add(1)
		go Worker(wg, id)
	}

	fmt.Println("Waiting for workers to finish")
	wg.Wait()
	fmt.Println("All Workers are done!")
}
