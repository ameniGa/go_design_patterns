package prototype

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetShirtsCloner(t *testing.T) {
	cloner := GetShirtsCloner()
	assert.NotNil(t, cloner)
	clone1, err := cloner.GetClone(White)
	assert.NoError(t, err)
	assert.NotEqual(t, &whitePrototype, &clone1)
	shirt1, ok := clone1.(*shirt)
	assert.True(t, ok)
	shirt1.SKU = "abbcc"

	clone2, err := cloner.GetClone(White)
	shirt2, ok := clone2.(*shirt)
	assert.NotEqual(t, shirt2, shirt1)
	assert.NotEqual(t, shirt1.SKU, shirt2.SKU)

	t.Logf("shirt1: %s", shirt1.GetInfo())
	t.Logf("shirt2: %s", shirt2.GetInfo())

}
