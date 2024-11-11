package main

import "fmt"

//==============================================
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, j)
		results <- j * 2
	}
}

//==============================================

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// Start workers
	numberOfWorkers := 3
	for w := 1; w <= numberOfWorkers; w++ {
		go worker(w, jobs, results)
	}

	// Send jobs
	numberOfJobs := 5
	for j := 1; j <= numberOfJobs; j++ {
		jobs <- j
	}
	close(jobs) // Close jobs channel after sending all jobs

	// Collect results
	for a := 1; a <= 5; a++ {
		<-results
	}
}

//==============================================
