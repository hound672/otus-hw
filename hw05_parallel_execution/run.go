package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var (
	ErrErrorsLimitExceeded = errors.New("errors limit exceeded")
	ErrNilTask             = errors.New("task is nil")
)

type Task func() error

// Run starts tasks in N goroutines and stops its work when receiving M errors from tasks.
func Run(tasks []Task, n int, m int) error {
	workerTasks := make(chan Task)
	var errorsCount int32

	wg := sync.WaitGroup{}
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			for task := range workerTasks {
				err := task()
				if err != nil {
					atomic.AddInt32(&errorsCount, 1)
				}
			}
		}()
	}

	err := func() error {
		for _, task := range tasks {
			task := task
			if task == nil {
				return ErrNilTask
			}

			if atomic.LoadInt32(&errorsCount) >= int32(m) {
				return ErrErrorsLimitExceeded
			}
			workerTasks <- task
		}

		return nil
	}()

	close(workerTasks)
	wg.Wait()

	return err
}
