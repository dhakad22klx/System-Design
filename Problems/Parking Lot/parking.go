package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

type ParkingSystem struct {
	mu     sync.Mutex
	Levels []*Level
}

type IParkingSystem interface {
	// AddLevel adds a new level to the parking system.
	AddLevel(level *Level)

	// ParkVehicle parks a vehicle in the parking system.
	ParkVehicle(vehicle *Vehicle) (string, error)

	ChargeParkingFee(vehicle *Vehicle) (float64, error)

	// UnparkVehicle unparks a vehicle from the parking system.
	UnparkVehicle(vehicle *Vehicle) error

	// getAvailableLots returns the available lots for a given vehicle type.
	getAvailableLot(vehicle *Vehicle) *Lot
}

func NewParkingSystem(levels []*Level) *ParkingSystem {
	return &ParkingSystem{
		Levels: levels,
	}
}

func (ps *ParkingSystem) AddLevel(level *Level) {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	ps.Levels = append(ps.Levels, level)
}

func (ps *ParkingSystem) ParkVehicle(vehicle *Vehicle) (string, error) {
	// Selecting and reserving a lot must happen as one operation.
	ps.mu.Lock()
	defer ps.mu.Unlock()

	availableLot := ps.getAvailableLot(vehicle)
	if availableLot == nil {
		return "", fmt.Errorf("no available lots for vehicle type: %d", vehicle.Type)
	}
	availableLot.Lock()
	vehicle.AssignedLot = availableLot
	vehicle.VehicleEntryTime = time.Now()
	return availableLot.GetId(), nil
}

func (ps *ParkingSystem) UnparkVehicle(vehicle *Vehicle) error {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	var lotToVacate *Lot = vehicle.AssignedLot
	lotToVacate.Unlock()
	return nil
}

func (ps *ParkingSystem) ChargeParkingFee(vehicle *Vehicle) (float64, error) {
	vehicle.VehicleExitTime = time.Now()
	var hoursParked = int(math.Ceil(vehicle.VehicleExitTime.Sub(vehicle.VehicleEntryTime).Hours()))
	return vehicle.ChargeVehicle(hoursParked), nil
}

// getAvailableLot must be called with ps.mu held.
func (ps *ParkingSystem) getAvailableLot(vehicle *Vehicle) *Lot {
	var availableLot *Lot = nil
	if vehicle.Type == SMALL_VEHICLE {
		for _, level := range ps.Levels {
			for _, lot := range level.SmallLots {
				if !lot.IsLocked() {
					availableLot = lot
					break
				}
			}
			if availableLot != nil {
				break
			}
		}
	} else if vehicle.Type == MEDIUM_VEHICLE {
		for _, level := range ps.Levels {
			for _, lot := range level.MediumLots {
				if !lot.IsLocked() {
					availableLot = lot
					break
				}
			}
			if availableLot != nil {
				break
			}
		}
	} else if vehicle.Type == LARGE_VEHICLE {
		for _, level := range ps.Levels {
			for _, lot := range level.LargeLots {
				if !lot.IsLocked() {
					availableLot = lot
					break
				}
			}
			if availableLot != nil {
				break
			}
		}
	}
	return availableLot
}
