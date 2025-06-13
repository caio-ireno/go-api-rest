package repository

import (
	"app/internal"
	"app/pkg/apperrors"
	"app/pkg/utils"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// NewVehicleMap is a function that returns a new instance of VehicleMap
func NewVehicleMap(db map[int]internal.Vehicle) *VehicleMap {
	// default db
	defaultDb := make(map[int]internal.Vehicle)
	if db != nil {
		defaultDb = db
	}
	return &VehicleMap{db: defaultDb}
}

// VehicleMap is a struct that represents a vehicle repository
type VehicleMap struct {
	db map[int]internal.Vehicle
}

// FindAll is a method that returns a map of all vehicles
func (r *VehicleMap) FindAll() (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	// copy db
	for key, value := range r.db {
		v[key] = value
	}

	return
}

func (r *VehicleMap) FindById(id string) (v internal.Vehicle, err error) {
	fmt.Println("Query parans", id)
	idInt, err := strconv.Atoi(id)

	if err != nil {
		return v, fmt.Errorf("invalid ID: %w", err)
	}

	for _, value := range r.db {
		if value.Id == idInt {
			v = value
		}
	}
	return
}

func (r *VehicleMap) DeleteById(id string) (err error) {
	idInt, err := strconv.Atoi(id)

	if err != nil {
		err = errors.New("id incorreto ou mal formatado")
		return
	}

	_, ok := r.db[idInt]
	if !ok {
		err = apperrors.ErrVehicleNotFound
		return
	}

	delete(r.db, idInt)

	return

}

func (r *VehicleMap) FindByTransmissionType(typeTransmission string) (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	for key, value := range r.db {
		if value.VehicleAttributes.Transmission == typeTransmission {
			v[key] = value
		}
	}
	return

}

func (r *VehicleMap) UpdateFuel(id int, fuel string) (v internal.Vehicle, err error) {

	_, ok := r.db[id]

	if !ok {
		err = apperrors.ErrVehicleNotFound
	}

	v = r.db[id]
	v.FuelType = fuel

	return

}
func (r *VehicleMap) FindTipoCombustivel(FuelType string) (v map[int]internal.Vehicle, err error) {
	fmt.Println("Query parans", FuelType)

	v = make(map[int]internal.Vehicle)

	for key, value := range r.db {
		if value.VehicleAttributes.FuelType == FuelType {
			v[key] = value
		}
	}
	return
}

func (r *VehicleMap) FindVelocidadeMediaMarca(brand string) (m float64, err error) {
	fmt.Println("Query parans", brand)
	brandCaptalize := utils.CapitalizeFirst(brand)

	sum := 0.0
	count := 0
	for _, value := range r.db {
		if value.Brand == brandCaptalize {
			sum += value.MaxSpeed
			count += 1

		}
	}

	if count == 0 {
		m = 0
		return
	}
	m = sum / float64(count)

	return
}

func (r *VehicleMap) FindByPeso(min, max string) (v map[int]internal.Vehicle, err error) {

	minFloat, err := strconv.ParseFloat(min, 64)
	v = make(map[int]internal.Vehicle)

	if err != nil {
		err = errors.New("erro ao converter dados")
		return
	}

	maxFloat, err := strconv.ParseFloat(max, 64)

	if err != nil {
		err = errors.New("erro ao converter dados")
		return
	}

	for key, value := range r.db {
		if value.VehicleAttributes.Weight >= minFloat &&
			value.VehicleAttributes.Weight <= maxFloat {
			v[key] = value
		}
	}
	if len(v) == 0 {
		err = apperrors.ErrVehicleNotFound
		return
	}

	return
}

// FindAll is a method that returns a map of all vehicles
func (r *VehicleMap) FindByMarcaAndYearInterval(brand, start_year, end_year string) (v map[int]internal.Vehicle, err error) {
	fmt.Println("Query parans", brand, start_year, end_year)
	brandCaptalize := utils.CapitalizeFirst(brand)

	v = make(map[int]internal.Vehicle)

	startYearInt, err := strconv.Atoi(start_year)
	if err != nil {
		return v, fmt.Errorf("invalid start_year: %w", err)
	}
	endYearInt, err := strconv.Atoi(end_year)
	if err != nil {
		return v, fmt.Errorf("invalid end_year: %w", err)
	}

	for key, value := range r.db {
		if value.Brand == brandCaptalize &&
			value.FabricationYear >= startYearInt &&
			value.FabricationYear <= endYearInt {
			v[key] = value
		}
	}

	return
}

func (r *VehicleMap) FindByColorAndYears(color, year string) (v map[int]internal.Vehicle, err error) {

	v = make(map[int]internal.Vehicle)

	yearInt, err := strconv.Atoi(year)

	if err != nil {
		return v, errors.New("erro ao converter parametro year")
	}

	for key, value := range r.db {
		if value.Color == color && value.FabricationYear == yearInt {
			v[key] = value
		}
	}

	return
}

func (r *VehicleMap) Save(vh *internal.VehicleAttributes) (v internal.Vehicle, err error) {
	attr := internal.Vehicle{
		VehicleAttributes: internal.VehicleAttributes{

			Brand:        vh.Brand,
			Model:        vh.Model,
			Registration: vh.Registration,

			Color:           vh.Color,
			FabricationYear: vh.FabricationYear,
			Capacity:        vh.Capacity,
			MaxSpeed:        vh.MaxSpeed,
			FuelType:        vh.FuelType,
			Transmission:    vh.Transmission,
			Weight:          vh.Weight,
			Dimensions: internal.Dimensions{
				Height: vh.Height,
				Width:  vh.Weight,
				Length: vh.Length,
			},
		},
	}

	attr.Id = len(r.db) + 1
	r.db[attr.Id] = attr

	v = attr

	return
}

func (r *VehicleMap) Patch(vh *internal.Vehicle) (v internal.Vehicle, err error) {
	attr := internal.Vehicle{
		Id: vh.Id,
		VehicleAttributes: internal.VehicleAttributes{

			Brand:        vh.Brand,
			Model:        vh.Model,
			Registration: vh.Registration,

			Color:           vh.Color,
			FabricationYear: vh.FabricationYear,
			Capacity:        vh.Capacity,
			MaxSpeed:        vh.MaxSpeed,
			FuelType:        vh.FuelType,
			Transmission:    vh.Transmission,
			Weight:          vh.Weight,
			Dimensions: internal.Dimensions{
				Height: vh.Height,
				Width:  vh.Weight,
				Length: vh.Length,
			},
		},
	}

	r.db[attr.Id] = attr

	v = attr

	return
}

func (r *VehicleMap) UpdateMaxSpeed(id int, maxSpeed float64) (v internal.Vehicle, err error) {

	vehicle, ok := r.db[id]

	if !ok {
		err = fmt.Errorf("vehicle with id %d not found", id)
		return
	}

	vehicle.MaxSpeed = maxSpeed
	r.db[id] = vehicle
	v = vehicle

	return
}

func (r *VehicleMap) FindMediaPessoaPorMarca(brand string) (m int, err error) {
	brandCapitalized := utils.CapitalizeFirst(brand)

	var count int
	var sum int

	for _, value := range r.db {
		if value.VehicleAttributes.Brand == brandCapitalized {
			fmt.Println(value)
			count += 1
			sum += (value.VehicleAttributes.Capacity)
		}
	}
	if count == 0 {
		err = apperrors.ErrVehicleBrand
		return
	}

	m = sum / count

	return
}

func (r *VehicleMap) FindByDimenssion(lengthParam, widthParam string) (v map[int]internal.Vehicle, err error) {
	lengthParams := strings.Split(lengthParam, "-")
	widthParams := strings.Split(widthParam, "-")

	v = make(map[int]internal.Vehicle)

	lengthMin, err := strconv.ParseFloat(lengthParams[0], 64)
	if err != nil {
		return v, fmt.Errorf("invalid length min: %w", err)
	}
	lengthMax, err := strconv.ParseFloat(lengthParams[1], 64)
	if err != nil {
		return v, fmt.Errorf("invalid length max: %w", err)
	}

	widthMin, err := strconv.ParseFloat(widthParams[0], 64)
	if err != nil {
		return v, fmt.Errorf("invalid width min: %w", err)
	}
	widthMax, err := strconv.ParseFloat(widthParams[1], 64)
	if err != nil {
		return v, fmt.Errorf("invalid width max: %w", err)
	}

	fmt.Println(lengthParams)
	fmt.Println(widthParams)

	for key, value := range r.db {
		if value.VehicleAttributes.Dimensions.Length >= lengthMin &&
			value.VehicleAttributes.Dimensions.Length <= lengthMax &&
			value.VehicleAttributes.Dimensions.Width >= widthMin &&
			value.VehicleAttributes.Dimensions.Width <= widthMax {
			fmt.Println(value)

			v[key] = value

		}
	}

	if len(v) == 0 {
		err = apperrors.ErrVehicleNotFound
	}

	return
}
