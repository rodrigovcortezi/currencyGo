package redis

import (
	"context"
	"currencyApi/currency"
	"encoding/json"

	"github.com/go-redis/redis/v9"
)

type currencyRepository struct {
	ctx context.Context
	rdb *redis.Client
}

func NewCurrencyRepository(ctx context.Context, rdb *redis.Client) *currencyRepository {
	return &currencyRepository{ctx, rdb}
}

func (repository currencyRepository) Add(c currency.Currency) (*currency.Currency, error) {
	value, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}
	err = repository.rdb.Set(repository.ctx, c.Code, value, 0).Err()
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (repository currencyRepository) Get(code string) (*currency.Currency, error) {
	value, err := repository.rdb.Get(repository.ctx, code).Result()
	if err != nil {
		return nil, err
	}
	currency := &currency.Currency{}
	json.Unmarshal([]byte(value), currency)

	return currency, nil
}

func (repository currencyRepository) GetAll() ([]currency.Currency, error) {
	cmd := redis.NewStringSliceCmd(repository.ctx, "KEYS", "*")
	repository.rdb.Process(repository.ctx, cmd)
	keys, err := cmd.Result()
	if err != nil {
		return nil, err
	}

	currencies := []currency.Currency{}
	for _, k := range keys {
		c, err := repository.Get(k)
		if err != nil {
			return nil, err
		}
		currencies = append(currencies, *c)
	}

	return currencies, nil
}

func (repository currencyRepository) Update(code string, c currency.Currency) (*currency.Currency, error) {
	oldCurrency, err := repository.Get(code)
	if err != nil {
		return nil, err
	}
	currency := *oldCurrency
	currency.Name = c.Name
	currency.Rate = c.Rate

	saved, err := repository.Add(currency)
	if err != nil {
		return nil, err
	}

	return saved, nil
}

func (repository currencyRepository) Remove(code string) (bool, error) {
	count, err := repository.rdb.Del(repository.ctx, code).Result()
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (repository currencyRepository) RemoveAll() (bool, error) {
	all, err := repository.GetAll()
	if err != nil {
		return false, err
	}

	allRemoved := true

	for _, c := range all {
		status, err := repository.Remove(c.Code)
		if err != nil {
			return false, err
		}

		allRemoved = allRemoved && status
	}

	return allRemoved, nil
}
