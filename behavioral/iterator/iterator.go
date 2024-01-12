package iterator

// Inventory is the iterable
type Inventory interface {
	GetIterator() InventoryIterator
}

type HomeInventory struct {
	// for the sake of example we used a slice, but it can be any complex type like a tree ...
	tools []string
}

func (ci *HomeInventory) GetIterator() InventoryIterator {
	return &HomeInventoryIterator{Tools: ci.tools}
}

type InventoryIterator interface {
	HasNext() bool
	Next()
	Current() any
}

type HomeInventoryIterator struct {
	Tools []string
	index int
}

func (ci *HomeInventoryIterator) HasNext() bool {
	if ci.index < len(ci.Tools) {
		return true
	}
	return false
}

func (ci *HomeInventoryIterator) Next() {
	if ci.HasNext() {
		ci.index++
	}
}

func (ci *HomeInventoryIterator) Current() any {
	return ci.Tools[ci.index]
}
