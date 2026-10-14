package main

type ParkingSystem struct {
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
	getAvailableLots(vehicle *Vehicle) []*Lot
}

func NewParkingSystem(levels []*Level) *ParkingSystem {
	return &ParkingSystem{
		Levels: levels,
	}
}
