package usecase

import (
	"errors"
	"time"

	"gonews/domain"
	"gonews/model"
)

type NewsUseCase struct {
	newsRepo domain.NewsRepo
}

func (n *NewsUseCase) GetNewsByTitle(title string) ([]model.News, time.Duration, error) {
	startTime := time.Now()
	res, err := n.newsRepo.GetNewsByTitle(title)
	duration := time.Since(startTime)

	if err != nil {
		return nil, duration, errors.New("internal server error: " + err.Error())
	}

	return res, duration, nil
}

func (n *NewsUseCase) GetNewsById(id int) (model.News, error) {
	res, err := n.newsRepo.GetNewsById(id)
	if err != nil {
		return model.News{}, errors.New("internal server error:" + err.Error())
	}

	return res, nil
}

func (n *NewsUseCase) CreateNews(createNews model.News) error {
	if err := n.newsRepo.CreateNews(createNews); err != nil {
		return err
	}

	return nil
}

// func (n *NewsUseCase) GetNews(page int, limit int) ([]model.News, int, error) {
// 	res, totalCount, err := n.newsRepo.GetNews(page, limit)
// 	if err != nil {
// 		return nil, 0, errors.New("internal server error: " + err.Error())
// 	}

// 	totalPages := (totalCount + limit - 1) / limit

// 	return res, totalPages, nil
// }

func (n *NewsUseCase) GetNews() ([]model.News, error) {
	res, err := n.newsRepo.GetNews()
	if err != nil {
		return nil, errors.New("internal server error:" + err.Error())
	}

	return res, nil
}

func NewNewsUseCase(newsRepo domain.NewsRepo) domain.NewsUseCase {
	return &NewsUseCase{
		newsRepo: newsRepo,
	}
}
