package service

import (
	zeroAgency "test-zero-agency"
	"test-zero-agency/package/repository"
)

type Authorization interface {
	CreateUser(user zeroAgency.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type NewsWithCategories interface {
	GetNewsById(newsId int) ([]zeroAgency.NewsWithCategories, error)
	GetAllNews() ([]zeroAgency.NewsWithCategories, error)
	UpdateNewsById(newsId zeroAgency.NewsWithCategories) error
	DeleteNewsById(newsId int) error
}

type Service struct {
	Authorization
	NewsWithCategories
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization:      NewAuthService(repos.Authorization),
		NewsWithCategories: NewNewsWithCategoriesService(repos.NewsWithCategories),
	}
}
