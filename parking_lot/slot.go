package parkinglot

import (
	"errors"
	"os/exec"
)

type Slot struct{
	SlotNo int
	Vehicle *Vehicle
	Token string
}

func NewSlot(slotNo int) (*Slot, error){
	if slotNo<0 {
		return nil, errors.New("Slot no. cannot be negative")
	}
	return &Slot{SlotNo: slotNo}, nil
}

func (s *Slot) isVehicleParked() bool {
	return s.Vehicle != nil
}

func (s *Slot) park(vehicle *Vehicle) (string, error) {
	if s.isVehicleParked() {
		return "", errors.New("Slot is already occupied")
	}
	s.Vehicle = vehicle;
	uuid,_ := exec.Command("uuidgen").Output()
	s.Token = string(uuid)
	return s.Token, nil 
}

func (s *Slot) unpark(token string) (*Vehicle, error) {
	if s.Token != token {
		return nil, errors.New("Token does not match")
	}
	vehicle := s.Vehicle;
	s.Vehicle = nil
	s.Token = ""
	return vehicle, nil 
}