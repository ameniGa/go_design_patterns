package publish_subcriber

import "fmt"

type Subscriber interface {
	Notify(interface{}) error
}

type WriterSubscriber struct {
}

func (ws *WriterSubscriber) Notify(data interface{}) error {
	msg, ok := data.(string)
	if !ok {
		return fmt.Errorf("invalid")
	}
	fmt.Println(msg)
	return nil
}

func (ws *WriterSubscriber) Close() {

}

type Publisher interface {
	Start()
	AddSubscriberCh() chan<- Subscriber
	RemoveSubscriberCh() chan<- Subscriber
	PublishingCh() chan<- interface{}
	Stop()
}

type publisher struct {
	subscribers []Subscriber
	addSubCh    chan Subscriber
	removeSubCh chan Subscriber
	in          chan interface{}
	stop        chan struct{}
}

func NewPublisher(b int) Publisher {
	return &publisher{
		subscribers: make([]Subscriber, 0),
		addSubCh:    make(chan Subscriber, b),
		removeSubCh: make(chan Subscriber, b),
		in:          make(chan interface{}, b),
		stop:        make(chan struct{}),
	}
}

func (p *publisher) Start() {
	for {
		select {
		case msg := <-p.in:
			for _, sub := range p.subscribers {
				sub.Notify(msg)
			}
		case sub := <-p.addSubCh:
			p.subscribers = append(p.subscribers, sub)
		case sub := <-p.removeSubCh:
			for i, candidate := range p.subscribers {
				if candidate == sub {
					p.subscribers = append(p.subscribers[:i], p.subscribers[i+1:]...)
					break
				}
			}
		case <-p.stop:
			close(p.addSubCh)
			close(p.removeSubCh)
			close(p.in)
			return
		}
	}
}

func (p *publisher) AddSubscriberCh() chan<- Subscriber {
	return p.addSubCh
}

func (p *publisher) RemoveSubscriberCh() chan<- Subscriber {
	return p.removeSubCh
}

func (p *publisher) PublishingCh() chan<- interface{} {
	return p.in
}

func (p *publisher) Stop() {
	close(p.stop)
}
