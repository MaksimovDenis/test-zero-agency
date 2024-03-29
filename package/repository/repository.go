package repository

import (
	zeroAgency "test-zero-agency"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user zeroAgency.User) (int, error)
	GetUser(username, password string) (zeroAgency.User, error)
}

type NewsWithCategories interface {
	GetNewsById(newsId int) ([]zeroAgency.NewsWithCategories, error)
	GetAllNews() ([]zeroAgency.NewsWithCategories, error)
	UpdateNewsById(newsId zeroAgency.NewsWithCategories) error
	DeleteNewsById(newsId int) error
}

type Repository struct {
	Authorization
	NewsWithCategories
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization:      NewAuthPostgres(db),
		NewsWithCategories: NewNewsWithCategoriesPostgres(db),
	}
}
