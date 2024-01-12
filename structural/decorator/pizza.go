package decorator

import (
	"errors"
	"fmt"
)

// IPizza is the component
type IPizza interface {
	AddIngredient() (string, error)
}

type Pizza struct {
}

func (ia *Pizza) AddIngredient() (string, error) {
	return "Pizza with the following ingredients:", nil
}

// MeatDecorator is a concrete decorator
type MeatDecorator struct {
	Ingredient IPizza
}

func (md *MeatDecorator) AddIngredient() (string, error) {
	if md.Ingredient == nil {
		return "", errors.New("missing ingredient ")
	}
	s, err := md.Ingredient.AddIngredient()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s %s", s, "meat"), nil
}

// OnionDecorator is a concrete decorator
type OnionDecorator struct {
	Ingredient IPizza
}

func (od *OnionDecorator) AddIngredient() (string, error) {
	if od.Ingredient == nil {
		return "", errors.New("missing ingredient ")
	}
	s, err := od.Ingredient.AddIngredient()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s %s", s, "onion"), nil
}
