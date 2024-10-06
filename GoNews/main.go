package main

import (
	"fmt"
	"log"

	"gonews/config"
	"gonews/controller"
	"gonews/database"
	"gonews/model"
	"gonews/repo"
	"gonews/router"
	"gonews/usecase"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("Hello, World!")

	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load environment variables", err)
	}

	// mysql
	db := database.ConnectionMysqlDB(&loadConfig)
	db.AutoMigrate(model.News{})

	rdb := database.ConnectionRedisDB(&loadConfig)

	startServer(db, rdb)
}

func startServer(db *gorm.DB, rdb *redis.Client) {
	app := fiber.New()
	app.Use(cors.New())

	newsRepo := repo.NewNewsRepo(db, rdb)
	newsUseCase := usecase.NewNewsUseCase(newsRepo)
	newsController := controller.NewNewsController(newsUseCase)

	routes := router.NewRouter(app, newsController)

	err := routes.Listen(":8000")
	if err != nil {
		log.Fatal("Cannot start server", err)
	}
}
