package decorator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPizzaDecorator_AddIngredient(t *testing.T) {
	pizza := &Pizza{}
	pizzaRes, _ := pizza.AddIngredient()
	expected := "Pizza with the following ingredients:"
	assert.Contains(t, pizzaRes, expected)
}

func TestOnionDecorator_AddIngredient(t *testing.T) {
	t.Run("nil ingredient", func(t *testing.T) {
		onion := OnionDecorator{}
		_, err := onion.AddIngredient()
		assert.Error(t, err)
	})
	t.Run("valid ingredient", func(t *testing.T) {
		onion := OnionDecorator{&Pizza{}}
		res, err := onion.AddIngredient()
		assert.NoError(t, err)
		assert.Contains(t, res, "onion")
	})
}

func TestMeat_AddIngredient(t *testing.T) {
	t.Run("nil ingredient", func(t *testing.T) {
		meat := MeatDecorator{}
		_, err := meat.AddIngredient()
		assert.Error(t, err)
	})
	t.Run("valid ingredient", func(t *testing.T) {
		meat := MeatDecorator{&Pizza{}}
		res, err := meat.AddIngredient()
		assert.NoError(t, err)
		assert.Contains(t, res, "meat")
	})
}

func TestPizzaDecorator_FullStack(t *testing.T) {
	pizza := &OnionDecorator{&MeatDecorator{&Pizza{}}}
	pizzaRes, _ := pizza.AddIngredient()
	assert.Contains(t, pizzaRes, "meat")
	assert.Contains(t, pizzaRes, "onion")
	t.Log(pizzaRes)
}
