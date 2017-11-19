package workerPool_test

import (
	"testing"

	"github.com/iamdimka/golib/workerPool"
)

func TestWorkerPool(t *testing.T) {
	pool := workerPool.New(10, 100)

	go pool.Handle(func(j workerPool.Job) {
	})

	for i := 0; i < 20000; i++ {
		go pool.AddJob(i)
	}

	pool.Stop()
}
