package hw05parallelexecution

import (
	"context"
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	workerTasks := make(chan Task, len(tasks))
	errorTasks := make(chan error, len(tasks))
	var errorsCount int

	// spawn workers
	wg.Add(n)
	for i := 0; i < n; i++ {
		go worker(ctx, &wg, workerTasks, errorTasks)
	}

	for _, task := range tasks {
		workerTasks <- task
	}

	for _ = range errorTasks {
		errorsCount++
		if errorsCount >= m {
			break
		}
	}
	cancel()

	close(workerTasks)
	wg.Wait()
	close(errorTasks) // close this channel only after working (to prevent writing close channel panic)

	return nil
}

func worker(ctx context.Context, wg *sync.WaitGroup, workerTasks <-chan Task, errorsTasks chan<- error) {
	defer wg.Done()

	for task := range workerTasks {
		if ctx.Err() != nil {
			return
		}

		err := task()
		if err != nil {
			errorsTasks <- err
		}
	}

	// the second solution
	// but the first one i like more since there is less the nesting
	//for {
	//	select {
	//	case <-ctx.Done():
	//		// cancel working
	//		return
	//	case task, ok := <-workerTasks:
	//		if !ok {
	//			// channel is closed
	//			return
	//		}
	//
	//		_ = task()
	//	}
	//}
}
