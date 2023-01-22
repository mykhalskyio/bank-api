package service

import (
	"bank-api/internal/entity"
)

type Storage interface {
	InsertClient(client *entity.CreateClientParams) error
	GetClientByEmail(email string) (*entity.Client, error)
	GetClients() ([]*entity.Client, error)

	InsertAccount(account *entity.CreateAccountParams) error
	GetAccountById(id int64) (*entity.Account, error)
	GetClientAccounts(email string) ([]*entity.Account, error)
	GetAccounts() ([]*entity.Account, error)

	CreateTransactionDeposit(transaction *entity.CreateTransactionParams) error
	CreateTransactionWithdraw(transaction *entity.CreateTransactionParams) error
	CreateTransactionTransfer(transaction *entity.CreateTransactionParams) error
	GetTransactionById(id int64) (*entity.Transaction, error)
	GetTransactions() ([]*entity.Transaction, error)
	GetAccountTransactions(accountId int64) ([]*entity.Transaction, error)
}

type Service struct {
	storage Storage
}

func NewService(storage Storage) *Service {
	return &Service{storage: storage}
}
