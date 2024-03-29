package repository

import (
	"fmt"
	logger "test-zero-agency/logs"

	_ "github.com/lib/pq"

	"github.com/jmoiron/sqlx"
)

const (
	userTable               = "users"
	newsTable               = "news"
	newsCategoriesTable     = "newsCategories"
	newsWithCategoriesTable = "newsWithCategories"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	//Check our connection to DB
	err = db.Ping()
	if err != nil {
		logger.Log.Errorf("Failed to connect db (repository)")
		return nil, err
	}
	return db, nil
}
