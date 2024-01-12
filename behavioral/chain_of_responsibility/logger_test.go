package chain_of_responsibility

import (
	"github.com/ameniGa/go_design_patterns/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWriterLogger_Next(t *testing.T) {
	t.Run("all chain ", func(t *testing.T) {
		writer := &mocks.WriterMock{}
		writerLogger := WriterLogger{Writer: writer}
		secondLogger := SecondLogger{nextChain: &writerLogger}
		firstLogger := FirstLogger{nextChain: &secondLogger}
		firstLogger.Next("hello world")
		assert.Equal(t, "Second logger: First logger: hello world", writer.Msg)
	})

	t.Run("exec stops at second chain ", func(t *testing.T) {
		writer := &mocks.WriterMock{}
		writerLogger := WriterLogger{Writer: writer}
		secondLogger := SecondLogger{nextChain: &writerLogger}
		firstLogger := FirstLogger{nextChain: &secondLogger}
		firstLogger.Next("learn design patterns")
		assert.Equal(t, "", writer.Msg)
	})
}
