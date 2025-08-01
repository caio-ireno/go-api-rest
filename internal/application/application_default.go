package application

import (
	"app/internal/handler"
	"app/internal/loader"
	"app/internal/repository"
	"app/internal/service"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// ConfigServerChi is a struct that represents the configuration for ServerChi
type ConfigServerChi struct {
	// ServerAddress is the address where the server will be listening
	ServerAddress string
	// LoaderFilePath is the path to the file that contains the vehicles
	LoaderFilePath string
}

// NewServerChi is a function that returns a new instance of ServerChi
func NewServerChi(cfg *ConfigServerChi) *ServerChi {
	// default values
	defaultConfig := &ConfigServerChi{
		ServerAddress: ":8080",
	}
	if cfg != nil {
		if cfg.ServerAddress != "" {
			defaultConfig.ServerAddress = cfg.ServerAddress
		}
		if cfg.LoaderFilePath != "" {
			defaultConfig.LoaderFilePath = cfg.LoaderFilePath
		}
	}

	return &ServerChi{
		serverAddress:  defaultConfig.ServerAddress,
		loaderFilePath: defaultConfig.LoaderFilePath,
	}
}

// ServerChi is a struct that implements the Application interface
type ServerChi struct {
	// serverAddress is the address where the server will be listening
	serverAddress string
	// loaderFilePath is the path to the file that contains the vehicles
	loaderFilePath string
}

// Run is a method that runs the application
func (a *ServerChi) Run() (err error) {
	// dependencies
	// - loader
	ld := loader.NewVehicleJSONFile(a.loaderFilePath)
	db, err := ld.Load()
	if err != nil {
		return
	}
	// - repository
	rp := repository.NewVehicleMap(db)
	// - service
	sv := service.NewVehicleDefault(rp)
	// - handler
	hd := handler.NewVehicleDefault(sv)
	// router
	rt := chi.NewRouter()
	// - middlewares
	rt.Use(middleware.Logger)
	rt.Use(middleware.Recoverer)
	// - endpoints
	rt.Route("/vehicles", func(rt chi.Router) {
		// - GET /vehicles
		rt.Get("/", hd.GetAll())
		// -  GET /GET /vehicles/brand/{brand}/between/{start_year}/{end_year}
		rt.Get("/brand/{brand}/between/{start_year}/{end_year}", hd.GetByMarcaAndYearInterval())
		// -  GET /GET /vehicles/average_speed/brand/{brand}
		rt.Get("/average_speed/brand/{brand}", hd.GetVelocidadeMediaMarca())

		///vehicles/fuel_type/{type}
		rt.Get("/fuel_type/{type}", hd.GetTipoCombustivel())

		// Rota 1 adicionar veiculo
		rt.Post("/", hd.Save())
		// - POST multiplos veiculos
		rt.Post("/batch", hd.SaveMultipleVehicles())

		// - PATCH - vehicles/{id}
		rt.Patch("/{id}", hd.Patch())
		// - PATCH - vehicles/{id}/update_speed
		rt.Patch("/{id}/update_speed", hd.UpdateMaxSpeed())
		// - PATCH /vehicles/{id}/update_fuel

		rt.Patch("/{id}/update_fuel", hd.UpdateFuel())

		// - PATCH - /vehicles/transmission/{type}
		rt.Get("/transmission/{type}", hd.GetTransmissionType())

		// - GET -  /vehicles/average_capacity/brand/{brand}
		// Obter a capacidade média de pessoas por marca
		rt.Get("/average_capacity/brand/{brand}", hd.GetMediaPessoaPorMarca())

		// - DELETE - /vehicles/{id}
		rt.Delete("/{id}", hd.DeleteById())

		// - GET - /vehicles/dimensions?length={min_length}-{max_length}&width={min_width}-{max_width}
		rt.Get("/dimensions", hd.GetByDimensions())

		// - GET /vehicles/weight?min={weight_min}&max={weight_max}
		rt.Get("/weight", hd.GetByPeso())

	})

	rt.Route("/vehiclesc", func(rt chi.Router) {
		// - GET /vehicles by color and years
		rt.Get("/", hd.GetByColorAndYears())

	})

	// run server
	err = http.ListenAndServe(a.serverAddress, rt)
	return
}
