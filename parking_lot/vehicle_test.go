package parkinglot

import (
	"testing"
)

func TestNewVehicleCreated(t *testing.T) {
	vehicle := NewVehicle("Jh01Bg1234", "White")

	if(vehicle == nil){
		t.Errorf("Expected vehicle no. is Jh01Bg1234, but received %s", vehicle.RegistrationNo)
	}
}

func TestColorAvailable(t *testing.T) {
	vehicle := NewVehicle("JH01888", "White")

	if !vehicle.ColorAvailable("White") {
		t.Errorf("Expected color 'white' to be available, but it is unavailable")
	}
}

func TestColorUnavailable(t *testing.T) {
	vehicle := NewVehicle("JH01888", "White")
	
	if vehicle.ColorAvailable("Black") {
		t.Errorf("Expected color 'black' to be unavailable, but it is available")
	}
}