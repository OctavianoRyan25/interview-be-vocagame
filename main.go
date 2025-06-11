package main

import "fmt"

type ParkingIOT struct {
	vehicles []*Vehicle
}

type Vehicle struct {
	vehicleID string
}

func create_parking_iot(size int) *ParkingIOT {
	return &ParkingIOT{
		vehicles: make([]*Vehicle, size),
	}
}

func (p *ParkingIOT) parkVehicle(vehicleID string) string {
	for i := 0; i < len(p.vehicles); i++ {
		if p.vehicles[i] == nil {
			p.vehicles[i] = &Vehicle{vehicleID: vehicleID}
			return "Allocated slot number: " + fmt.Sprint(i+1)
		}
	}
	return "Sorry, parking iot is full"
}

func (p *ParkingIOT) removeVehicle(vehicleID string, hours int) string {
	for i := 0; i < len(p.vehicles); i++ {
		if p.vehicles[i] != nil && p.vehicles[i].vehicleID == vehicleID {
			p.vehicles[i] = nil
			if hours <= 2 {
				return "Registration number " + vehicleID + " with Slot Number " + fmt.Sprint(i+1) + " is free with Charge $" + fmt.Sprint(hours*10)
			}
			return "Registration number " + vehicleID + " with Slot Number " + fmt.Sprint(i+1) + " is free with Charge $" + fmt.Sprint((hours-1)*10)
		}
	}
	return "Registration number " + vehicleID + " not found"
}

func (p *ParkingIOT) status() {
	fmt.Println("Slot No.    Registration No.")
	for i, slot := range p.vehicles {
		if slot != nil {
			fmt.Printf("%d          %s\n", i+1, slot.vehicleID)
		}
	}
}

func main() {
	parkingIOT := create_parking_iot(6)
	// Example usage
	fmt.Println(parkingIOT.parkVehicle("KA-01-HH-1234"))
	fmt.Println(parkingIOT.parkVehicle("KA-01-HH-9999"))
	fmt.Println(parkingIOT.parkVehicle("KA-01-BB-0001"))
	fmt.Println(parkingIOT.parkVehicle("KA-01-HH-7777"))
	fmt.Println(parkingIOT.parkVehicle("KA-01-HH-2701"))
	fmt.Println(parkingIOT.parkVehicle("KA-01-HH-3141"))
	fmt.Println(parkingIOT.removeVehicle("KA-01-HH-3141", 4))
	parkingIOT.status()
	fmt.Println(parkingIOT.parkVehicle("KA-01-P-333"))
	fmt.Println(parkingIOT.parkVehicle("DL-12-AA-9999"))
	fmt.Println(parkingIOT.removeVehicle("KA-01-HH-1234", 4))
	fmt.Println(parkingIOT.removeVehicle("KA-01-BB-0001", 6))
	fmt.Println(parkingIOT.removeVehicle("DL-12-AA-9999", 2))
	fmt.Println(parkingIOT.parkVehicle("KA-09-HH-0987"))
	fmt.Println(parkingIOT.parkVehicle("CA-09-IO-1111"))
	fmt.Println(parkingIOT.parkVehicle("KA-09-HH-0123"))
	parkingIOT.status()
}
