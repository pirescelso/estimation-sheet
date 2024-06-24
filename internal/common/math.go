package common

import (
	"github.com/shopspring/decimal"
)

func RoundToTwoDecimals(number float64) float64 {
	v, _ := decimal.NewFromFloat(number).Round(2).Float64()
	return v
}

func RoundToFourDecimals(number float64) float64 {
	v, _ := decimal.NewFromFloat(number).Round(4).Float64()
	return v
}

func IsTwoDecimals(number float64) bool {
	d, _ := decimal.NewFromFloat(number).Round(2).Float64()
	return d == number
}
