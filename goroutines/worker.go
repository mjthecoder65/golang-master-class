package goroutines

import (
	"fmt"
	"sync"
	"time"
)

func Worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, job)
		time.Sleep(time.Second) // Simulate the code.
		fmt.Printf("Worker %d finished job %d\n", id, job)
		results <- job * 2
	}
}

func WorkerMain() {
	const NUMBER_OF_JOBS int = 5

	jobs := make(chan int, NUMBER_OF_JOBS)
	results := make(chan int, NUMBER_OF_JOBS)

	var wg sync.WaitGroup

	for id := 1; id <= 3; id++ {
		wg.Add(1)
		go Worker(id, jobs, results, &wg)
	}

	for j := 1; j < NUMBER_OF_JOBS; j++ {
		jobs <- j
	}

	close(jobs)
	wg.Wait()
	close(results)

	for result := range results {
		fmt.Println("Result: ", result)
	}
}
