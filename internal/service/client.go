package service

import (
	"bank-api/internal/entity"
	"database/sql"
)

func (s Service) InsertClient(client *entity.CreateClientParams) error {
	_, err := s.storage.GetClientByEmail(client.UserEmail)
	if err != nil {
		if err == sql.ErrNoRows {
			err = s.storage.InsertClient(client)
			if err != nil {
				return err
			}
			return nil
		}
		return err
	}
	return entity.EmailInUseError
}

func (s Service) GetClient(email string) (*entity.Client, error) {
	client, err := s.storage.GetClientByEmail(email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, entity.ClientNotFoundError
		}
		return nil, err
	}
	return client, nil
}

func (s Service) GetClients() ([]*entity.Client, error) {
	clients, err := s.storage.GetClients()
	if err != nil {
		return nil, err
	}
	return clients, nil
}
