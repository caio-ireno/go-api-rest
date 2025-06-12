package repository

import (
	"app/internal"
	"app/pkg/utils"
	"errors"
	"fmt"
	"strconv"
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
