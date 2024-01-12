package builder

type VehicleProduct struct {
	Seats     int
	Wheels    int
	Structure string
}

// BuildProcess defindes the common steps used to create a vehicle product
type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	Build() VehicleProduct
}

type ManufacturingBuilder struct {
	builder BuildProcess
}

func (mf *ManufacturingBuilder) SetBuilder(process BuildProcess) {
	mf.builder = process
}

func (mf *ManufacturingBuilder) Construct() {
	mf.builder.SetWheels().SetSeats().SetStructure()
}
