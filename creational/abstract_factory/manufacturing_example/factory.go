package manufacturing_example

import "errors"

type Chair interface {
	SetColor(color string)
}

type Sofa interface {
	SetSize(size int)
}

type CoffeeTable interface {
	SetSize(size int)
	SetShape(shape string)
}

type ManufacturingFactory interface {
	CreateChair() Chair
	CreateCoffeeTable() CoffeeTable
	CreateSofa() Sofa
}

const (
	ModernManufacturing    = 1
	VictorianManufacturing = 2
)

func GetManufacturingFactory(factoryType int) (ManufacturingFactory, error) {
	switch factoryType {
	case ModernManufacturing:
		return &ModernFactory{}, nil
	case VictorianManufacturing:
		return &VictorianFactory{}, nil
	default:
		return nil, errors.New("unsupported")
	}
}
