package auth

import (
	"database/sql"
	"errors"
	"github.com/go-sql-driver/mysql"
	"os"
)

func NewConnectionPool() (*sql.DB, error) {
	dbIP := os.Getenv("DB_IP")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbUsername := os.Getenv("DB_USERNAME")

	db, err := sql.Open("mysql", dbUsername+":"+dbPassword+"@("+dbIP+")/"+dbName+
		"?parseTime=true")
	if err != nil {
		return nil, err
	}
	return db, nil
}

var UserAlreadyExist = errors.New("user already exist")

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

func GetUser(db *sql.DB, username string) (*User, error) {
	row := db.QueryRow("SELECT * FROM users WHERE username=?", username)
	user := &User{}

	err := row.Scan(&user.Username, &user.PwHash)
	if err != nil {
		return nil, err
	}
	return user, nil
}
