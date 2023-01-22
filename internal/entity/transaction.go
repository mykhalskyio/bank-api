package entity

type Transaction struct {
	Id            int64  `db:"id" json:"id"`
	Type          string `db:"type" json:"type"`
	FromAccountId int64  `db:"from_account_id" json:"from_account_id"`
	ToAccountId   int64  `db:"to_account_id" json:"to_account_id"`
	Amount        int64  `db:"amount" json:"amount"`
	CreatedAt     string `db:"created_at" json:"created_at"`
}

type CreateTransactionParams struct {
	Type          string `db:"type" json:"type"`
	FromAccountId int64  `db:"from_account_id" json:"from_account_id"`
	ToAccountId   int64  `db:"to_account_id" json:"to_account_id"`
	Amount        int64  `db:"amount" json:"amount"`
}
