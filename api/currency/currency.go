package currency

type CurrencyType string

const (
	TypeReal       CurrencyType = "REAL"
	TypeFictitious CurrencyType = "FICTITIOUS"
)

type Currency struct {
	Name string
	Code string
	Type CurrencyType
	Rate float32
}

func (c Currency) IsReal() bool {
	return c.Type == TypeReal
}

func (c Currency) IsFictitious() bool {
	return c.Type == TypeFictitious
}
