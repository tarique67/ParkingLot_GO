package parkinglot

import "errors"

type ParkingLot struct {
	slots []*Slot
}

func NewParkingLot(size int) (*ParkingLot, error) {
	if(size<=0){
		return nil, errors.New("cannot create Parking lot of size less than 1")
	}
	parkingLot := &ParkingLot{}
	parkingLot.slots = make([]*Slot, size)
	for i :=0; i<size; i++ {
		parkingLot.slots[i],_ = NewSlot(i)
	}
	return parkingLot, nil
}

func (p *ParkingLot) park(vehicle *Vehicle) (string, error) {
	for i :=0; i<len(p.slots); i++ {
		if(!p.slots[i].isVehicleParked()){
			token, err := p.slots[i].park(vehicle)
			if err != nil {
				return "", err
			}
			return token,nil
		}
	}
	return "", errors.New("parking lot full")
}

func (p *ParkingLot) unpark(token string) (*Vehicle, error) {
	for i :=0; i<len(p.slots); i++ {
		if(p.slots[i].isVehicleParked()){
			returnedVehicle, err := p.slots[i].unpark(token)
			if err != nil {
				return nil, err
			}
			return returnedVehicle,nil
		}
	}
	return nil, errors.New("token did not match.")
}

func (p *ParkingLot) isFull() bool {
	for i:=0; i<len(p.slots); i++ {
		if(!p.slots[i].isVehicleParked()){
			return false
		}
	}
	return true
}