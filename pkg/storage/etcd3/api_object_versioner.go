package etcd3

import (
	"strconv"

	metav1 "ooneko.github.com/vehicle-insight/pkg/apis/meta/v1"
)

type APIObjectVersioner struct{}

func (a APIObjectVersioner) UpdateObject(obj metav1.Object, resourceVersion uint64) error {
	versionString := ""
	if resourceVersion != 0 {
		versionString = strconv.FormatUint(resourceVersion, 10)
	}
	obj.SetResourceVersion(versionString)
	return nil
}

func (a APIObjectVersioner) ObjectResourceVersion(obj metav1.Object) (uint64, error) {
	version := obj.GetResourceVersion()
	if len(version) == 0 {
		return 0, nil
	}
	return strconv.ParseUint(version, 10, 64)
}

func (a APIObjectVersioner) PrepareObjectForStorage(obj metav1.Object) error {
	obj.SetResourceVersion("")
	return nil
}
