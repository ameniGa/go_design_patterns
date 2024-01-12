package manufacturing_example

type VictorianFactory struct {
}

func (vf *VictorianFactory) CreateChair() Chair {
	return &VictorianChair{}
}

func (vf *VictorianFactory) CreateCoffeeTable() CoffeeTable {
	return &VictorianCoffeeTable{}
}

func (vf *VictorianFactory) CreateSofa() Sofa {
	return &VictorianSofa{}
}
