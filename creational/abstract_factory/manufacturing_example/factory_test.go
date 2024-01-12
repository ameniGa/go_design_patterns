package manufacturing_example_test

import (
	. "github.com/ameniGa/go_design_patterns/creational/abstract_factory/manufacturing_example"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetManufacturingFactory(t *testing.T) {
	t.Run("modern factory", func(t *testing.T) {
		factory, err := GetManufacturingFactory(ModernManufacturing)
		assert.NoError(t, err)
		chair := factory.CreateChair()
		sofa := factory.CreateSofa()
		coffeeTable := factory.CreateCoffeeTable()
		_, ok := chair.(*ModernChair)
		assert.True(t, ok)
		_, ok = sofa.(*ModernSofa)
		assert.True(t, ok)
		_, ok = coffeeTable.(*ModernCoffeeTable)
		assert.True(t, ok)
	})
	t.Run("victorian factory", func(t *testing.T) {
		factory, err := GetManufacturingFactory(VictorianManufacturing)
		assert.NoError(t, err)
		chair := factory.CreateChair()
		sofa := factory.CreateSofa()
		coffeeTable := factory.CreateCoffeeTable()
		_, ok := chair.(*VictorianChair)
		assert.True(t, ok)
		_, ok = sofa.(*VictorianSofa)
		assert.True(t, ok)
		_, ok = coffeeTable.(*VictorianCoffeeTable)
		assert.True(t, ok)
	})
	t.Run("invalid factory", func(t *testing.T) {
		_, err := GetManufacturingFactory(10)
		assert.Error(t, err)
	})
}
