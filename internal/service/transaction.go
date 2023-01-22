package service

import (
	"bank-api/internal/entity"
	"database/sql"
)

func (s Service) InsertTransaction(transaction *entity.CreateTransactionParams) error {
	var (
		accountFrom *entity.Account
		accountTo   *entity.Account
		err         error
	)
	accountTo, err = s.storage.GetAccountById(transaction.ToAccountId)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.ToClientNotFoundError
		}
		return err
	}
	switch transaction.Type {
	case "deposit":
		err = s.storage.CreateTransactionDeposit(transaction)
	case "withdraw":
		if (accountTo.Balance - transaction.Amount) < 0 {
			return entity.NotEnoughMoneyError
		}
		err = s.storage.CreateTransactionWithdraw(transaction)
	case "transfer":
		accountFrom, err = s.storage.GetAccountById(transaction.FromAccountId)
		if err != nil {
			if err == sql.ErrNoRows {
				return entity.FromClientNotFoundError
			}
			return err
		}
		if accountFrom.Currency != accountTo.Currency {
			return entity.DifferentCurrencyError
		}
		if (accountFrom.Balance - transaction.Amount) < 0 {
			return entity.NotEnoughMoneyError
		}
		err = s.storage.CreateTransactionTransfer(transaction)
	}
	if err != nil {
		return err
	}
	return nil
}

func (s Service) GetTransactions(id int64) ([]*entity.Transaction, error) {
	var (
		transactions []*entity.Transaction
		err          error
	)
	if id != 0 {
		_, err = s.storage.GetAccountById(id)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, entity.AccountNotFoundError
			}
			return nil, err
		}
		transactions, err = s.storage.GetAccountTransactions(id)
		if err != nil {
			return nil, err
		}
	} else {
		transactions, err = s.storage.GetTransactions()
		if err != nil {
			return nil, err
		}
	}
	return transactions, nil
}

func (s Service) GetTransaction(id int64) (*entity.Transaction, error) {
	transaction, err := s.storage.GetTransactionById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, entity.TransactionNotFoundError
		}
		return nil, err
	}
	return transaction, nil
}
