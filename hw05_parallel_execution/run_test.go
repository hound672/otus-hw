package hw05parallelexecution

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/goleak"
)

func TestRun(t *testing.T) {
	defer goleak.VerifyNone(t)

	t.Run("if were errors in first M tasks, than finished not more N+M tasks", func(t *testing.T) {
		tasksCount := 50
		tasks := make([]Task, 0, tasksCount)

		var runTasksCount int32

		for i := 0; i < tasksCount; i++ {
			err := fmt.Errorf("error from task %d", i)
			tasks = append(tasks, func() error {
				time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
				atomic.AddInt32(&runTasksCount, 1)
				return err
			})
		}

		workersCount := 10
		maxErrorsCount := 23
		err := Run(tasks, workersCount, maxErrorsCount)

		require.Truef(t, errors.Is(err, ErrErrorsLimitExceeded), "actual err - %v", err)
		require.LessOrEqual(t, runTasksCount, int32(workersCount+maxErrorsCount), "extra tasks were started")
	})

	t.Run("tasks without errors", func(t *testing.T) {
		tasksCount := 50
		tasks := make([]Task, 0, tasksCount)

		var runTasksCount int32
		var sumTime time.Duration

		for i := 0; i < tasksCount; i++ {
			taskSleep := time.Millisecond * time.Duration(rand.Intn(100))
			sumTime += taskSleep

			tasks = append(tasks, func() error {
				time.Sleep(taskSleep)
				atomic.AddInt32(&runTasksCount, 1)
				return nil
			})
		}

		workersCount := 5
		maxErrorsCount := 1

		start := time.Now()
		err := Run(tasks, workersCount, maxErrorsCount)
		elapsedTime := time.Since(start)
		require.NoError(t, err)

		require.Equal(t, runTasksCount, int32(tasksCount), "not all tasks were completed")
		require.LessOrEqual(t, int64(elapsedTime), int64(sumTime/2), "tasks were run sequentially?")
	})

	t.Run("allowed zero errors", func(t *testing.T) {
		tasksCount := 50
		tasks := make([]Task, 0, tasksCount)

		var runTasksCount int32

		for i := 0; i < tasksCount; i++ {
			i := i

			err := fmt.Errorf("error from task %d", i)
			tasks = append(tasks, func() error {
				atomic.AddInt32(&runTasksCount, 1)
				return err
			})
		}

		workersCount := 10
		maxErrorsCount := 0
		result := Run(tasks, workersCount, maxErrorsCount)

		require.Equal(t, ErrErrorsLimitExceeded, result)
		require.LessOrEqual(t, runTasksCount, int32(workersCount+maxErrorsCount), "extra tasks were started")
	})
}

func TestNilTasks(t *testing.T) {
	defer goleak.VerifyNone(t)

	t.Run("single task nil", func(t *testing.T) {
		taskCount := 1
		tasks := make([]Task, 0, taskCount)

		for i := 0; i < taskCount; i++ {
			tasks = append(tasks, nil)
		}

		workersCount := 10
		maxErrorsCount := taskCount
		result := Run(tasks, workersCount, maxErrorsCount)

		require.Equal(t, ErrNilTask, result)
	})

	t.Run("some tasks is nil", func(t *testing.T) {
		tasksCount := 50
		tasks := make([]Task, 0, tasksCount)

		var runTasksCount int32
		var sumTime time.Duration

		for i := 0; i < tasksCount; i++ {
			taskSleep := time.Millisecond * time.Duration(rand.Intn(100))
			sumTime += taskSleep

			if i%2 == 0 {
				tasks = append(tasks, nil)
				continue
			}

			tasks = append(tasks, func() error {
				time.Sleep(taskSleep)
				atomic.AddInt32(&runTasksCount, 1)
				return nil
			})
		}

		workersCount := 5
		maxErrorsCount := 1

		result := Run(tasks, workersCount, maxErrorsCount)

		require.Equal(t, ErrNilTask, result)
	})
}

func TestWithEventually(t *testing.T) {
	defer goleak.VerifyNone(t)

	t.Run("with eventually (positive)", func(t *testing.T) {
		tasksCount := 50
		tasks := make([]Task, 0, tasksCount)

		var runTasksCount int32
		var sumTime time.Duration

		for i := 0; i < tasksCount; i++ {
			taskSleep := time.Millisecond * time.Duration(rand.Intn(100))
			sumTime += taskSleep

			tasks = append(tasks, func() error {
				time.Sleep(taskSleep)
				atomic.AddInt32(&runTasksCount, 1)
				return nil
			})
		}

		workersCount := 5
		maxErrorsCount := 1

		var wg sync.WaitGroup

		wg.Add(1)
		go func() {
			defer wg.Done()
			condition := func() bool {
				finishedTasksCount := atomic.LoadInt32(&runTasksCount)
				return finishedTasksCount >= int32(tasksCount)
			}
			require.Eventually(t, condition, sumTime/2, time.Millisecond, "tasks were run sequentially?")
		}()

		err := Run(tasks, workersCount, maxErrorsCount)
		require.NoError(t, err)

		wg.Wait()
	})

	t.Run("with eventually (negative)", func(t *testing.T) {
		tasksCount := 50
		tasks := make([]Task, 0, tasksCount)

		var runTasksCount int32
		var sumTime time.Duration

		for i := 0; i < tasksCount; i++ {
			taskSleep := time.Millisecond * time.Duration(rand.Intn(100))
			sumTime += taskSleep

			tasks = append(tasks, func() error {
				time.Sleep(taskSleep)
				atomic.AddInt32(&runTasksCount, 1)
				return nil
			})
		}

		workersCount := 2
		maxErrorsCount := 1

		var wg sync.WaitGroup

		wg.Add(1)
		go func() {
			defer wg.Done()
			condition := func() bool {
				finishedTasksCount := atomic.LoadInt32(&runTasksCount)
				return finishedTasksCount >= int32(tasksCount)
			}
			require.Never(t, condition, sumTime/2, time.Millisecond, "tasks were run sequentially?")
		}()

		err := Run(tasks, workersCount, maxErrorsCount)
		require.NoError(t, err)

		wg.Wait()
	})
}
