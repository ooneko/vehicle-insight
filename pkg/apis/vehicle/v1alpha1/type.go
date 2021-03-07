package v1alpha1

import (
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Vehicle struct {
	metav1.TypeMeta
	metav1.ObjectMeta

	Spec   VehicleSpec   `json:"spec,omitempty"`
	Status VehicleStatus `json:"status,omitempty"`
}

type VehicleList struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Items []Vehicle `json:"items"`
}

type VehicleSpec struct {
	Model              string  `json:"model,omitempty"`
	Vendor             string  `json:"vendor,omitempty"`
	Spares             []Spare `json:"parts,omitempty"`
	MaintenanceRecords []MaintenanceRecord
}

type VehicleStatus struct {
	Kilometer int32 `json:"kilometer,omitempty"`
	Days      int32
}

type Spare struct {
	metav1.TypeMeta
	metav1.ObjectMeta

	Spec   SpareSpec
	Status SpareStatus
}

type SpareList struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Items []Spare `json:"items"`
}

type SpareSpec struct {
	Type string `json:"type,omitempty"`
	Life Life   `json:"life,omitempty"`
}

type SpareStatus struct {
	Life Life `json:"life,omitempty"`
}

type Life struct {
	Kilometer int32 `json:"kilometer,omitempty"`
	Days      int32 `json:"days,omitempty"`
}

type MaintenanceRecord struct {
	metav1.TypeMeta
	metav1.ObjectMeta

	Spec   MaintenanceRecordSpec
	Status MaintenanceRecordStatus `json:"status,omitempty"`
}

type MaintenanceRecordList struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Items []MaintenanceRecord `json:"items"`
}

type MaintenanceRecordSpec struct {
	ServiceDate time.Time `json:"service_date,omitempty"`
	CreatedDate time.Time `json:"created_date,omitempty"`

	KilometerAt   int32   `json:"kilometer_at,omitempty"`
	ReplaceSpares []Spare `json:"replace_spares,omitempty"`
}

type MaintenanceRecordStatus struct {
}
