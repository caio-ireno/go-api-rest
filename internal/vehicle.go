package internal

import "errors"

// Dimensions is a struct that represents a dimension in 3d
type Dimensions struct {
	// Height is the height of the dimension
	Height float64
	// Length is the length of the dimension
	Length float64
	// Width is the width of the dimension
	Width float64
}

// VehicleAttributes is a struct that represents the attributes of a vehicle
type VehicleAttributes struct {
	// Brand is the brand of the vehicle
	Brand string
	// Model is the model of the vehicle
	Model string
	// Registration is the registration of the vehicle
	Registration string
	// Color is the color of the vehicle
	Color string
	// FabricationYear is the fabrication year of the vehicle
	FabricationYear int
	// Capacity is the capacity of people of the vehicle
	Capacity int
	// MaxSpeed is the maximum speed of the vehicle
	MaxSpeed float64
	// FuelType is the fuel type of the vehicle
	FuelType string
	// Transmission is the transmission of the vehicle
	Transmission string
	// Weight is the weight of the vehicle
	Weight float64
	// Dimensions is the dimensions of the vehicle
	Dimensions
}

// Vehicle is a struct that represents a vehicle
type Vehicle struct {
	// Id is the unique identifier of the vehicle
	Id int

	// VehicleAttribue is the attributes of a vehicle
	VehicleAttributes
}

type UpdateMaxSpeedRequest struct {
	MaxSpeed float64 `json:"MaxSpeed"`
}

func (v *VehicleAttributes) ToDomain() *Vehicle {
	return &Vehicle{
		VehicleAttributes: *v,
	}
}

func (v *VehicleAttributes) Validate() error {
	if v.Brand == "" {
		return errors.New("brand is required")
	}
	if v.Model == "" {
		return errors.New("model is required")
	}
	if v.Registration == "" {
		return errors.New("registration is required")
	}
	if v.Color == "" {
		return errors.New("color is required")
	}
	if v.FabricationYear <= 0 {
		return errors.New("fabrication year must be a positive integer")
	}
	if v.Capacity <= 0 {
		return errors.New("capacity must be a positive integer")
	}
	if v.MaxSpeed <= 0 {
		return errors.New("max speed must be greater than zero")
	}
	if v.FuelType == "" {
		return errors.New("fuel type is required")
	}
	if v.Transmission == "" {
		return errors.New("transmission is required")
	}
	if v.Weight <= 0 {
		return errors.New("weight must be greater than zero")
	}
	if v.Height <= 0 {
		return errors.New("dimensions: height must be greater than zero")
	}
	if v.Width <= 0 {
		return errors.New("dimensions: width must be greater than zero")
	}
	return nil
}
