package postgres

import "bank-api/internal/entity"

func (p Postgres) InsertAccount(account *entity.CreateAccountParams) error {
	_, err := p.db.Exec(
		"INSERT INTO account (owner_email, balance, currency) VALUES ($1, $2, $3);",
		account.OwnerEmail,
		account.Balance,
		account.Currency,
	)
	if err != nil {
		return err
	}
	return nil
}

func (p Postgres) GetAccountById(id int64) (*entity.Account, error) {
	account := &entity.Account{}
	err := p.db.Get(account, "SELECT * FROM account WHERE id = $1;", id)
	if err != nil {
		return nil, err
	}
	return account, nil
}

func (p Postgres) GetClientAccounts(email string) ([]*entity.Account, error) {
	var accounts []*entity.Account
	err := p.db.Select(&accounts, "SELECT * FROM account WHERE owner_email = $1;", email)
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

func (p Postgres) GetAccounts() ([]*entity.Account, error) {
	var accounts []*entity.Account
	err := p.db.Select(&accounts, "SELECT * FROM account;")
	if err != nil {
		return nil, err
	}
	return accounts, nil
}
