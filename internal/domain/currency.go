package domain

type Currency string

const (
	BRL Currency = "BRL"
	USD Currency = "USD"
	EUR Currency = "EUR"
)

func (c Currency) String() string {
	return string(c)
}

func (c Currency) IsValid() bool {
	switch c {
	case BRL, USD, EUR:
		return true
	default:
		return false
	}
}

func (c Currency) IsBRL() bool {
	return c == BRL
}

func (c Currency) IsUSD() bool {
	return c == USD
}

func (c Currency) IsEUR() bool {
	return c == EUR
}
