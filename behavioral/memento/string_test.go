package memento

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCareTaker_Add(t *testing.T) {
	originator := originator{}
	originator.state = State{Description: "idle"}

	carTaker := careTaker{}
	mem := originator.NewMemento()
	assert.Equal(t, "idle", mem.state.Description)

	currentLen := len(carTaker.mementoList)
	carTaker.Add(mem)
	assert.Equal(t, currentLen+1, len(carTaker.mementoList))
}

func TestCareTaker_Memento(t *testing.T) {
	originator := originator{}
	carTaker := careTaker{}

	originator.state = State{Description: "idle"}
	carTaker.Add(originator.NewMemento())

	mem, err := carTaker.Memento(0)
	assert.NoError(t, err)
	assert.Equal(t, "idle", mem.state.Description)

	_, err = carTaker.Memento(-1)
	assert.Error(t, err)
}

func TestOriginator_ExtractAndStoreState(t *testing.T) {
	originator := originator{state: State{Description: "idle"}}
	mem := originator.NewMemento()
	originator.state.Description = "working"
	originator.RestoreMemento(&mem)
	assert.Equal(t, "idle", mem.state.Description)
}
