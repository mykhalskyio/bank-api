package entity

type Account struct {
	Id         int64  `db:"id" json:"id"`
	OwnerEmail string `db:"owner_email" json:"owner_email"`
	Balance    int64  `db:"balance" json:"balance"`
	Currency   string `db:"currency" json:"currency"`
	CreatedAt  string `db:"created_at" json:"created_at"`
	UpdatedAt  string `db:"updated_at" json:"updated_at"`
}

type CreateAccountParams struct {
	OwnerEmail string `db:"owner_email" json:"owner_email"`
	Balance    int64  `db:"balance" json:"balance"`
	Currency   string `db:"currency" json:"currency"`
}
