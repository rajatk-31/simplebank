package util

const (
	INR = "INR"
	USD = "USD"
	EUR = "EUR"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case INR, USD, EUR:
		return true
	}
	return false
}
