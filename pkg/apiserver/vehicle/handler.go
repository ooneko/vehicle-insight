package vehicle

import (
	"go.uber.org/zap"
	"ooneko.github.com/vehicle-insight/pkg/models/vehicle"
	"ooneko.github.com/vehicle-insight/pkg/storage/storagebackend/config"
)

type handler struct {
	operator vehicle.Operator
}

func newHandler(config config.Config, logger *zap.Logger) (*handler, error) {
	o, err := vehicle.NewOperator(config, logger)
	if err != nil {
		return nil, err
	}

	return &handler{
		operator: o,
	}, nil
}

func (h handler) AddVehicle()    {}
func (h handler) DeleteVehicle() {}
func (h handler) UpdateVehicle() {}
