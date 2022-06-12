package service

import (
	"database/sql"
	"prival-api/internal/entity"
	"time"
)

type SecretsService interface {
	CreateSecret(*entity.Secret) error
	ListSecrets(userID int) ([]entity.Secret, error)
	DeleteUsersSecretByID(userId, secretID int) (bool, error)
}

type secretsService struct {
	db *sql.DB
}

func NewSecretsService(db *sql.DB) SecretsService {
	return &secretsService{
		db: db,
	}
}

func (s *secretsService) CreateSecret(data *entity.Secret) error {
	query := `INSERT INTO secrets (user_id, description, created_at) VALUES($1, $2, $3)`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(data.UserID, data.Description, time.Now())
	if err != nil {
		return err
	}

	return nil
}

func (s *secretsService) ListSecrets(userID int) ([]entity.Secret, error) {
	secrets := []entity.Secret{}

	query := `SELECT id, user_id, description, created_at FROM secrets WHERE user_id = $1`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return secrets, err
	}

	rows, err := stmt.Query(userID)
	if err != nil {
		return secrets, err
	}

	for rows.Next() {
		row := entity.Secret{}

		if err = rows.Scan(&row.ID, &row.UserID, &row.Description, &row.CreatedAt); err != nil {
			return secrets, err
		}

		secrets = append(secrets, row)
	}

	return secrets, nil
}

func (s *secretsService) DeleteUsersSecretByID(userID, secretID int) (bool, error) {
	query := `DELETE FROM secrets WHERE id = $1`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return false, err
	}

	result, err := stmt.Exec(secretID)
	if err != nil {
		return false, err
	}

	n, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	return n > 0, nil
}
