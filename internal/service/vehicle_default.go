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

func (s *VehicleDefault) UpdateFuel(id int, fuel string) (v internal.Vehicle, err error) {
	v, err = s.rp.UpdateFuel(id, fuel)

	if err != nil {
		return
	}

	return
}
func (s *VehicleDefault) FindById(id string) (v internal.Vehicle, err error) {
	v, err = s.rp.FindById(id)

	if err != nil {
		return
	}

	return
}

func (s *VehicleDefault) DeleteById(id string) (err error) {
	err = s.rp.DeleteById(id)

	if err != nil {
		return
	}
	return
}

func (s *VehicleDefault) FindByTransmissionType(typeTransmission string) (v map[int]internal.Vehicle, err error) {

	v, err = s.rp.FindByTransmissionType(typeTransmission)

	if err != nil {
		err = apperrors.ErrInvalidVehicleData
		return
	}

	if len(v) == 0 {
		err = apperrors.ErrVehicleNotFound
		return
	}

	return
}

func (s *VehicleDefault) FindTipoCombustivel(FuelType string) (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.FindTipoCombustivel(FuelType)

	if err != nil {
		return
	}

	if len(v) == 0 {
		err = apperrors.ErrVehicleNotFound
	}

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

func (s *VehicleDefault) SaveMultipleVehicles(vh *[]internal.VehicleAttributes) (v map[int]internal.Vehicle, err error) {

	for _, vehicle := range *vh {
		err = vehicle.Validate()
	}
	if err != nil {
		return
	}

	vehicleExisting, _ := s.rp.FindAll()

	// criar um MAP com registro como chave
	registroMap := make(map[string]struct{}) // é menos do que veiculos existentes  pois tem registros repetidos
	// registroError := make(map[string]string)
	for _, ve := range vehicleExisting {
		registroMap[ve.Registration] = struct{}{}
	}

	v = make(map[int]internal.Vehicle)

	for _, vehicle := range *vh {

		// similar ao in no python
		_, exists := registroMap[vehicle.Registration]

		if exists {
			err = apperrors.ErrVehicleAlreadyExists
			return
		}

		// Salve cada veículo individualmente
		saved, saveErr := s.rp.Save(&vehicle)

		if saveErr != nil {
			err = saveErr
			return
		}
		v[saved.Id] = saved

		// Atualize o map para evitar duplicidade entre os novos também
		registroMap[vehicle.Registration] = struct{}{}
	}

	return
}

func (s *VehicleDefault) Patch(vh *internal.Vehicle) (v internal.Vehicle, err error) {

	err = vh.VehicleAttributes.Validate()

	if err != nil {
		return
	}

	v, err = s.rp.Patch(vh)
	return
}

func (s *VehicleDefault) UpdateMaxSpeed(id int, maxSpeed float64) (v internal.Vehicle, err error) {

	v, err = s.rp.UpdateMaxSpeed(id, maxSpeed)
	if err != nil {
		err = apperrors.ErrVehicleNotFound
	}
	return
}
