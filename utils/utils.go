package utils

import (
	"net/mail"
)

func IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return false
	}
	return true
}

func IsValidCurrency(currency string) bool {
	switch currency {
	case "USD":
		return true
	case "COP":
		return true
	case "MXN":
		return true
	}
	return false
}

func IsValidTransactionType(transactionType string) bool {
	switch transactionType {
	case "deposit":
		return true
	case "withdraw":
		return true
	case "transfer":
		return true
	}
	return false
}
