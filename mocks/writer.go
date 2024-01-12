package mocks

import "errors"

type WriterMock struct {
	Msg string
}

func (wm *WriterMock) Write(p []byte) (int, error) {
	n := len(p)
	if n > 0 {
		wm.Msg = string(p)
		return n, nil
	}
	return n, errors.New("received empty content")
}
