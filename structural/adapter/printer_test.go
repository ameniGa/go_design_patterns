package adapter_test

import (
	. "github.com/ameniGa/go_design_patterns/structural/adapter"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLegacyPrinter_Print(t *testing.T) {
	t.Run("valid legacy interface", func(t *testing.T) {
		msg := "hello world!"
		adapter := PrinterAdapter{OldPrinter: &LegacyPrinter{}, Msg: msg}
		s := adapter.PrintStored()
		assert.Contains(t, s, "legacy")
	})
	t.Run("nil legacy interface", func(t *testing.T) {
		msg := "hello world!"
		adapter := PrinterAdapter{OldPrinter: nil, Msg: msg}
		s := adapter.PrintStored()
		assert.Equal(t, s, "hello world!")
	})

}
