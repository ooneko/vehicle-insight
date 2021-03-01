package v1alpha1

import (
	vehiclev1alpha1 "ooneko.github.com/vehicle-insight/pkg/apis/vehicle/v1alpha1"
	"ooneko.github.com/vehicle-insight/pkg/simple/client/etcd"
)

type Interface interface {
	List() []*vehiclev1alpha1.Vehicle
}

func New(client etcd.Client) Interface {
	return &Operator{}
}

type Operator struct {
}

func (u *Operator) List() []*vehiclev1alpha1.Vehicle {

}
