package main

import (
	"fmt"
	feestrategy "parkinglot/fee-strategy"
)

const (
	SMALL  int = 1
	MEDIUM int = 2
	LARGE  int = 3
)

func main() {
	fmt.Println("Parking Lot Problems")
	var parkingSystem IParkingSystem = NewParkingSystem([]*Level{
		NewLevel("1", []*Lot{
			NewLot("L1-1", LotSize(SMALL)),
		}, []*Lot{
			NewLot("L1-2", LotSize(MEDIUM)),
		}, []*Lot{
			NewLot("L1-3", LotSize(LARGE)),
		}),
		NewLevel("2", []*Lot{
			NewLot("L2-1", LotSize(SMALL)),
		}, []*Lot{
			NewLot("L2-2", LotSize(MEDIUM)),
		}, []*Lot{
			NewLot("L2-3", LotSize(LARGE)),
		}),
	})

	vehicle1 := NewVehicle("KA-01-HH-1234", SMALL_VEHICLE, &feestrategy.SmallVehicleFee{})
	vehicle2 := NewVehicle("KA-01-HH-9999", MEDIUM_VEHICLE, &feestrategy.MediumVehicleFee{})
	vehicle3 := NewVehicle("KA-01-BB-0001", LARGE_VEHICLE, &feestrategy.LargeVehicleFee{})

	if ticket, err := parkingSystem.ParkVehicle(vehicle1); err == nil {
		fmt.Printf("Vehicle %s parked at lot %s\n", vehicle1.VehicleNumber, ticket)
	} else {
		fmt.Printf("Failed to park vehicle %s: %v\n", vehicle1.VehicleNumber, err)
	}

	if ticket, err := parkingSystem.ParkVehicle(vehicle2); err == nil {
		fmt.Printf("Vehicle %s parked at lot %s\n", vehicle2.VehicleNumber, ticket)
	} else {
		fmt.Printf("Failed to park vehicle %s: %v\n", vehicle2.VehicleNumber, err)
	}

	if ticket, err := parkingSystem.ParkVehicle(vehicle3); err == nil {
		fmt.Printf("Vehicle %s parked at lot %s\n", vehicle3.VehicleNumber, ticket)
	} else {
		fmt.Printf("Failed to park vehicle %s: %v\n", vehicle3.VehicleNumber, err)
	}

	if fee, err := parkingSystem.ChargeParkingFee(vehicle1); err == nil {
		fmt.Printf("Parking fee for vehicle %s: %.2f\n", vehicle1.VehicleNumber, fee)
	} else {
		fmt.Printf("Failed to charge parking fee for vehicle %s: %v\n", vehicle1.VehicleNumber, err)
	}

	if fee, err := parkingSystem.ChargeParkingFee(vehicle2); err == nil {
		fmt.Printf("Parking fee for vehicle %s: %.2f\n", vehicle2.VehicleNumber, fee)
	} else {
		fmt.Printf("Failed to charge parking fee for vehicle %s: %v\n", vehicle2.VehicleNumber, err)
	}

	if fee, err := parkingSystem.ChargeParkingFee(vehicle3); err == nil {
		fmt.Printf("Parking fee for vehicle %s: %.2f\n", vehicle3.VehicleNumber, fee)
	} else {
		fmt.Printf("Failed to charge parking fee for vehicle %s: %v\n", vehicle3.VehicleNumber, err)
	}
}
