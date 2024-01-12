package barrier

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type BarrierRes struct {
	Err  error
	Resp string
}

func Barrier(endpoints ...string) {
	reqCount := len(endpoints)
	in := make(chan BarrierRes, reqCount)
	defer close(in)

	responses := make([]BarrierRes, reqCount)

	for _, endpoint := range endpoints {
		go makeRequest(in, endpoint)
	}

	hasErr := false
	for i := 0; i < reqCount; i++ {
		resp := <-in
		if resp.Err != nil {
			fmt.Println("ERROR: ", resp.Err)
			hasErr = true
		}
		responses[i] = resp
	}
	if !hasErr {
		for _, resp := range responses {
			fmt.Println(resp.Resp)
		}
	}

}

var timeOutMilliseconds = 5000

func makeRequest(out chan<- BarrierRes, url string) {
	res := BarrierRes{}
	client := http.Client{Timeout: time.Duration(timeOutMilliseconds) * time.Millisecond}

	resp, err := client.Get(url)
	if err != nil {
		res.Err = err
		out <- res
		return
	}
	by, err := io.ReadAll(resp.Body)
	if err != nil {
		res.Err = err
		out <- res
		return
	}
	res.Resp = string(by)
	out <- res
}

func captureBarrierOutput(endpoints ...string) string {
	reader, writer, _ := os.Pipe()
	os.Stdout = writer

	out := make(chan string)

	go func() {
		var buf bytes.Buffer
		_, _ = io.Copy(&buf, reader)
		out <- buf.String()
	}()

	Barrier(endpoints...)
	_ = writer.Close()
	temp := <-out
	return temp
}
