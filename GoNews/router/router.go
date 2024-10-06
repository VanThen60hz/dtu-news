package router

import (
	"gonews/controller"

	"github.com/gofiber/fiber/v2"
)

func NewRouter(router *fiber.App, newsController *controller.NewsController) *fiber.App {
	router.Get("/news/search", newsController.GetNewsByTitle)
	router.Get("/news/:id", newsController.GetNewsById)
	router.Get("/news", newsController.GetNews)
	router.Post("/news", newsController.CreateNews)

	return router
}
