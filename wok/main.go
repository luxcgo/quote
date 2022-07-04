package main

import (
	"fmt"
	"sync"
)

func main() {
	nbJobs := 30
	concurrentGoroutines := make(chan struct{}, 2)

	var wg sync.WaitGroup

	for i := 0; i < nbJobs; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			concurrentGoroutines <- struct{}{}
			worker(i)
			<-concurrentGoroutines

		}(i)

	}
	wg.Wait()

}

func worker(i int) { fmt.Println("doing work on", i) }
