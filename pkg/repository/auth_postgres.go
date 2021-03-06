package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	todo "go-application/model"
)

type AuthPostgres struct {
	DB *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{DB: db}
}

func (r *AuthPostgres) CreateUser(user todo.User) (int, error) {
	queryToTable := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id", usersTableName)
	row := r.DB.QueryRow(queryToTable, user.Name, user.Username, user.PasswordHash)
	var id int
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(username, passwordHash string) (todo.User, error) {
	var user todo.User
	queryToTable := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", usersTableName)
	err := r.DB.Get(&user, queryToTable, username, passwordHash)

	return user, err
}
