package apperrors

import "errors"

var (
	ErrVehicleWithCriteria  = errors.New("no vehicle found with these criteria")
	ErrVehicleBrand         = errors.New("no brand found")
	ErrVehicleAlreadyExists = errors.New("vehicle identifier already exists")
	ErrInvalidVehicleData   = errors.New("required or invalid vehicle data")
	ErrVehicleNotFound      = errors.New("vehicle not Found")
)
