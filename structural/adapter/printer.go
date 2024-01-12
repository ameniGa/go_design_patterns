package adapter

import "fmt"

// ILegacyPrinter is the adapted interface
type ILegacyPrinter interface {
	Print(string) string
}

// IModernPrinter is the target
type IModernPrinter interface {
	PrintStored() string
}

/*****************************************************/

type LegacyPrinter struct {
}

func (lp *LegacyPrinter) Print(text string) string {
	return fmt.Sprintf("legacy printer prints: %v", text)
}

/*****************************************************/

// PrinterAdapter have to implement the target interface used by the client
type PrinterAdapter struct {
	OldPrinter ILegacyPrinter
	Msg        string
}

func (pa *PrinterAdapter) PrintStored() string {
	if pa.OldPrinter == nil {
		return pa.Msg
	}
	return pa.OldPrinter.Print(pa.Msg)
}

/*****************************************************/
