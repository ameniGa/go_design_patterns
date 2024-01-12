package future

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestStringOrError_Execute(t *testing.T) {
	future := MaybeString{}
	t.Run("success", func(t *testing.T) {
		wg := new(sync.WaitGroup)
		wg.Add(1)
		future.Success(func(s string) {
			assert.Equal(t, "hello world", s)
			wg.Done()
		}).Fail(func(err error) {
			assert.NoError(t, err)
			wg.Done()
		})
		future.Execute(func() (string, error) {
			return "hello world", nil
		})
		wg.Wait()
	})
	t.Run("failure", func(t *testing.T) {
		wg := new(sync.WaitGroup)
		wg.Add(1)
		future.Success(func(s string) {
			assert.Equal(t, "", s)
			wg.Done()
		}).Fail(func(err error) {
			assert.Error(t, err)
			wg.Done()
		})
		future.Execute(func() (string, error) {
			return "", fmt.Errorf("error")
		})
		wg.Wait()
	})
}
