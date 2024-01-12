package builder_test

import (
	"github.com/ameniGa/go_design_patterns/creational/builder"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestManufacturingBuilder(t *testing.T) {
	mf := builder.ManufacturingBuilder{}
	carBuilder := &builder.CarBuilder{}
	mf.SetBuilder(carBuilder)
	mf.Construct()
	car := carBuilder.Build()
	assert.Equal(t, 4, car.Wheels)
	assert.Equal(t, 5, car.Seats)
	assert.Equal(t, "Car", car.Structure)

	bikeBuilder := &builder.BikeBuilder{}
	mf.SetBuilder(bikeBuilder)
	mf.Construct()
	bike := bikeBuilder.Build()
	assert.Equal(t, 2, bike.Wheels)
	assert.Equal(t, 1, bike.Seats)
	assert.Equal(t, "Bike", bike.Structure)
}
