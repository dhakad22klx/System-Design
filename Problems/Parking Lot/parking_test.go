package main

import (
	"fmt"
	feestrategy "parkinglot/fee-strategy"
	"sync"
	"testing"
	"time"
)

func TestParkingSystemParkVehicleConcurrently(t *testing.T) {
	const (
		levelCount      = 3
		lotsPerSize     = 8
		capacityPerType = levelCount * lotsPerSize
	)

	vehicleTypes := []struct {
		vehicleType VehicleType
		feeStrategy feestrategy.IFeeStrategy
	}{
		{SMALL_VEHICLE, &feestrategy.SmallVehicleFee{}},
		{MEDIUM_VEHICLE, &feestrategy.MediumVehicleFee{}},
		{LARGE_VEHICLE, &feestrategy.LargeVehicleFee{}},
	}

	for _, tc := range []struct {
		name            string
		vehiclesPerType int
	}{
		{name: "at capacity", vehiclesPerType: capacityPerType},
		{name: "over capacity", vehiclesPerType: capacityPerType * 2},
	} {
		t.Run(tc.name, func(t *testing.T) {
			// Build a fresh environment with every lot size on each level.
			levels := make([]*Level, 0, levelCount)
			lotsByID := make(map[string]*Lot)
			for levelIndex := 1; levelIndex <= levelCount; levelIndex++ {
				lotsBySize := make([][]*Lot, len(vehicleTypes))
				for sizeIndex := range lotsBySize {
					for lotIndex := 1; lotIndex <= lotsPerSize; lotIndex++ {
						id := fmt.Sprintf("L%d-S%d-%d", levelIndex, sizeIndex+1, lotIndex)
						lot := NewLot(id, LotSize(sizeIndex+1))
						lotsBySize[sizeIndex] = append(lotsBySize[sizeIndex], lot)
						lotsByID[id] = lot
					}
				}
				levels = append(levels, NewLevel(fmt.Sprint(levelIndex),
					lotsBySize[0], lotsBySize[1], lotsBySize[2]))
			}
			var parkingSystem IParkingSystem = NewParkingSystem(levels[:1])
			for _, level := range levels[1:] {
				parkingSystem.AddLevel(level)
			}

			type parkingResult struct {
				vehicle *Vehicle
				ticket  string
				err     error
			}
			vehicleCount := tc.vehiclesPerType * len(vehicleTypes)
			results := make(chan parkingResult, vehicleCount)
			start := make(chan struct{})
			var ready, workers sync.WaitGroup
			ready.Add(vehicleCount)
			workers.Add(vehicleCount)
			for _, spec := range vehicleTypes {
				for vehicleIndex := 0; vehicleIndex < tc.vehiclesPerType; vehicleIndex++ {
					vehicle := NewVehicle(fmt.Sprintf("V%d-%d", spec.vehicleType, vehicleIndex),
						spec.vehicleType, spec.feeStrategy)
					go func() {
						defer workers.Done()
						ready.Done()
						<-start
						ticket, err := parkingSystem.ParkVehicle(vehicle)
						results <- parkingResult{vehicle: vehicle, ticket: ticket, err: err}
					}()
				}
			}
			ready.Wait()
			startedAt := time.Now()
			close(start)
			workers.Wait()
			finishedAt := time.Now()
			close(results)

			parkedByType := make(map[VehicleType]int)
			rejectedByType := make(map[VehicleType]int)
			assignedLots := make(map[string]string)
			for result := range results {
				vehicle := result.vehicle
				if result.err != nil {
					rejectedByType[vehicle.Type]++
					if result.ticket != "" || vehicle.AssignedLot != nil || !vehicle.VehicleEntryTime.IsZero() {
						t.Errorf("rejected vehicle %s has a ticket, assigned lot, or entry time", vehicle.VehicleNumber)
					}
					continue
				}

				parkedByType[vehicle.Type]++
				lot, exists := lotsByID[result.ticket]
				if !exists {
					t.Errorf("vehicle %s received unknown lot %q", vehicle.VehicleNumber, result.ticket)
					continue
				}
				if previousVehicle, duplicate := assignedLots[result.ticket]; duplicate {
					t.Errorf("lot %s assigned to both %s and %s", result.ticket, previousVehicle, vehicle.VehicleNumber)
				}
				assignedLots[result.ticket] = vehicle.VehicleNumber
				if vehicle.AssignedLot != lot {
					t.Errorf("vehicle %s assigned lot does not match ticket %s", vehicle.VehicleNumber, result.ticket)
				}
				if lot.GetSize() != LotSize(vehicle.Type) {
					t.Errorf("vehicle %s of type %d parked in lot of size %d", vehicle.VehicleNumber, vehicle.Type, lot.GetSize())
				}
				if vehicle.VehicleEntryTime.Before(startedAt) || vehicle.VehicleEntryTime.After(finishedAt) {
					t.Errorf("vehicle %s entry time is outside the parking interval", vehicle.VehicleNumber)
				}
			}

			for _, spec := range vehicleTypes {
				if got := parkedByType[spec.vehicleType]; got != capacityPerType {
					t.Errorf("type %d: parked %d vehicles, want %d", spec.vehicleType, got, capacityPerType)
				}
				wantRejected := tc.vehiclesPerType - capacityPerType
				if got := rejectedByType[spec.vehicleType]; got != wantRejected {
					t.Errorf("type %d: rejected %d vehicles, want %d", spec.vehicleType, got, wantRejected)
				}
			}
			for id, lot := range lotsByID {
				if !lot.IsLocked() {
					t.Errorf("lot %s is unlocked after filling the parking system", id)
				}
			}
		})
	}
}
