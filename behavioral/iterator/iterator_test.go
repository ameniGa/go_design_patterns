package iterator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConcreteIterable_GetIterator(t *testing.T) {
	tools := []string{"tool1", "tool2", "tool3"}
	iterable := HomeInventory{tools: tools}
	it := iterable.GetIterator()
	i := 0
	for it.HasNext() {
		current := it.Current()
		it.Next()
		assert.Equal(t, tools[i], current)
		i++
	}
}
