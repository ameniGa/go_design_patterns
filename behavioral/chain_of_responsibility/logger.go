package chain_of_responsibility

import (
	"fmt"
	"io"
	"strings"
)

type ChainLogger interface {
	Next(msg string)
	//SetNext(next ChainLogger)
}

type FirstLogger struct {
	nextChain ChainLogger
}

func (fl *FirstLogger) Next(msg string) {
	newMsg := fmt.Sprintf("First logger: %s", msg)
	fmt.Println(newMsg)
	if fl.nextChain != nil {
		fl.nextChain.Next(newMsg)
	}
}

type SecondLogger struct {
	nextChain ChainLogger
}

func (sl *SecondLogger) Next(msg string) {
	if strings.Contains(msg, "hello") {
		newMsg := fmt.Sprintf("Second logger: %s", msg)
		fmt.Println(newMsg)
		if sl.nextChain != nil {
			sl.nextChain.Next(newMsg)
		}
		return
	}
	fmt.Println("logging finished in Second Logger")
}

type WriterLogger struct {
	nextChain ChainLogger
	Writer    io.Writer
}

func (thl *WriterLogger) Next(msg string) {
	if thl.Writer != nil {
		_, _ = thl.Writer.Write([]byte(msg))
	}

	if thl.nextChain != nil {
		thl.nextChain.Next(msg)
	}
}
