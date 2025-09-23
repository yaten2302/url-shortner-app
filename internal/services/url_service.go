package services

import (
	"context"
	"time"

	"url-shortner-app/internal/models"
	"url-shortner-app/internal/repository"
	"url-shortner-app/pkg/utils"
)

type URLService struct {
	repo *repository.URLRepository
}

func NewURLService(repo *repository.URLRepository) *URLService {
	return &URLService{repo: repo}
}

func (s *URLService) CreateShortURL(ctx context.Context, originalURL string) (*models.ShortURL, error) {
	shortCode, err := utils.GenerateShortCode(len(originalURL))
	if err != nil {
		return nil, err
	}
	now := time.Now()

	shortURL := models.ShortURL{
		URL:         originalURL,
		ShortCode:   shortCode,
		CreatedAt:   now,
		UpdatedAt:   now,
		AccessCount: 0,
	}

	if err := s.repo.Insert(ctx, shortURL); err != nil {
		return nil, err
	}

	return &shortURL, nil
}

func (s *URLService) GetOriginalURL(ctx context.Context, shortCode string) (string, error) {
	url, err := s.repo.FindByShortCode(ctx, shortCode)
	if err != nil {
		return "", err
	}
	if url == nil {
		return "", nil
	}

	// Update access count
	url.AccessCount++
	url.UpdatedAt = time.Now()
	if err := s.repo.Update(ctx, *url); err != nil {
		return "", err
	}

	return url.URL, nil
}

func (s *URLService) DeleteURL(ctx context.Context, shortCode string) error {
	return s.repo.DeleteByShortCode(ctx, shortCode)
}
