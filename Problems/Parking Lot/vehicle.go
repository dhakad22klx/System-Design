package main

type VehicleType string

const (
	SMALL_VEHICLE  VehicleType = "SMALL"
	MEDIUM_VEHICLE VehicleType = "MEDIUM"
	LARGE_VEHICLE  VehicleType = "LARGE"
)

type Vehicle struct {
	VehicleNumber string
	Type          VehicleType
}

type IVehicle interface {
	//ChargeVehicle charges the vehicle based on the parking fee.
	ChargeVehicle() float64
}

func NewVehicle(vehicleNumber string, vehicleType VehicleType) *Vehicle {
	return &Vehicle{
		VehicleNumber: vehicleNumber,
		Type:          vehicleType,
	}
}
