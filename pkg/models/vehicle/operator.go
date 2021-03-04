package vehicle

import (
	"go.uber.org/zap"
	"ooneko.github.com/vehicle-insight/pkg/storage"
	"ooneko.github.com/vehicle-insight/pkg/storage/storagebackend"
	"ooneko.github.com/vehicle-insight/pkg/storage/storagebackend/config"
)

type Operator interface {
	AddVehicle() error
}

func NewOperator(config config.Config, logger *zap.Logger) (Operator, error) {
	storage, err := storagebackend.NewETCD3Storage(config, logger)
	if err != nil {
		return nil, err
	}
	return &operator{
		storage: storage,
	}, nil
}

type operator struct {
	storage storage.Interface
}

func (o *operator) AddVehicle() error {
	return nil
}
