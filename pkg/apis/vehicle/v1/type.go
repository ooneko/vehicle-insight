package v1

import (
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Vehicle struct {
	metav1.TypeMeta
	metav1.ObjectMeta

	Spec VehicleSpec
}

type VehicleList struct {
	metav1.TypeMeta
	metav1.ObjectMeta

	Items []Vehicle
}

type VehicleSpec struct {
	Model     Model  `json:"model,omitempty"`
	Kilometer int32  `json:"kilometer,omitempty"`
	Parts     []Part `json:"parts,omitempty"`
}
type Part struct {
	Name string `json:"name,omitempty"`
	Life Life   `json:"life,omitempty"`
}

type Life struct {
	Kilometer int32 `json:"life,omitempty"`
	Days      int32 `json:"days,omitempty"`
}

type Model struct {
	Brand   string `json:"brand,omitempty"`
	Version string `json:"version,omitempty"`
}

type MaintenanceRecord struct {
	UUID        string    `json:uuid,omitempty`
	ServiceDate time.Time `json:"service_date,omitempty"`
	CreatedDate time.Time `json:"created_date,omitempty"`
	UseLife     Life      `json:"use_life,omitempty"`
}
