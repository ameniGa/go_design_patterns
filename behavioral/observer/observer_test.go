package observer

import "testing"

func Test_Observer(t *testing.T) {
	subscriber := subscriber{id: 10}
	var publisher publisher
	publisher.subscribe(subscriber)
	publisher.notify(event{data: 100})
	//time.Sleep(3*time.Second)
	publisher.unsubscribe(subscriber)
}
