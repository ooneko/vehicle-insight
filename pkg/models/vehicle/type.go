package vehicle

import "time"

type Vehicle struct {
	Name      string `json:"name,omitempty"`
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

type UUID string

type MaintenanceRecord struct {
	UUID        UUID
	ServiceDate time.Time `json:"service_date,omitempty"`
	CreatedDate time.Time `json:"created_date,omitempty"`
	UseLife     Life      `json:"use_life,omitempty"`
}
