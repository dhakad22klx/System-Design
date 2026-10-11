package main

type ParkingSystem struct {
	Levels []*Level
}

type IParkingSystem interface {
	// AddLevel adds a new level to the parking system.
	AddLevel(level *Level)

	// ParkVehicle parks a vehicle in the parking system.
	ParkVehicle(vehicle *Vehicle) (string, error)

	ChargeParkingFee(vehicleNumber string) (float64, error)

	// UnparkVehicle unparks a vehicle from the parking system.
	UnparkVehicle(vehicleNumber string) error

	// getAvailableLots returns the available lots for a given vehicle type.
	getAvailableLots(vehicleType VehicleType) []*Lot
}

func NewParkingSystem(levels []*Level) *ParkingSystem {
	return &ParkingSystem{
		Levels: levels,
	}
}
