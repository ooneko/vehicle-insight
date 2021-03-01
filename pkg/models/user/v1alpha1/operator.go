package v1alpha1

import "ooneko.github.com/vehicle-insight/pkg/simple/client/etcd"

type Interface interface{}

type userOperator struct {
}

func New(client etcd.Client) Interface {

}

func (u *userOperator) List()
