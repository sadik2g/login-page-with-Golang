package database

import (
	"database/sql"
	"fmt"
	"log"
	"complexloginapp/pkg/models"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() error {
	connStr := "user=your_username dbname=your_database sslmode=disable"
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	fmt.Println("Connected to the database")

	return nil
}

func GetUser(username, password string) (*models.User, error) {
	query := "SELECT id, username, password FROM users WHERE username=$1 AND password=$2"
	row := db.QueryRow(query, username, password)

	var user models.User
	err := row.Scan(&user.ID, &user.Username, &user.Password)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &user, nil
}
