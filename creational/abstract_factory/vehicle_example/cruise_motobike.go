package vehicle_example

type CruiseMotorbike struct {
}

func (cm *CruiseMotorbike) GetMotorbikeType() int {
	return CruiseMotorbikeType
}

func (cm *CruiseMotorbike) NumWheels() int {
	return 2
}

func (cm *CruiseMotorbike) NumSeats() int {
	return 2
}
