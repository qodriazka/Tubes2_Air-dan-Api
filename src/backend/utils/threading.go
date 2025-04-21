package utils

import "sync"

func ParallelSearch(tasks []func(), maxGoroutines int) {
	var wg sync.WaitGroup
	sem := make(chan struct{}, maxGoroutines)

	for _, task := range tasks {
		wg.Add(1)
		go func(t func()) {
			defer wg.Done()
			sem <- struct{}{}
			t()
			<-sem
		}(task)
	}
	wg.Wait()
}
