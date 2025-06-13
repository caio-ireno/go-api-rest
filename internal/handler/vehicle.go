package handler

import (
	"app/internal"
	"app/pkg/apperrors"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

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

func (h *VehicleDefault) GetTipoCombustivel() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fuel_type := chi.URLParam(r, "type")

		fmt.Println("Query parans", fuel_type)

		v, err := h.sv.FindTipoCombustivel(fuel_type)

		if err != nil {
			if errors.Is(err, apperrors.ErrVehicleNotFound) {
				response.JSON(w, http.StatusNotFound, map[string]any{
					"message": "Não foram encontrados veículos com esse tipo de combustível.",
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

func (h *VehicleDefault) DeleteById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		fmt.Println("Query parans", id)

		err := h.sv.DeleteById(id)

		if err != nil {
			if errors.Is(err, apperrors.ErrVehicleNotFound) {
				response.JSON(w, http.StatusNotFound, map[string]any{
					"message": "Veiculo não encontrado",
					"data":    nil,
				})
				return
			}
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": "Erro Interno no servidor",
				"data":    nil,
			})
			return
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "Veículo removido com sucesso.",
		})

	}
}

func (h *VehicleDefault) UpdateFuel() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		idInt, errId := strconv.Atoi(id)

		if errId != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": "bad request:Id mal formatado ou fora do padrão",
				"data":    nil,
			})
			return
		}

		var reqBody internal.UpdateFuel

		err := json.NewDecoder(r.Body).Decode(&reqBody)

		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": "bad request: Tipo de combustivel do veículo mal formatados ou incompletos.",
				"data":    nil,
			})
			return
		}

		vh, err := h.sv.UpdateFuel(idInt, reqBody.FuelType)

		if err != nil {
			if errors.Is(err, apperrors.ErrVehicleNotFound) {
				response.JSON(w, http.StatusNotFound, map[string]any{
					"message": "Veículo não encontrado.",
					"data":    nil,
				})
				return
			}

			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": "bad request: Tipo de combustivel do veículo mal formatados ou incompletos.",
				"data":    nil,
			})
			return
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "Tipo de combustivel do veículo atualizado",
			"data":    vh,
		})

	}
}

func (h *VehicleDefault) GetTransmissionType() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		typeTransmission := chi.URLParam(r, "type")

		fmt.Println("Query parans", typeTransmission)

		v, err := h.sv.FindByTransmissionType(typeTransmission)

		if err != nil {
			if errors.Is(err, apperrors.ErrInvalidVehicleData) {
				response.JSON(w, http.StatusNotFound, map[string]any{
					"message": "Nenhum veículo encontrado com esses critérios.",
					"data":    nil,
				})
				return
			}

			if errors.Is(err, apperrors.ErrVehicleNotFound) {
				response.JSON(w, http.StatusNotFound, map[string]any{
					"message": "Não foram encontrados veículos com esse tipo de transmissão.",
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

func (h *VehicleDefault) SaveMultipleVehicles() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var reqBody []internal.VehicleAttributes

		err := json.NewDecoder(r.Body).Decode(&reqBody)

		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": "bad request: Dados do veículo mal formatados ou incompletos.",
				"data":    nil,
			})
			return
		}

		v, err := h.sv.SaveMultipleVehicles(&reqBody)
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

func (h *VehicleDefault) Patch() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vehicleId := chi.URLParam(r, "id")
		vehicleIdInt, errId := strconv.Atoi(vehicleId)

		if errId != nil {
			response.JSON(w, http.StatusNotFound, map[string]any{
				"message": "Id Errado ou Invalido",
				"data":    nil,
			})
			return
		}

		vehicle, ok := h.sv.FindById(vehicleId)

		if ok != nil {
			if errors.Is(ok, apperrors.ErrVehicleWithCriteria) {
				response.JSON(w, http.StatusNotFound, map[string]any{
					"message": "Nenhum veículo encontrado com esses critérios.",
					"data":    nil,
				})
				return
			}
		}

		reqBody := internal.VehicleAttributes{
			Brand:        vehicle.Brand,
			Model:        vehicle.Model,
			Registration: vehicle.Registration,

			Color:           vehicle.Color,
			FabricationYear: vehicle.FabricationYear,
			Capacity:        vehicle.Capacity,
			MaxSpeed:        vehicle.MaxSpeed,
			FuelType:        vehicle.FuelType,
			Transmission:    vehicle.Transmission,
			Weight:          vehicle.Weight,
			Dimensions: internal.Dimensions{
				Height: vehicle.Height,
				Width:  vehicle.Weight,
			},
		}

		err := json.NewDecoder(r.Body).Decode(&reqBody)

		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": "bad request: Dados do veículo mal formatados ou incompletos.",
				"data":    nil,
			})
			return
		}

		vh := reqBody.ToDomain()
		vh.Id = vehicleIdInt

		fmt.Println(vh)

		v, err := h.sv.Patch(vh)

		if err != nil {
			if errors.Is(err, apperrors.ErrVehicleNotFound) {
				response.JSON(w, http.StatusConflict, map[string]any{
					"message": "Veiculo não encontrado",
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
			"message": "Velocidade do veículo atualizada com sucesso.",
			"data":    v,
		})
	}
}

func (h *VehicleDefault) UpdateMaxSpeed() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vehicleId := chi.URLParam(r, "id")
		vehicleIdInt, errId := strconv.Atoi(vehicleId)

		if errId != nil {
			response.JSON(w, http.StatusNotFound, map[string]any{
				"message": "Id Errado ou Invalido",
				"data":    nil,
			})
			return
		}

		_, ok := h.sv.FindById(vehicleId)

		if ok != nil {
			if errors.Is(ok, apperrors.ErrVehicleWithCriteria) {
				response.JSON(w, http.StatusNotFound, map[string]any{
					"message": "Nenhum veículo encontrado com esses critérios.",
					"data":    nil,
				})
				return
			}
		}

		var reqBody internal.UpdateMaxSpeedRequest

		err := json.NewDecoder(r.Body).Decode(&reqBody)

		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": "bad request: Velocidade do veículo mal formatados ou incompletos.",
				"data":    nil,
			})
			return
		}

		v, err := h.sv.UpdateMaxSpeed(vehicleIdInt, reqBody.MaxSpeed)

		if err != nil {
			if errors.Is(err, apperrors.ErrVehicleNotFound) {
				response.JSON(w, http.StatusConflict, map[string]any{
					"message": "Veiculo não encontrado",
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
			"message": "Velocidade do veículo atualizada com sucesso.",
			"data":    v,
		})
	}
}

func (h *VehicleDefault) GetMediaPessoaPorMarca() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		brand := chi.URLParam(r, "brand")

		m, err := h.sv.FindMediaPessoaPorMarca(brand)

		if err != nil {
			if errors.Is(err, apperrors.ErrVehicleBrand) {
				response.JSON(w, http.StatusNotFound, map[string]any{
					"message": " Não foram encontrados veículos dessa marca.",
					"data":    nil,
				})
				return
			}
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": "Bad Request: Erro no servidor",
				"data":    nil,
			})
			return
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "capacidade média de pessoas dos veículos da marca",
			"data":    m,
		})

	}
}
