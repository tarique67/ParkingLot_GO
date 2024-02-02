package parkinglot

type Vehicle struct{
	RegistrationNo string
	Color string
}

func NewVehicle(registrationNo string, color string) *Vehicle {
	return &Vehicle{RegistrationNo: registrationNo, Color: color}
}

func (v *Vehicle) ColorAvailable(color string) bool {
	return v.Color == color
}
