package mediator

import "testing"

func TestMediator(t *testing.T) {
	mediator := StationManager{isPlatformFree: true, trainQueue: make([]train, 0)}
	train1 := PassengerTrain{mediator: &mediator}
	train2 := FreightTrain{mediator: &mediator}
	train1.Arrive()
	train2.Arrive()
	train1.Depart()

}
