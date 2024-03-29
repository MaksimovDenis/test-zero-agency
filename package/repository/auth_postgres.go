package repository

import (
	"fmt"
	zeroAgency "test-zero-agency"

	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (a *AuthPostgres) CreateUser(user zeroAgency.User) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s(username, password_hash) values ($1, $2) RETURNING id", userTable)

	row := a.db.QueryRow(query, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (a *AuthPostgres) GetUser(username, password string) (zeroAgency.User, error) {
	var user zeroAgency.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", userTable)
	err := a.db.Get(&user, query, username, password)

	return user, err

}
