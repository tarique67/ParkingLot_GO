package parkinglot

import (
	"testing"
)

func TestAttendantCreated(t *testing.T) {
	attendant := NewAttendant()

	if attendant == nil {
		t.Errorf("Expected attendant to be created but was not.")
	}
}

func TestAttendantAbleToParkAfterAssigningParkingLot(t *testing.T) {
	attendant := NewAttendant()
	parkingLot,_ := NewParkingLot(1)
	attendant.assign(parkingLot)
	vehicle := NewVehicle("JH01BK0087", "White")

	token,_ := attendant.park(vehicle)

	if token == "" {
		t.Errorf("Expected a token but did not receive")
	}
}

func TestExceptionWhenParkingIfAllLotsFull(t *testing.T) {
	attendant := NewAttendant()
	parkingLot,_ := NewParkingLot(1)
	attendant.assign(parkingLot)
	vehicle := NewVehicle("JH01BK0087", "White")
	vehicle2 := NewVehicle("BR01KM0098", "Blue")
	attendant.park(vehicle)

	_,err := attendant.park(vehicle2)

	if err == nil {
		t.Errorf("Expected exception but did not receive")
	}
}

func TestExceptionWhenTryingToUnparkFromEmptySlot(t *testing.T) {
	attendant := NewAttendant()
	parkingLot,_ := NewParkingLot(1)
	attendant.assign(parkingLot)

	_,err := attendant.unpark("abc")

	if err == nil {
		t.Errorf("Expected exception while unparking from vacant lot but did not receive.")
	}
}

func TestCorrectVehicleReturnedWhenUnparking(t *testing.T) {
	attendant := NewAttendant()
	parkingLot,_ := NewParkingLot(1)
	attendant.assign(parkingLot)
	vehicle := NewVehicle("MP01BH0032", "Blue")
	token,_ := attendant.park(vehicle)

	returnedVehicle,_ := attendant.unpark(token)

	if  vehicle != returnedVehicle{
		t.Errorf("Expected correct vehicle returned when unparked but got another")
	}
}