package domain

import (
	"errors"

	"github.com/celsopires1999/estimation/internal/common"
)

type exchange struct {
	rateMap map[currencyYear]float64
}

type currencyYear struct {
	currency Currency
	year     int
}

type rate struct {
	currency Currency
	year     int
	rate     float64
}

func NewExchange(rates []rate) *exchange {
	rateMap := make(map[currencyYear]float64, len(rates))
	for _, rate := range rates {
		rateMap[currencyYear{rate.currency, rate.year}] = rate.rate
	}

	return &exchange{
		rateMap,
	}
}

var ErrInvalidCurrencyYear = errors.New("currency and year combination not found")

func (e *exchange) ConvertToBRL(value float64, currency Currency, year int) (float64, error) {
	rate, ok := e.rateMap[currencyYear{currency, year}]
	if !ok {
		return 0, common.NewDomainValidationError(ErrInvalidCurrencyYear)
	}

	return common.RoundToTwoDecimals(value * rate), nil
}
