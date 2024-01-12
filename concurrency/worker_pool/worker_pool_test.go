package worker_pool

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func Test_WorkerPool(t *testing.T) {
	dispatcher := NewDispatcher(100)
	workers := 3
	for i := 0; i < workers; i++ {
		execWorker := &ExecSthgWorker{}
		dispatcher.LaunchWorker(execWorker)
	}

	requests := 10

	wg := new(sync.WaitGroup)
	wg.Add(requests)
	for i := 0; i < requests; i++ {
		req := Request{
			Data: fmt.Sprintf("msg n°: %v", i),
			Handler: func(i interface{}) {
				fmt.Println(i)
				s, ok := i.(string)
				assert.True(t, ok)
				assert.Contains(t, s, "msg n°")
				wg.Done()
			},
		}
		dispatcher.MakeRequest(req)
	}
	dispatcher.Stop()
	wg.Wait()
}
