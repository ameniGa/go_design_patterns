package vehicle_example

import "errors"

type Vehicle interface {
	NumWheels() int
	NumSeats() int
}

type VehicleFactory interface {
	NewVehicle(v int) (Vehicle, error)
}

const (
	CarFactoryType       = 1
	MotorbikeFactoryType = 2
)

// BuildFactory is the abstract factory
func BuildFactory(f int) (VehicleFactory, error) {
	switch f {
	case CarFactoryType:
		return &CarFactory{}, nil
	case MotorbikeFactoryType:
		return &MotorbikeFactory{}, nil
	default:
		return nil, errors.New("unsupported factory type")
	}
}

const (
	LuxuryCarType = 1
	FamilyCarType = 2
)

type CarFactory struct {
}

func (cf *CarFactory) NewVehicle(v int) (Vehicle, error) {
	switch v {
	case LuxuryCarType:
		return new(LuxuryCar), nil
	case FamilyCarType:
		return new(FamilyCar), nil
	default:
		return nil, errors.New("unsupported type")
	}
}

const (
	SportMotorbikeType  = 1
	CruiseMotorbikeType = 2
)

type MotorbikeFactory struct {
}

func (mf *MotorbikeFactory) NewVehicle(v int) (Vehicle, error) {
	switch v {
	case SportMotorbikeType:
		return new(SportMotorbike), nil
	case CruiseMotorbikeType:
		return new(CruiseMotorbike), nil
	default:
		return nil, errors.New("unsupported type")
	}
}
