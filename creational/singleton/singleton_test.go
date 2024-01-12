package singleton_test

import (
	"github.com/ameniGa/go_design_patterns/creational/singleton"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetInstance(t *testing.T) {
	t.Run("new instance", func(t *testing.T) {
		instance := singleton.GetInstance()
		assert.NotNil(t, instance)
		count := instance.AddOne()
		assert.Equal(t, 1, count)

		newInstance := singleton.GetInstance()
		assert.Equal(t, instance, newInstance)
		count2 := instance.AddOne()
		assert.Equal(t, 2, count2)
	})
}
