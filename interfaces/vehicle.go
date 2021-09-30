package interfaces

type Vehicle interface {
	GetBrand() string
}

type Car struct {
	Brand string
}

type Bike struct {
	Brand string
}

func (c Car) GetBrand() string {
	return c.Brand
}

func (b Bike) GetBrand() string {
	return b.Brand
}
