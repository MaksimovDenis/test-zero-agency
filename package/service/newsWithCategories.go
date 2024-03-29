package service

import (
	zeroAgency "test-zero-agency"
	"test-zero-agency/package/repository"
)

type NewsWithCategoriesService struct {
	repo repository.NewsWithCategories
}

func NewNewsWithCategoriesService(repo repository.NewsWithCategories) *NewsWithCategoriesService {
	return &NewsWithCategoriesService{repo: repo}
}

func (n *NewsWithCategoriesService) GetNewsById(newsId int) ([]zeroAgency.NewsWithCategories, error) {
	return n.repo.GetNewsById(newsId)
}

func (n *NewsWithCategoriesService) GetAllNews() ([]zeroAgency.NewsWithCategories, error) {
	return n.repo.GetAllNews()
}

func (n *NewsWithCategoriesService) UpdateNewsById(newsId zeroAgency.NewsWithCategories) error {
	return n.repo.UpdateNewsById(newsId)
}

func (n *NewsWithCategoriesService) DeleteNewsById(newsId int) error {
	return n.repo.DeleteNewsById(newsId)
}
