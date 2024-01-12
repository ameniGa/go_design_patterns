package manufacturing_example

type ModernFactory struct {
}

func (mf *ModernFactory) CreateChair() Chair {
	return &ModernChair{}
}

func (mf *ModernFactory) CreateCoffeeTable() CoffeeTable {
	return &ModernCoffeeTable{}
}

func (mf *ModernFactory) CreateSofa() Sofa {
	return &ModernSofa{}
}
