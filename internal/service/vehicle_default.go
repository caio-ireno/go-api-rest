package service

import (
	"app/internal"
	"app/pkg/apperrors"
)

// NewVehicleDefault is a function that returns a new instance of VehicleDefault
func NewVehicleDefault(rp internal.VehicleRepository) *VehicleDefault {
	return &VehicleDefault{rp: rp}
}

// VehicleDefault is a struct that represents the default service for vehicles
type VehicleDefault struct {
	// rp is the repository that will be used by the service
	rp internal.VehicleRepository
}

// FindAll is a method that returns a map of all vehicles
func (s *VehicleDefault) FindAll() (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.FindAll()
	return
}

func (s *VehicleDefault) FindByColorAndYears(color, year string) (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.FindByColorAndYears(color, year)

	if err != nil {
		return
	}

	if len(v) == 0 {
		err = apperrors.ErrVehicleWithCriteria
	}

	return
}

func (s *VehicleDefault) FindVelocidadeMediaMarca(brand string) (m float64, err error) {
	m, err = s.rp.FindVelocidadeMediaMarca(brand)

	if err != nil {
		return
	}

	if m == 0 {
		err = apperrors.ErrVehicleBrand
		return
	}

	return
}

func (s *VehicleDefault) FindByMarcaAndYearInterval(brand, start_year, end_year string) (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.FindByMarcaAndYearInterval(brand, start_year, end_year)

	if err != nil {
		return
	}

	if len(v) == 0 {
		err = apperrors.ErrVehicleWithCriteria
	}

	return
}

func (s *VehicleDefault) Save(vh *internal.VehicleAttributes) (v internal.Vehicle, err error) {
	err = vh.Validate()

	if err != nil {
		return
	}
	vehicle, _ := s.rp.FindAll()

	for _, vehicle := range vehicle {

		if vehicle.Registration == vh.Registration {
			err = apperrors.ErrVehicleAlreadyExists
			return
		}
	}
	v, err = s.rp.Save(vh)
	return
}
