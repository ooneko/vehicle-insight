package vehicle

type Operator interface {
	AddVehicle() error
}

func NewOperator() Operator {
	return nil
}
