package util

//Константы для всех поддерживаемых валют
const (
	USD = "USD"
	EUR = "EUR"
	TJS = "TJS"
)

// IsSupportedCurrency возвращает true, если валюта поддерживается.
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, TJS:
		return true
	}
	return false
}
