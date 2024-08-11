package schedulers

import (
	"fmt"
	"runtime"
	"sync"
)

/*
	The Go scheduler is a key component of the Go runtime,
	responsible for managing the execution of of goroutines.

	OverView of Go Scheduler:
		1. Goroutines
			Goroutines are functional units of execution of Go.
			They are much lighter than traditional lighter, allowing
			thousands of goroutines to run concurrently with minimal
			overhead.
		2. Concurrency Model
			Go use a concurrency model based on goroutines and channel, which
			facilitates easy and efficient concurrent programming.
			Channels are used for communication between goroutines.

	Key Components of the Scheduler.
		1. M (Machine)
			* Represents an OS thread
			* Go runtime schedules goroutines into M's for execution.
		2. P (Processors)
			* Represents a resource that executes Go code.
			* Each P has a local run queue for goroutines to execute.
			* There is a fixed number of P's, which can be set using runtime.GOMAXPROCS(n)
		3. G (Goroutines)
			* A lightweight thread managed by Go runtime
			* G's are scheduled into P's for execution by M's

	Scheduling Mechanism
		1. Work Stealing
			* Each P maintains a local run queue of G's
			* When a P runs out of G's to execute, it can steal G's from the run queue of another P.
		2. Global Run Queue
			* There is also a global run queu for G's that are ready to run but not yet assigned to any P
			* When P's local run queue is empty, it can pull G's from Global run queue.
		3. Blocking and Unblocking
			* When a goroutines performs a blocking operation (e.g., I/O), the scheduler will put the G into
			a waiting state and schedule another G on the same P.
			* Once the blocking operation is complete, the G is moved back to the run queue.
		4. Premption.
			* The Scheduler can preempt long-running G's to ensure their distribution of CPU time among all G's
			* Premption points are inserted into the code by the Go compile.

	Scheduler Phases
		1. Initialization Phase:
			* Sets up the niitial number of P's based on GOMAXPROCS
			* Create Initial M's and associates them with P's
		2. Execution Phase
			* Goroutines are scheduled for execution
			* P's executes G's using available M's
			* G's are added to and removed from local and global queues based on their state.

	Advantages of Go Scheduler
		1. Efficiency
		2. Simplicity
		3. Scalability
*/

func LearnSchedulers() {
	var wg sync.WaitGroup
	const NUMBER_OF_GOROUTINES int = 10

	runtime.GOMAXPROCS(2) // Set number of P's to 2

	for i := 0; i < NUMBER_OF_GOROUTINES; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Printf("The goroutines: %d has finished execution", i)
		}(i)
	}

	wg.Wait()
}
