package apperrors

import "errors"

var (
	ErrVehicleWithCriteria  = errors.New("no vehicle found with these criteria")
	ErrVehicleAlreadyExists = errors.New("vehicle identifier already exists")
	ErrInvalidVehicleData   = errors.New("required or invalid vehicle data")
)
