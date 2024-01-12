package state

import "fmt"

type Gate interface {
	//Request()

	Enter()
	Pay()
	PayOK()
	PayFailed()
	ChangeState(state GateState)
}

type MyGate struct {
	state GateState
}

func (g *MyGate) Enter() {
	g.state.Enter()
}

func (g *MyGate) Pay() {
	g.state.Pay()
}

func (g *MyGate) PayOK() {
	g.state.PayOK()
}

func (g *MyGate) PayFailed() {
	g.state.PayFailed()
}

func (g *MyGate) ChangeState(state GateState) {
	g.state = state
}

type GateState interface {
	// Handle()

	Enter()
	Pay()
	PayOK()
	PayFailed()
}

type OpenGateState struct {
	gate Gate
}

func (ogs *OpenGateState) Enter() {
	fmt.Println("entering")
	ogs.gate.ChangeState(&ProcessingGateState{gate: ogs.gate})
}

func (ogs *OpenGateState) Pay() {
	fmt.Println("paying")
	ogs.gate.ChangeState(&ProcessingGateState{gate: ogs.gate})
}

func (ogs *OpenGateState) PayOK() {
	fmt.Println("let user in")
	ogs.gate.ChangeState(&ClosedGateState{gate: ogs.gate})

}

func (ogs *OpenGateState) PayFailed() {
	fmt.Println("pay failed")
	ogs.gate.ChangeState(&ClosedGateState{gate: ogs.gate})
}

/********************************************/

type ClosedGateState struct {
	gate Gate
}

func (cgs *ClosedGateState) Enter() {
	fmt.Println("entering")
	cgs.gate.ChangeState(&ProcessingGateState{gate: cgs.gate})
}

func (cgs *ClosedGateState) Pay() {
	fmt.Println("paying")
	cgs.gate.ChangeState(&ProcessingGateState{gate: cgs.gate})
}

func (cgs *ClosedGateState) PayOK() {
	fmt.Println("pay ok")
	cgs.gate.ChangeState(&OpenGateState{gate: cgs.gate})
}

func (cgs *ClosedGateState) PayFailed() {
	fmt.Println("pay failed")
}

/********************************************/

type ProcessingGateState struct {
	gate Gate
}

func (pgs *ProcessingGateState) Enter() {
	fmt.Println("entering")
}

func (pgs *ProcessingGateState) Pay() {
	fmt.Println("paying")
}

func (pgs *ProcessingGateState) PayOK() {
	fmt.Println("pay ok")
	pgs.gate.ChangeState(&OpenGateState{gate: pgs.gate})
}

func (pgs *ProcessingGateState) PayFailed() {
	fmt.Println("pay failed")
	pgs.gate.ChangeState(&ClosedGateState{gate: pgs.gate})
}
