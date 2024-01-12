package mediator

import "fmt"

type train interface {
	Arrive()
	Depart()
	PermitArrival()
}

type Mediator interface {
	canArrive(train train) bool
	notifyAboutDeparture()
}

// StationManager is the mediator
type StationManager struct {
	isPlatformFree bool
	trainQueue     []train
}

func (sm *StationManager) canArrive(train train) bool {
	if sm.isPlatformFree {
		sm.isPlatformFree = false
		return true
	}
	sm.trainQueue = append(sm.trainQueue, train)
	return false
}

func (sm *StationManager) notifyAboutDeparture() {
	if !sm.isPlatformFree {
		sm.isPlatformFree = true
	}
	if len(sm.trainQueue) > 0 {
		firstTrainInQueue := sm.trainQueue[0]
		sm.trainQueue = sm.trainQueue[1:]
		firstTrainInQueue.PermitArrival()
	}
}

type PassengerTrain struct {
	mediator Mediator
}

func (g *PassengerTrain) Arrive() {
	if !g.mediator.canArrive(g) {
		fmt.Println("PassengerTrain: Arrival blocked, waiting")
		return
	}
	fmt.Println("PassengerTrain: Arrived")
}

func (g *PassengerTrain) Depart() {
	fmt.Println("PassengerTrain: Leaving")
	g.mediator.notifyAboutDeparture()
}

func (g *PassengerTrain) PermitArrival() {
	fmt.Println("PassengerTrain: Arrival permitted, arriving")
	g.Arrive()
}

type FreightTrain struct {
	mediator Mediator
}

func (g *FreightTrain) Arrive() {
	if !g.mediator.canArrive(g) {
		fmt.Println("FreightTrain: Arrival blocked, waiting")
		return
	}
	fmt.Println("FreightTrain: Arrived")
}

func (g *FreightTrain) Depart() {
	fmt.Println("FreightTrain: Leaving")
	g.mediator.notifyAboutDeparture()
}

func (g *FreightTrain) PermitArrival() {
	fmt.Println("FreightTrain: Arrival permitted")
	g.Arrive()
}
