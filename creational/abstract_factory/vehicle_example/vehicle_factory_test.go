package vehicle_example_test

import (
	. "github.com/ameniGa/go_design_patterns/creational/abstract_factory/vehicle_example"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuildFactory(t *testing.T) {
	t.Run("create car factory", func(t *testing.T) {
		vehicleFactory, err := BuildFactory(CarFactoryType)
		assert.NoError(t, err)
		_, ok := vehicleFactory.(*CarFactory)
		assert.True(t, ok)
		vehicle, err := vehicleFactory.NewVehicle(LuxuryCarType)
		assert.NoError(t, err)
		assert.Equal(t, 5, vehicle.NumSeats())
		assert.Equal(t, 4, vehicle.NumWheels())
		luxCar, ok := vehicle.(*LuxuryCar)
		assert.True(t, ok)
		assert.Equal(t, 4, luxCar.NumDoors())
	})

	t.Run("create motorbike factory", func(t *testing.T) {
		vehicleFactory, err := BuildFactory(MotorbikeFactoryType)
		assert.NoError(t, err)
		_, ok := vehicleFactory.(*MotorbikeFactory)
		assert.True(t, ok)
		vehicle, err := vehicleFactory.NewVehicle(CruiseMotorbikeType)
		assert.NoError(t, err)
		assert.Equal(t, 2, vehicle.NumSeats())
		assert.Equal(t, 2, vehicle.NumWheels())
		luxCar, ok := vehicle.(*CruiseMotorbike)
		assert.True(t, ok)
		assert.Equal(t, CruiseMotorbikeType, luxCar.GetMotorbikeType())
	})

	t.Run("invalid factory", func(t *testing.T) {
		_, err := BuildFactory(100)
		assert.Error(t, err)
	})
}
