package parkinglot

import "errors"

type Attendant struct {
	parkingLots []*ParkingLot
}

func NewAttendant() *Attendant {
	attendant := &Attendant{}
	attendant.parkingLots = []*ParkingLot{}
	return attendant
}

func (attendant *Attendant) assign(parkingLot *ParkingLot) {
	attendant.parkingLots = append(attendant.parkingLots, parkingLot)
}

func (attendant *Attendant) park(vehicle *Vehicle) (string, error) {
	for i := 0; i<len(attendant.parkingLots); i++ {
		if(!attendant.parkingLots[i].isFull()) {
			return attendant.parkingLots[i].park(vehicle)
		}
	}
	return "", errors.New("all parking lots full")
}

func (attendant *Attendant) unpark(token string) (*Vehicle, error) {
	for i := 0; i<len(attendant.parkingLots); i++ {
		vehicle,err := attendant.parkingLots[i].unpark(token)
		if(err == nil){
			return vehicle,nil
		}
	}
	return nil,errors.New("token did not match")
}