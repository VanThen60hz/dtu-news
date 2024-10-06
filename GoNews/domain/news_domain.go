package domain

import (
	"time"

	"gonews/model"
)

type NewsRepo interface {
	// GetNews(page int, limit int) ([]model.News, int, error)
	GetNews() ([]model.News, error)
	CreateNews(createNews model.News) error
	GetNewsById(id int) (model.News, error)
	GetNewsByTitle(title string) ([]model.News, error)
}

type NewsUseCase interface {
	// GetNews(page int, limit int) ([]model.News, int, error)
	GetNews() ([]model.News, error)
	CreateNews(createNews model.News) error
	GetNewsById(id int) (model.News, error)
	GetNewsByTitle(title string) ([]model.News, time.Duration, error)
}
