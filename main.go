package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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

func (p *ParkingIOT) park(car_number string) string {
	for i := 0; i < len(p.vehicles); i++ {
		if p.vehicles[i] == nil {
			p.vehicles[i] = &Vehicle{vehicleID: car_number}
			return "Allocated slot number: " + fmt.Sprint(i+1)
		}
	}
	return "Sorry, parking iot is full"
}

func (p *ParkingIOT) leave(car_number string, hours int) string {
	for i := 0; i < len(p.vehicles); i++ {
		if p.vehicles[i] != nil && p.vehicles[i].vehicleID == car_number {
			p.vehicles[i] = nil
			if hours <= 2 {
				return "Registration number " + car_number + " with Slot Number " + fmt.Sprint(i+1) + " is free with Charge $" + fmt.Sprint(hours*10)
			}
			return "Registration number " + car_number + " with Slot Number " + fmt.Sprint(i+1) + " is free with Charge $" + fmt.Sprint((hours-1)*10)
		}
	}
	return "Registration number " + car_number + " not found"
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
	// Init size of parking iot
	// parkingIOT := create_parking_iot(6)

	// Command
	// fmt.Println(parkingIOT.park("KA-01-HH-1234"))
	// fmt.Println(parkingIOT.park("KA-01-HH-9999"))
	// fmt.Println(parkingIOT.park("KA-01-BB-0001"))
	// fmt.Println(parkingIOT.park("KA-01-HH-7777"))
	// fmt.Println(parkingIOT.park("KA-01-HH-2701"))
	// fmt.Println(parkingIOT.park("KA-01-HH-3141"))
	// fmt.Println(parkingIOT.leave("KA-01-HH-3141", 4))
	// parkingIOT.status()
	// fmt.Println(parkingIOT.park("KA-01-P-333"))
	// fmt.Println(parkingIOT.park("DL-12-AA-9999"))
	// fmt.Println(parkingIOT.leave("KA-01-HH-1234", 4))
	// fmt.Println(parkingIOT.leave("KA-01-BB-0001", 6))
	// fmt.Println(parkingIOT.leave("DL-12-AA-9999", 2))
	// fmt.Println(parkingIOT.park("KA-09-HH-0987"))
	// fmt.Println(parkingIOT.park("CA-09-IO-1111"))
	// fmt.Println(parkingIOT.park("KA-09-HH-0123"))
	// parkingIOT.status()

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <input_file>")
		return
	}

	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var parking *ParkingIOT
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Fields(line)

		if len(tokens) == 0 {
			continue
		}

		switch tokens[0] {
		case "create_parking_lot":
			if len(tokens) < 2 {
				fmt.Println("Usage: create_parking_lot <size>")
				continue
			}
			var size int
			fmt.Sscanf(tokens[1], "%d", &size)
			parking = create_parking_iot(size)
			fmt.Printf("Created a parking lot with %d slots\n", size)

		case "park":
			if parking == nil {
				fmt.Println("Parking lot not created yet")
				continue
			}
			if len(tokens) < 2 {
				fmt.Println("Usage: park <vehicle_id>")
				continue
			}
			fmt.Println(parking.park(tokens[1]))

		case "leave":
			if parking == nil {
				fmt.Println("Parking lot not created yet")
				continue
			}
			if len(tokens) < 3 {
				fmt.Println("Usage: leave <vehicle_id> <hours>")
				continue
			}
			var hours int
			fmt.Sscanf(tokens[2], "%d", &hours)
			fmt.Println(parking.leave(tokens[1], hours))

		case "status":
			if parking == nil {
				fmt.Println("Parking lot not created yet")
				continue
			}
			parking.status()

		default:
			fmt.Println("Unknown command:", tokens[0])
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
