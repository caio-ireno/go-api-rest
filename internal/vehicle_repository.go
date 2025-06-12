package internal

// VehicleRepository is an interface that represents a vehicle repository
type VehicleRepository interface {
	// FindAll is a method that returns a map of all vehicles
	FindAll() (v map[int]Vehicle, err error)
	Save(vh *VehicleAttributes) (v Vehicle, err error)
	FindByColorAndYears(color, year string) (v map[int]Vehicle, err error)
	FindByMarcaAndYearInterval(brand, start_year, end_year string) (v map[int]Vehicle, err error)
}
