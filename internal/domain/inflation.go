package domain

import (
	"errors"

	"github.com/celsopires1999/estimation/internal/common"
)

type inflation struct {
	factors []factor
	table   map[period]float64
}

type factor struct {
	year      int
	inflation float64
}

type period struct {
	start int
	end   int
}

var (
	ErrInvalidPeriod    = errors.New("start greater than end")
	ErrNoStartInflation = errors.New("there is no inflation for start month")
	ErrNoEndInflation   = errors.New("there is no inflation for end month")
)

func NewInflation(factors []factor) *inflation {
	return &inflation{
		factors: factors,
		table:   make(map[period]float64),
	}
}

func (i *inflation) ApplyInflation(value float64, start int, end int) (float64, error) {
	if start > end {
		return 0.0, common.NewDomainValidationError(ErrInvalidPeriod)
	}
	if start == end {
		return common.RoundToTwoDecimals(value), nil
	}
	if acumulatedInflation, ok := i.table[period{start, end}]; ok {
		return common.RoundToTwoDecimals(acumulatedInflation * value), nil
	}

	hasStart := false
	hasEnd := false

	acumulatedInflation := float64(1)
	for _, inflation := range i.factors {
		if inflation.year <= start {
			continue
		}
		if inflation.year == start+1 {
			hasStart = true
		}
		if inflation.year == end {
			hasEnd = true
		}
		acumulatedInflation *= 1 + (inflation.inflation / 100)
		if inflation.year == end {
			break
		}
	}

	if !hasStart {
		return 0.0, common.NewDomainValidationError(ErrNoStartInflation)
	}
	if !hasEnd {
		return 0.0, common.NewDomainValidationError(ErrNoEndInflation)
	}

	acumulatedInflation = common.RoundToFourDecimals(acumulatedInflation)

	p := period{start, end}
	i.table[p] = acumulatedInflation

	return common.RoundToTwoDecimals(acumulatedInflation * value), nil
}
