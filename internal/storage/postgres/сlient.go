package postgres

import "bank-api/internal/entity"

func (p Postgres) InsertClient(client *entity.CreateClientParams) error {
	_, err := p.db.Exec(
		"INSERT INTO client (username, user_email) VALUES ($1, $2);",
		client.Username,
		client.UserEmail,
	)
	if err != nil {
		return err
	}
	return nil
}

func (p Postgres) GetClientByEmail(email string) (*entity.Client, error) {
	client := &entity.Client{}
	err := p.db.Get(client, "SELECT * FROM client WHERE user_email = $1 ORDER BY id;", email)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (p Postgres) GetClients() ([]*entity.Client, error) {
	var clients []*entity.Client
	err := p.db.Select(&clients, "SELECT * FROM client ORDER BY id;")
	if err != nil {
		return nil, err
	}
	return clients, nil
}
