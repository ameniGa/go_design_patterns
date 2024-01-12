package memento

import "fmt"

type State struct {
	Description string
}

type Memento struct {
	state State
}

func (m *Memento) GetSavedState() State {
	return m.state
}

// originator creates the Memento and store the current active state
type originator struct {
	state State
}

func (o *originator) NewMemento() Memento {
	return Memento{state: o.state}
}

func (o *originator) RestoreMemento(m *Memento) {
	o.state = m.GetSavedState()
}

func (o *originator) GetState() State {
	return o.state
}

func (o *originator) SetState(state State) {
	o.state = state
}

type careTaker struct {
	mementoList []Memento
}

func (c *careTaker) Add(m Memento) {
	c.mementoList = append(c.mementoList, m)
}

func (c *careTaker) Memento(i int) (Memento, error) {
	if len(c.mementoList) < i || i < 0 {
		return Memento{}, fmt.Errorf("invalid index")
	}
	return c.mementoList[i], nil
}
