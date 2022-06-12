package service

import (
	"database/sql"
	"fmt"
	"log"
	"prival-api/internal/entity"
	"prival-api/internal/helpers"
	"time"
)

type UsersService interface {
	GetUsers() []entity.User
	CreateUser(*entity.User)
	LoginUser(string, string) (*entity.User, error)
}

type usersService struct {
	db *sql.DB
}

func NewUsersService(db *sql.DB) UsersService {
	return &usersService{
		db: db,
	}
}

func (s *usersService) GetUsers() []entity.User {
	return []entity.User{
		{
			ID:        1,
			Username:  "Danny",
			Email:     "danny@example.com",
			Password:  "large-piglet-123",
			CreatedAt: "",
			UpdatedAt: "",
		},
		{
			ID:        2,
			Username:  "Anna",
			Email:     "anna@example.com",
			Password:  "I-love-apples",
			CreatedAt: "",
			UpdatedAt: "",
		},
	}
}

func (s *usersService) CreateUser(u *entity.User) {
	query := `INSERT INTO users (username, email, password, created_at) VALUES ($1, $2, $3, $4)`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		log.Printf("Error preparing statement for query: %s\n", query)
		log.Fatal(err.Error())
	}

	// Encrypt password
	hash, err := helpers.HashString(u.Password)
	if err != nil {
		log.Fatal(err.Error())
	} else {
		u.Password = hash
	}

	_, err = stmt.Exec(&u.Username, &u.Email, &u.Password, time.Now())
	if err != nil {
		log.Printf("Error executing query %s\n", query)
		log.Fatal(err.Error())
	}
}

func (s *usersService) LoginUser(username, password string) (*entity.User, error) {
	query := `SELECT id, username, email, password, created_at FROM users WHERE username = $1`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("could not prepare query to select user")
	}

	row := stmt.QueryRow(username)

	user := &entity.User{}
	err = row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("no user was found with username: %s", username)
	}

	if !helpers.ValidateHash(user.Password, password) {
		return nil, fmt.Errorf("password or username don't match")
	}

	return user, nil
}
