package bridge_test

import (
	"github.com/ameniGa/go_design_patterns/mocks"
	"github.com/ameniGa/go_design_patterns/structural/bridge"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrinterImpl1_Print(t *testing.T) {
	t.Run("implementation 1", func(t *testing.T) {
		impl1 := bridge.PrinterImpl1{}
		err := impl1.PrintMessage("hello")
		assert.NoError(t, err)
	})
}

func TestPrinterImpl2_Print(t *testing.T) {

	t.Run("nil writer", func(t *testing.T) {
		impl1 := bridge.PrinterImpl2{}
		err := impl1.PrintMessage("hello")
		assert.Error(t, err)
	})
	t.Run("valid writer", func(t *testing.T) {
		testWriter := mocks.WriterMock{}
		impl2 := bridge.PrinterImpl2{Writer: &testWriter}
		err := impl2.PrintMessage("hello")
		assert.NoError(t, err)
		assert.Contains(t, testWriter.Msg, "hello")
	})
}

func TestNormalPrinter_Print(t *testing.T) {
	t.Run("nil api", func(t *testing.T) {
		np := bridge.NormalPrinter{Msg: "hello"}
		err := np.Print()
		assert.Error(t, err)
	})
	t.Run("valid api : impl1", func(t *testing.T) {
		np := bridge.NormalPrinter{Msg: "hello", Printer: &bridge.PrinterImpl1{}}
		err := np.Print()
		assert.NoError(t, err)
	})
	t.Run("valid api: impl2", func(t *testing.T) {
		testWriter := mocks.WriterMock{}
		np := bridge.NormalPrinter{Msg: "hello", Printer: &bridge.PrinterImpl2{Writer: &testWriter}}
		err := np.Print()
		assert.NoError(t, err)
		assert.Contains(t, testWriter.Msg, "hello")
	})
}

func TestPacktPrinter_Print(t *testing.T) {
	t.Run("nil api", func(t *testing.T) {
		np := bridge.PacktPrinter{Msg: "hello"}
		err := np.Print()
		assert.Error(t, err)
	})
	t.Run("valid api : impl1", func(t *testing.T) {
		np := bridge.PacktPrinter{Msg: "hello", Printer: &bridge.PrinterImpl1{}}
		err := np.Print()
		assert.NoError(t, err)
	})
	t.Run("valid api: impl2", func(t *testing.T) {
		testWriter := mocks.WriterMock{}
		np := bridge.PacktPrinter{Msg: "hello", Printer: &bridge.PrinterImpl2{Writer: &testWriter}}
		err := np.Print()
		assert.NoError(t, err)
		assert.Contains(t, testWriter.Msg, "Message from Packt: hello")
	})
}
