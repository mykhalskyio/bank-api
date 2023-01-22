package service

import (
	"bank-api/internal/entity"
	"database/sql"
)

func (s Service) InsertAccount(account *entity.CreateAccountParams) error {
	_, err := s.storage.GetClientByEmail(account.OwnerEmail)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.ClientNotFoundError
		}
		return err
	}
	err = s.storage.InsertAccount(account)
	if err != nil {
		return err
	}
	return nil
}

func (s Service) GetAccount(id int64) (*entity.Account, error) {
	account, err := s.storage.GetAccountById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, entity.AccountNotFoundError
		}
		return nil, err
	}
	return account, nil
}

func (s Service) GetAccounts(email string) ([]*entity.Account, error) {
	var (
		accounts []*entity.Account
		err      error
	)
	if email == "" {
		accounts, err = s.storage.GetAccounts()
		if err != nil {
			return nil, err
		}
	} else {
		accounts, err = s.storage.GetClientAccounts(email)
		if err != nil {
			return nil, err
		}
	}
	return accounts, nil
}
