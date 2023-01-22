package entity

type Client struct {
	Id        int64  `db:"id" json:"id"`
	Username  string `db:"username" json:"username"`
	UserEmail string `db:"user_email" json:"user_email"`
	CreatedAt string `db:"created_at" json:"created_at"`
}

type CreateClientParams struct {
	Username  string `db:"username" json:"username"`
	UserEmail string `db:"user_email" json:"user_email"`
}
