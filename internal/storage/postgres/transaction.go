package postgres

import "bank-api/internal/entity"

func (p Postgres) CreateTransactionDeposit(transaction *entity.CreateTransactionParams) error {
	tx := p.db.MustBegin()
	_, err := tx.Exec("UPDATE account SET balance = balance + $1 WHERE id = $2;", transaction.Amount, transaction.ToAccountId)
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = tx.Exec(
		"INSERT INTO transaction (type, from_account_id, to_account_id, amount) VALUES ($1, $2, $3, $4);",
		transaction.Type,
		transaction.ToAccountId,
		transaction.ToAccountId,
		transaction.Amount,
	)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (p Postgres) CreateTransactionWithdraw(transaction *entity.CreateTransactionParams) error {
	tx := p.db.MustBegin()
	_, err := tx.Exec("UPDATE account SET balance = balance - $1 WHERE id = $2;", transaction.Amount, transaction.ToAccountId)
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = tx.Exec(
		"INSERT INTO transaction (type, from_account_id, to_account_id, amount) VALUES ($1, $2, $3, $4);",
		transaction.Type,
		transaction.ToAccountId,
		transaction.ToAccountId,
		transaction.Amount,
	)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (p Postgres) CreateTransactionTransfer(transaction *entity.CreateTransactionParams) error {
	tx := p.db.MustBegin()
	_, err := tx.Exec("UPDATE account SET balance = balance - $1 WHERE id = $2;", transaction.Amount, transaction.FromAccountId)
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = tx.Exec("UPDATE account SET balance = balance + $1 WHERE id = $2;", transaction.Amount, transaction.ToAccountId)
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = tx.Exec(
		"INSERT INTO transaction (type, from_account_id, to_account_id, amount) VALUES ($1, $2, $3, $4);",
		transaction.Type,
		transaction.FromAccountId,
		transaction.ToAccountId,
		transaction.Amount,
	)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (p Postgres) GetTransactionById(id int64) (*entity.Transaction, error) {
	transaction := &entity.Transaction{}
	err := p.db.Get(transaction, "SELECT * FROM transaction WHERE id = $1;", id)
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

func (p Postgres) GetTransactions() ([]*entity.Transaction, error) {
	var transactions []*entity.Transaction
	err := p.db.Select(&transactions, "SELECT * FROM transaction;")
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (p Postgres) GetAccountTransactions(accountId int64) ([]*entity.Transaction, error) {
	var transactions []*entity.Transaction
	err := p.db.Select(&transactions, "SELECT * FROM transaction WHERE to_account_id = $1 OR from_account_id = $2;", accountId, accountId)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}
