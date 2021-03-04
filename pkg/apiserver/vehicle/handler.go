package vehicle

import "ooneko.github.com/vehicle-insight/pkg/models/vehicle"

type handler struct {
	operator vehicle.Operator
}

func newHandler() *handler {
	return &handler{
		operator: vehicle.NewOperator(),
	}
}

func (h handler) AddVehicle()    {}
func (h handler) DeleteVehicle() {}
func (h handler) UpdateVehicle() {}
