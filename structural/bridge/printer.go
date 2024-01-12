package bridge

import (
	"errors"
	"fmt"
	"io"
)

// PrinterAPI is the implementor
type PrinterAPI interface {
	PrintMessage(string) error
}

type PrinterImpl1 struct {
}

func (pi PrinterImpl1) PrintMessage(text string) error {
	fmt.Println(text)
	return nil
}

type PrinterImpl2 struct {
	Writer io.Writer
}

func (pi PrinterImpl2) PrintMessage(text string) error {
	if pi.Writer == nil {
		return errors.New("missing writer")
	}
	_, err := pi.Writer.Write([]byte(text))
	return err
}

/*************************************************************************/

// PrinterAbstraction is the abstraction
type PrinterAbstraction interface {
	Print() error
}

type NormalPrinter struct {
	Msg     string
	Printer PrinterAPI
}

func (pa *NormalPrinter) Print() error {
	if pa.Printer == nil {
		return errors.New("missing printer api")
	}
	return pa.Printer.PrintMessage(pa.Msg)
}

type PacktPrinter struct {
	Msg     string
	Printer PrinterAPI
}

func (pp *PacktPrinter) Print() error {
	if pp.Printer == nil {
		return errors.New("missing printer api")
	}
	return pp.Printer.PrintMessage(fmt.Sprintf("Message from Packt: %v", pp.Msg))
}
