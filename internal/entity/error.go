package entity

import "errors"

var (
	InvalidRequestBodyError     = errors.New("invalid request body")
	InvalidEmailError           = errors.New("invalid email address")
	InvalidCurrencyError        = errors.New("invalid currency")
	InvalidIdError              = errors.New("invalid id")
	InvalidTransactionTypeError = errors.New("invalid transaction type")
	ClientNotFoundError         = errors.New("client not found")
	AccountNotFoundError        = errors.New("account not found")
	FromClientNotFoundError     = errors.New("from client not found")
	ToClientNotFoundError       = errors.New("to client not found")
	TransactionNotFoundError    = errors.New("transaction not found")
	DifferentCurrencyError      = errors.New("different currency")
	NotEnoughMoneyError         = errors.New("not enough money")
	EmailInUseError             = errors.New("this email is already in use")
	TooSmallMoneyError          = errors.New("too small an amount of money")
	SameIdError                 = errors.New("same id")
	DatabaseError               = errors.New("problem with the database")
)
