package handler

import (
	"app/internal"
	"app/pkg/apperrors"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
)

// VehicleJSON is a struct that represents a vehicle in JSON format
type VehicleJSON struct {
	ID              int     `json:"id"`
	Brand           string  `json:"brand"`
	Model           string  `json:"model"`
	Registration    string  `json:"registration"`
	Color           string  `json:"color"`
	FabricationYear int     `json:"year"`
	Capacity        int     `json:"passengers"`
	MaxSpeed        float64 `json:"max_speed"`
	FuelType        string  `json:"fuel_type"`
	Transmission    string  `json:"transmission"`
	Weight          float64 `json:"weight"`
	Height          float64 `json:"height"`
	Length          float64 `json:"length"`
	Width           float64 `json:"width"`
}

// NewVehicleDefault is a function that returns a new instance of VehicleDefault
func NewVehicleDefault(sv internal.VehicleService) *VehicleDefault {
	return &VehicleDefault{sv: sv}
}

// VehicleDefault is a struct with methods that represent handlers for vehicles
type VehicleDefault struct {
	// sv is the service that will be used by the handler
	sv internal.VehicleService
}

// GetAll is a method that returns a handler for the route GET /vehicles
func (h *VehicleDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		v, err := h.sv.FindAll()
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, nil)
			return
		}

		// response
		data := make(map[int]VehicleJSON)
		for key, value := range v {
			data[key] = VehicleJSON{
				ID:              value.Id,
				Brand:           value.Brand,
				Model:           value.Model,
				Registration:    value.Registration,
				Color:           value.Color,
				FabricationYear: value.FabricationYear,
				Capacity:        value.Capacity,
				MaxSpeed:        value.MaxSpeed,
				FuelType:        value.FuelType,
				Transmission:    value.Transmission,
				Weight:          value.Weight,
				Height:          value.Height,
				Length:          value.Length,
				Width:           value.Width,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

func (h *VehicleDefault) GetByColorAndYears() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		color := r.URL.Query().Get("color")
		year := r.URL.Query().Get("year")

		fmt.Println("Query parans", color, year)

		v, err := h.sv.FindByColorAndYears(color, year)

		if err != nil {
			if errors.Is(err, apperrors.ErrVehicleWithCriteria) {
				response.JSON(w, http.StatusNotFound, map[string]any{
					"message": "Nenhum veículo encontrado com esses critérios.",
					"data":    nil,
				})
				return
			}

			response.JSON(w, http.StatusInternalServerError, map[string]any{
				"message": err.Error(),
				"data":    nil,
			})
			return
		}

		data := make(map[int]VehicleJSON)
		for key, value := range v {
			data[key] = VehicleJSON{
				ID:              value.Id,
				Brand:           value.Brand,
				Model:           value.Model,
				Registration:    value.Registration,
				Color:           value.Color,
				FabricationYear: value.FabricationYear,
				Capacity:        value.Capacity,
				MaxSpeed:        value.MaxSpeed,
				FuelType:        value.FuelType,
				Transmission:    value.Transmission,
				Weight:          value.Weight,
				Height:          value.Height,
				Length:          value.Length,
				Width:           value.Width,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})

	}
}

func (h *VehicleDefault) GetVelocidadeMediaMarca() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		brand := chi.URLParam(r, "brand")

		fmt.Println("Query parans", brand)

		m, err := h.sv.FindVelocidadeMediaMarca(brand)

		if err != nil {
			if errors.Is(err, apperrors.ErrVehicleBrand) {
				response.JSON(w, http.StatusNotFound, map[string]any{
					"message": "Nenhuma marca encontrado.",
					"data":    nil,
				})
				return
			}

			response.JSON(w, http.StatusInternalServerError, map[string]any{
				"message": err.Error(),
				"data":    nil,
			})
			return
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    m,
		})

	}
}

func (h *VehicleDefault) GetByMarcaAndYearInterval() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		brand := chi.URLParam(r, "brand")
		start_year := chi.URLParam(r, "start_year")
		end_year := chi.URLParam(r, "end_year")

		fmt.Println("Query parans", brand, start_year, end_year)

		v, err := h.sv.FindByMarcaAndYearInterval(brand, start_year, end_year)

		if err != nil {
			if errors.Is(err, apperrors.ErrVehicleWithCriteria) {
				response.JSON(w, http.StatusNotFound, map[string]any{
					"message": "Nenhum veículo encontrado com esses critérios.",
					"data":    nil,
				})
				return
			}

			response.JSON(w, http.StatusInternalServerError, map[string]any{
				"message": err.Error(),
				"data":    nil,
			})
			return
		}

		data := make(map[int]VehicleJSON)
		for key, value := range v {
			data[key] = VehicleJSON{
				ID:              value.Id,
				Brand:           value.Brand,
				Model:           value.Model,
				Registration:    value.Registration,
				Color:           value.Color,
				FabricationYear: value.FabricationYear,
				Capacity:        value.Capacity,
				MaxSpeed:        value.MaxSpeed,
				FuelType:        value.FuelType,
				Transmission:    value.Transmission,
				Weight:          value.Weight,
				Height:          value.Height,
				Length:          value.Length,
				Width:           value.Width,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})

	}
}

func (h *VehicleDefault) Save() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var reqBody internal.VehicleAttributes

		err := json.NewDecoder(r.Body).Decode(&reqBody)

		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": "bad request: Dados do veículo mal formatados ou incompletos.",
				"data":    nil,
			})
			return
		}

		v, err := h.sv.Save(&reqBody)
		if err != nil {
			if errors.Is(err, apperrors.ErrVehicleAlreadyExists) {
				response.JSON(w, http.StatusConflict, map[string]any{
					"message": "Identificador do veículo já existente",
					"data":    nil,
				})
				return
			}

			response.JSON(w, http.StatusInternalServerError, map[string]any{
				"message": err.Error(),
				"data":    nil,
			})
			return
		}

		response.JSON(w, http.StatusCreated, map[string]any{
			"message": "Veículo criado com sucesso.",
			"data":    v,
		})
	}
}
