package currency

type CurrencyRepository interface {
	Add(data Currency) (*Currency, error)
	Get(code string) (*Currency, error)
	GetAll() ([]Currency, error)
	Update(code string, data Currency) (*Currency, error)
	Remove(code string) (bool, error)
	RemoveAll() (bool, error)
}
