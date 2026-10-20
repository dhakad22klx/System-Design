package main

import (
	feestrategy "parkinglot/fee-strategy"
	"time"
)

type VehicleType int

const (
	SMALL_VEHICLE  VehicleType = 1
	MEDIUM_VEHICLE VehicleType = 2
	LARGE_VEHICLE  VehicleType = 3
)

type Vehicle struct {
	VehicleNumber    string
	Type             VehicleType
	FeeStrategy      feestrategy.IFeeStrategy
	AssignedLot      *Lot
	VehicleEntryTime time.Time
	VehicleExitTime  time.Time
}

type IVehicle interface {
	//ChargeVehicle charges the vehicle based on the parking fee.
	ChargeVehicle(hours int) float64
}

func NewVehicle(vehicleNumber string, vehicleType VehicleType, feeStrategy feestrategy.IFeeStrategy) *Vehicle {
	return &Vehicle{
		VehicleNumber: vehicleNumber,
		Type:          vehicleType,
		FeeStrategy:   feeStrategy,
	}
}

func (v *Vehicle) ChargeVehicle(hours int) float64 {
	if v.FeeStrategy == nil {
		panic("Fee strategy not set for vehicle type: " + string(v.VehicleNumber))
	}
	return v.FeeStrategy.CalculateFee(hours)
}
