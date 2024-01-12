package prototype

import "errors"

type ItemInfoGetter interface {
	GetInfo() string
}

type ShirtCloner interface {
	GetClone(s int) (ItemInfoGetter, error)
}

const (
	White = 1
	Black = 2
	Blue  = 3
)

var whitePrototype = &shirt{
	Price: 15.00,
	SKU:   "empty",
	Color: White,
}

var blackPrototype = &shirt{
	Price: 15.00,
	SKU:   "empty",
	Color: Black,
}

var bluePrototype = &shirt{
	Price: 15.00,
	SKU:   "empty",
	Color: Blue,
}

func GetShirtsCloner() ShirtCloner {
	return &ShirtCache{}
}

type ShirtCache struct {
}

func (shc *ShirtCache) GetClone(s int) (ItemInfoGetter, error) {
	switch s {
	case Black:
		newClone := *blackPrototype
		return &newClone, nil
	case White:
		newClone := *whitePrototype
		return &newClone, nil
	case Blue:
		newClone := *bluePrototype
		return &newClone, nil
	default:
		return nil, errors.New("invalid")
	}
}
