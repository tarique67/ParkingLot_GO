package parkinglot

import (
	"testing"
)

func TestParkingLotCreated(t *testing.T) {
	parkingLot,_ := NewParkingLot(1)

	if(parkingLot == nil){
		t.Errorf("Expected parking lot to be created but was not created.")
	}
}

func TestExceptionThrownIfParkingLotCreatedWith0Size(t *testing.T) {
	_,err := NewParkingLot(0)

	if(err == nil){
		t.Errorf("Expected exception when creating a Parking lot with negative size but was not thrown.")
	}
}

func TestAbleToParkOnEmptyLot(t *testing.T) {
	parkinglot,_ := NewParkingLot(1);
	vehicle := NewVehicle("JH01HN0554", "White")

	token,_ := parkinglot.park(vehicle)

	if(token == "" ){
		t.Errorf("Expected vehicle to be parked, but was not parked.")
	}
}

func TestExceptionThrownIfParkingOnOccupiedLot(t *testing.T) {
	parkinglot,_ := NewParkingLot(1);
	vehicle := NewVehicle("JH01HN0554", "White")
	vehicle2 := NewVehicle("KA04GG0554", "White")
	parkinglot.park(vehicle)

	_,err := parkinglot.park(vehicle2)

	if(err == nil ){
		t.Errorf("Expected exception on parking in a full lot, but not thrown")
	}
}

func TestExceptionThrownOnTryingToUnparkWithInvalidToken(t *testing.T) {
	parkinglot,_ := NewParkingLot(1);
	vehicle := NewVehicle("JH01HN0554", "White")
	parkinglot.park(vehicle)

	
	_,err := parkinglot.unpark("abc")
	if(err == nil ){
		t.Errorf("Expected exception but was not thrown")
	}
}

func TestCorrectVehicleReturnedOnUnparkingWithValidToken(t *testing.T) {
	parkinglot,_ := NewParkingLot(1);
	vehicle := NewVehicle("JH01HN0554", "White")
	token,_ := parkinglot.park(vehicle)

	
	returnedVehicle,_ := parkinglot.unpark(token)
	
	if(returnedVehicle != vehicle ){
		t.Errorf("Expected correct vehicle but got another vehicle")
	}
}

func TestTrueForLotFull(t *testing.T) {
	parkinglot,_ := NewParkingLot(1);
	vehicle := NewVehicle("JH01HN0554", "White")
	parkinglot.park(vehicle)
	
	if !parkinglot.isFull() {
		t.Errorf("Expected parking lot to be full but was not full")
	}
}

func TestFalseForLotFull(t *testing.T) {
	parkinglot,_ := NewParkingLot(2);
	vehicle := NewVehicle("JH01HN0554", "White")
	parkinglot.park(vehicle)
	
	if parkinglot.isFull() {
		t.Errorf("Expected parking lot not to be full but was full")
	}
}