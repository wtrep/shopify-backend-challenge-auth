package auth

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
)

// Thrown when a user already exist in the DB with the same username
var UserAlreadyExist = errors.New("user already exist")

// Return a new DB Connection Pool
func NewConnectionPool() (*sql.DB, error) {
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USERNAME")
	dbIP, ok := os.LookupEnv("DB_IP")
	if !ok {
		dbIP = "127.0.0.1"
	}

	var dbURI string
	dbURI = fmt.Sprintf("%s:%s@(%s)/%s?parseTime=true", dbUser, dbPassword, dbIP, dbName)

	db, err := sql.Open("mysql", dbURI)
	if err != nil {
		return nil, err
	}

	err = createImageTableIfNotExist(db)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// Create the users table if it doesn't exist
func createImageTableIfNotExist(db *sql.DB) error {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS users (username varchar(32) not null primary key," +
		"pwHash binary(60) not null)")
	if err != nil {
		return err
	}
	return nil
}

// Create a new user in the database
func CreateUser(db *sql.DB, user User) error {
	_, err := db.Exec("INSERT INTO users (username, pwHash) VALUEs (?, ?)", user.Username, user.PwHash)
	if err, ok := err.(*mysql.MySQLError); ok {
		if err.Number == 1062 {
			return UserAlreadyExist
		}
		return errors.New(err.Error())
	}
	return nil
}

// Get an existing user from the database
func GetUser(db *sql.DB, username string) (*User, error) {
	row := db.QueryRow("SELECT * FROM users WHERE username=?", username)
	user := &User{}

	err := row.Scan(&user.Username, &user.PwHash)
	if err != nil {
		return nil, err
	}
	return user, nil
}
