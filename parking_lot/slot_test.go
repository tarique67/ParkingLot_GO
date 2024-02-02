package parkinglot

import (
	"os/exec"
	"testing"
)

func TestNewSlotThrowsError(t *testing.T) {
	_, error := NewSlot(-1)

	if error == nil{
		t.Errorf("Expected error to be thrown with negative slot no. but was not thrown")
	}
}

func TestNewSlotCreated(t *testing.T) {
	slot,_ := NewSlot(0)

	if (slot == nil) {
		t.Errorf("Expected Slot to be created but was not created")
	}
}

func TestSlotFree(t *testing.T) {
	slot, _ := NewSlot(1)

	if(slot.isVehicleParked()) {
		t.Errorf("Expected slot to be free but got vehicle is parked.")
	}
}

func TestPark(t *testing.T) {
	vehicle := NewVehicle("Jh01na7612", "Black")
	slot,_ := NewSlot(1)

	_,error := slot.park(vehicle)

	if(error != nil){
		t.Errorf("Expected vehicle to be parked but an error thrown")
	}
}

func TestUnableToParkOnFullSlot(t *testing.T) {
	vehicle := NewVehicle("Jh01na7612", "Black")
	vehicle2 := NewVehicle("MH05DS7612", "Black")
	slot,_ := NewSlot(1)

	_,err1 := slot.park(vehicle)
	_,err2 := slot.park(vehicle2)

	if(err1 == nil && err2 == nil){
		t.Errorf("Expected exception to be thrown when parking on full slot but was not thrown.")
	}
}

func TestExceptionThrownIfUnparkWithInvalidToken(t *testing.T) {
	vehicle := NewVehicle("Jh01na7612", "Black")
	slot,_ := NewSlot(1)

	slot.park(vehicle)
	uuid,_ := exec.Command("uuidgen").Output()
	invalidToken := string(uuid)
	_,err := slot.unpark(invalidToken)
	if(err == nil){
		t.Errorf("Expected exception when unparking with invalid token but got no exception")
	}
	
}

func TestCorrectVehicleReturnedOnUnparkWithCorrectToken(t *testing.T) {
	vehicle := NewVehicle("Jh01na7612", "Black")
	slot,_ := NewSlot(1)

	token,err1 := slot.park(vehicle)
	if(err1 != nil ){
		t.Errorf("Expected vehicle to parked but was not parked.")
	}

	vehicleReturned,_ := slot.unpark(token)
	if(vehicleReturned != vehicle){
		t.Errorf("Expected to be able to unpark but could not unpark.")
	}
	
}
