package repo

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"
	"time"

	"gonews/domain"
	"gonews/model"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type newsRepo struct {
	db  *gorm.DB
	rdb *redis.Client
}

func (n *newsRepo) GetNewsByTitle(title string) ([]model.News, error) {
	var newsList []model.News
	ctx := context.Background()

	result, err := n.rdb.Get(ctx, "news:title:contains:"+title).Result()
	if err != nil && err != redis.Nil {
		return newsList, err
	}

	if len(result) > 0 {
		err := json.Unmarshal([]byte(result), &newsList)
		return newsList, err
	}

	err = n.db.Model(&model.News{}).
		Select("id", "title", "summary", "content").
		Where("title LIKE ?", "%"+title+"%").
		Find(&newsList).Error
	if err != nil {
		return newsList, err
	}

	jsonBytes, err := json.Marshal(newsList)
	if err != nil {
		return newsList, err
	}
	jsonString := string(jsonBytes)

	err = n.rdb.Set(ctx, "news:title:contains:"+title, jsonString, 24*time.Hour).Err()
	if err != nil {
		return newsList, err
	}

	return newsList, nil
}

func (n *newsRepo) GetNewsById(id int) (model.News, error) {
	var news model.News
	ctx := context.Background()

	result, err := n.rdb.Get(ctx, "news:"+strconv.Itoa(id)).Result()
	if err != nil && err != redis.Nil {
		return news, err
	}

	if len(result) > 0 {
		err := json.Unmarshal([]byte(result), &news)
		return news, err
	}

	err = n.db.Model(&model.News{}).Select("id", "title", "summary", "content").Where("id = ?", id).Find(&news).Error
	if err != nil {
		return news, err
	}

	jsonBytes, err := json.Marshal(news)
	if err != nil {
		return news, err
	}
	jsonString := string(jsonBytes)

	err = n.rdb.Set(ctx, "news:"+strconv.Itoa(id), jsonString, 24*time.Hour).Err()
	if err != nil {
		return news, err
	}

	return news, nil
}

func (n *newsRepo) CreateNews(createNews model.News) error {
	if err := n.db.Create(&createNews).Error; err != nil {
		return errors.New("internal server error: Failed to create news")
	}

	return nil
}

// func (n *newsRepo) GetNews(page int, limit int) ([]model.News, int, error) {
// 	var newsList []model.News
// 	var totalCount int64
// 	offset := (page - 1) * limit

// 	if err := n.db.Model(&model.News{}).Count(&totalCount).Error; err != nil {
// 		return newsList, 0, err
// 	}

// 	err := n.db.Model(&model.News{}).
// 		Select("id", "title", "summary", "content").
// 		Limit(limit).
// 		Offset(offset).
// 		Find(&newsList).Error
// 	if err != nil {
// 		return newsList, 0, err
// 	}

// 	return newsList, int(totalCount), nil
// }

func (n *newsRepo) GetNews() ([]model.News, error) {
	var newsList []model.News
	err := n.db.Model(&model.News{}).Select("id", "title", "summary", "content").Find(&newsList).Error
	if err != nil {
		return newsList, err
	}

	return newsList, nil
}

func NewNewsRepo(db *gorm.DB, rdb *redis.Client) domain.NewsRepo {
	return &newsRepo{
		db:  db,
		rdb: rdb,
	}
}
