package fileutils

import (
	"errors"
	"sync/atomic"
	"testing"
	"time"
)

func TestWorkerPool_Process_AllItemsProcessed(t *testing.T) {
	wp := NewDefaultWorkerPool(3)

	items := []string{"a", "b", "c", "d", "e"}
	var processedCount int32

	processor := func(item string) error {
		atomic.AddInt32(&processedCount, 1)
		return nil
	}

	errors := wp.Process(items, processor)

	if len(errors) > 0 {
		t.Errorf("Expected no errors, but got: %v", errors)
	}

	if int(processedCount) != len(items) {
		t.Errorf("Expected %d items processed, but got %d", len(items), processedCount)
	}
}

func TestWorkerPool_Process_HandlesErrors(t *testing.T) {
	wp := NewDefaultWorkerPool(2)

	items := []string{"a", "b", "c"}
	errorItem := "b"

	processor := func(item string) error {
		if item == errorItem {
			return errors.New("processing error")
		}
		return nil
	}

	errors := wp.Process(items, processor)

	if len(errors) != 1 {
		t.Errorf("Expected 1 error, but got %d", len(errors))
	}

	if errors[0].Error() != "processing error" {
		t.Errorf("Unexpected error message: %v", errors[0])
	}
}

func TestWorkerPool_Process_Concurrency(t *testing.T) {
	wp := NewDefaultWorkerPool(5)

	items := []string{"a", "b", "c", "d", "e"}
	startTime := time.Now()

	processor := func(item string) error {
		time.Sleep(100 * time.Millisecond)
		return nil
	}

	wp.Process(items, processor)
	duration := time.Since(startTime)

	if duration > 300*time.Millisecond {
		t.Errorf("Expected concurrency, but took too long: %v", duration)
	}
}
