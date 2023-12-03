package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	const numJobs = 5
	jobs := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go func(i int) {
			worker(i, jobs)
			defer wg.Done()
		}(w)

	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)
	wg.Wait()

}

func worker(id int, jobs <-chan int) {
	for j := range jobs {
		fmt.Println("worker ", id, "started job ", j)
		time.Sleep(time.Second)
		fmt.Println("worker ", id, "finished job ", j)
	}
}
