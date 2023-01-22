package http

import (
	"bank-api/internal/entity"
)

type Service interface {
	InsertClient(client *entity.CreateClientParams) error
	GetClient(email string) (*entity.Client, error)
	GetClients() ([]*entity.Client, error)

	InsertAccount(account *entity.CreateAccountParams) error
	GetAccount(id int64) (*entity.Account, error)
	GetAccounts(email string) ([]*entity.Account, error)

	InsertTransaction(transaction *entity.CreateTransactionParams) error
	GetTransaction(id int64) (*entity.Transaction, error)
	GetTransactions(id int64) ([]*entity.Transaction, error)
}

type Controller struct {
	service Service
}

func newController(service Service) *Controller {
	return &Controller{service: service}
}
