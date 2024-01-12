package barrier

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBarrier(t *testing.T) {
	t.Run("2 success", func(t *testing.T) {
		endpoints := []string{"http://httpbin.org/headers", "http://httpbin.org/user-agent"}
		res := captureBarrierOutput(endpoints...)
		assert.Contains(t, res, "Accept-Encoding", "User-Agent")
	})

	t.Run("1 failure", func(t *testing.T) {
		endpoints := []string{"http://malformed-url", "http://httpbin.org/User-Agent"}
		res := captureBarrierOutput(endpoints...)
		assert.Contains(t, res, "ERROR")
	})

	t.Run("very short timeout", func(t *testing.T) {
		endpoints := []string{"http://httpbin.org/headers", "http://httpbin.org/user-agent"}
		timeOutMilliseconds = 1
		res := captureBarrierOutput(endpoints...)
		assert.Contains(t, res, "Timeout")
	})
}
