package currency

import "errors"

type CurrencyService interface {
	Find() ([]Currency, error)
	FindOne(code string) (*Currency, error)
	Create(currency Currency) (*Currency, error)
	Update(code string, currency Currency) (*Currency, error)
	Delete(code string) (bool, error)
}

type service struct {
	repository CurrencyRepository
}

func NewCurrencyService(repository CurrencyRepository) CurrencyService {
	return &service{repository}
}

func (s service) Find() ([]Currency, error) {
	return s.repository.GetAll()
}

func (s service) FindOne(code string) (*Currency, error) {
	currency, err := s.repository.Get(code)
	if err != nil {
		return nil, errors.New("Not found.")
	}

	return currency, nil
}

func (s service) Create(currency Currency) (*Currency, error) {
	saved, err := s.repository.Add(currency)
	if err != nil {
		return nil, err
	}

	return saved, nil
}

func (s service) Update(code string, currency Currency) (*Currency, error) {
	pSaved, err := s.repository.Get(code)
	if err != nil {
		return nil, err
	}

	if pSaved == nil {
		return nil, errors.New("Not found.")
	}

	return s.repository.Update(code, currency)
}

func (s service) Delete(code string) (bool, error) {
	return s.repository.Remove(code)
}
