package feestrategy

type FeeByType float64

const (
	// All unites in INR
	SMALL_VEHICLE_CHARGES  FeeByType = 10
	MEDIUM_VEHICLE_CHARGES FeeByType = 20
	LARGE_VEHICLE_CHARGES  FeeByType = 30
)

type IFeeStrategy interface {
	CalculateFee(hours int) float64
}

type SmallVehicleFee struct {
}

func (s *SmallVehicleFee) CalculateFee(hours int) float64 {
	return float64(SMALL_VEHICLE_CHARGES) * float64(hours)
}

type MediumVehicleFee struct {
}

func (m *MediumVehicleFee) CalculateFee(hours int) float64 {
	return float64(MEDIUM_VEHICLE_CHARGES) * float64(hours)
}

type LargeVehicleFee struct {
}

func (l *LargeVehicleFee) CalculateFee(hours int) float64 {
	return float64(LARGE_VEHICLE_CHARGES) * float64(hours)
}
