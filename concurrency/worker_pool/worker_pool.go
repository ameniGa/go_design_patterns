package worker_pool

import "time"

type RequestHandler func(interface{})

type Request struct {
	Data    interface{}
	Handler RequestHandler
}

type WorkerLauncher interface {
	launchWorker(in chan Request)
}

type Dispatcher interface {
	LaunchWorker(w WorkerLauncher)
	MakeRequest(request Request)
	Stop()
}

type dispatcher struct {
	inCh chan Request
}

func NewDispatcher(b int) Dispatcher {
	return &dispatcher{inCh: make(chan Request, b)}
}

func (d *dispatcher) LaunchWorker(w WorkerLauncher) {
	w.launchWorker(d.inCh)
}

func (d *dispatcher) MakeRequest(request Request) {
	select {
	case d.inCh <- request:
	case <-time.After(time.Second * 5):
		return
	}
}

func (d *dispatcher) Stop() {
	close(d.inCh)
}

type ExecSthgWorker struct {
}

func (wl *ExecSthgWorker) launchWorker(in chan Request) {
	go func() {
		for req := range in {
			req.Handler(req.Data)
		}
	}()
}
