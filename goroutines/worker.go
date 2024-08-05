package goroutines

import (
	"fmt"
	"sync"
	"time"
)

/*
Here is the worker, of which we will run several concurrent instances.
These workers will receive work on the jobs channel and send the the corresponding
results on results.
We will sleep a second to simulate an expensive task.
*/
func Worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, job)
		time.Sleep(time.Second) // Simulate an expensive job.
		fmt.Printf("Worker %d finished job %d\n", id, job)
		results <- job * 2
	}
}

func WorkerMain() {
	const NUMBER_OF_JOBS int = 5
	/*
		In order to use our pool of workers, we need to send them work and
		collect their results. We will make two channels for this one.
	*/
	jobs := make(chan int, NUMBER_OF_JOBS)
	results := make(chan int, NUMBER_OF_JOBS)

	var wg sync.WaitGroup

	// Start three workers, initially blocked before there are not jobs yet.
	for id := 1; id <= 3; id++ {
		wg.Add(1)
		go Worker(id, jobs, results, &wg)
	}

	// Here we send five jobs and then close that channel to indicate that
	// that's all the work we have.
	for j := 1; j < NUMBER_OF_JOBS; j++ {
		jobs <- j
	}
	close(jobs)

	// Finally, we collect results of the work. This ensure that the work
	// goroutine has finished.
	// An alternative to way to wait for multiple goroutines to is use a WaitGroup.
	for result := range results {
		fmt.Println("Result: ", result)
	}
}
