package builder

type CarBuilder struct {
	v VehicleProduct
}

func (cp *CarBuilder) SetWheels() BuildProcess {
	cp.v.Wheels = 4
	return cp
}

func (cp *CarBuilder) SetSeats() BuildProcess {
	cp.v.Seats = 5
	return cp
}
func (cp *CarBuilder) SetStructure() BuildProcess {
	cp.v.Structure = "Car"
	return cp
}

func (cp *CarBuilder) Build() VehicleProduct {
	return cp.v
}
