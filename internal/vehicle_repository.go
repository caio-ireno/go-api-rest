package internal

// VehicleRepository is an interface that represents a vehicle repository
type VehicleRepository interface {
	// FindAll is a method that returns a map of all vehicles
	FindAll() (v map[int]Vehicle, err error)
	Save(vh *VehicleAttributes) (v Vehicle, err error)
	FindByColorAndYears(color, year string) (v map[int]Vehicle, err error)
	FindByMarcaAndYearInterval(brand, start_year, end_year string) (v map[int]Vehicle, err error)
	FindVelocidadeMediaMarca(brand string) (m float64, err error)
	FindByTransmissionType(typeTransmission string) (v map[int]Vehicle, err error)

	FindTipoCombustivel(typeFuel string) (v map[int]Vehicle, err error)

	FindById(id string) (v Vehicle, err error)

	Patch(vh *Vehicle) (v Vehicle, err error)
	UpdateMaxSpeed(id int, maxSpeed float64) (v Vehicle, err error)
	UpdateFuel(id int, fuelType string) (v Vehicle, err error)

	DeleteById(id string) (err error)
}
