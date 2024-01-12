package vehicle_example

type SportMotorbike struct {
}

func (sm *SportMotorbike) GetMotorbikeType() int {
	return SportMotorbikeType
}

func (sm *SportMotorbike) NumWheels() int {
	return 2
}

func (sm *SportMotorbike) NumSeats() int {
	return 1
}
