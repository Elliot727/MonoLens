package fileutils

import (
	"sync"
)

type WorkerPool interface {
	Process(items []string, processor func(string) error) []error
}

type DefaultWorkerPool struct {
	numWorkers int
}

func NewDefaultWorkerPool(numWorkers int) *DefaultWorkerPool {
	return &DefaultWorkerPool{
		numWorkers: numWorkers,
	}
}

func (wp *DefaultWorkerPool) Process(items []string, processor func(string) error) []error {
	var wg sync.WaitGroup
	itemChan := make(chan string, len(items))
	errChan := make(chan error, len(items))

	worker := func() {
		for item := range itemChan {
			if err := processor(item); err != nil {
				errChan <- err
			}
		}
		wg.Done()
	}

	for i := 0; i < wp.numWorkers; i++ {
		wg.Add(1)
		go worker()
	}

	for _, item := range items {
		itemChan <- item
	}
	close(itemChan)

	wg.Wait()
	close(errChan)

	var errors []error
	for err := range errChan {
		errors = append(errors, err)
	}

	return errors
}
