package service

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

type connection struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func (c *connection) ToString() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.Host, c.Port, c.User, c.Password, c.DBName,
	)
}

func InitDB() error {
	var err error

	conn := connection{
		Host:     "127.0.0.1",
		Port:     "5432",
		User:     "almighty",
		Password: "secret",
		DBName:   "prival_db",
	}

	DB, err = sql.Open("postgres", conn.ToString())
	if err != nil {
		log.Println(err.Error())
		log.Fatal("Could not connect to database")
		return err
	}

	err = DB.Ping()
	if err != nil {
		log.Println(err.Error())
		log.Fatal("Error pinging the database")
	}

	log.Println("Connection to database established!")
	return nil
}
