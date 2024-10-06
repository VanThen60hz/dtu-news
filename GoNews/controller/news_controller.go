package controller

import (
	"net/http"

	"gonews/domain"
	"gonews/model"

	"github.com/gofiber/fiber/v2"
)

type NewsController struct {
	NewsUseCase domain.NewsUseCase
}

func NewNewsController(newsUseCase domain.NewsUseCase) *NewsController {
	return &NewsController{
		NewsUseCase: newsUseCase,
	}
}

// func (n *NewsController) GetNews(ctx *fiber.Ctx) error {
// 	page, err := strconv.Atoi(ctx.Query("page", "1"))
// 	if err != nil {
// 		page = 1
// 	}

// 	limit, err := strconv.Atoi(ctx.Query("limit", "10"))
// 	if err != nil {
// 		limit = 10
// 	}

// 	res, totalPages, err := n.NewsUseCase.GetNews(page, limit)
// 	if err != nil {
// 		response := model.Response{
// 			StatusCode: http.StatusInternalServerError,
// 			Message:    err.Error(),
// 		}

// 		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
// 	}

// 	response := model.Response{
// 		StatusCode: http.StatusOK,
// 		Message:    "Success",
// 		Data:       res,
// 		// Include total pages in the response
// 		Pagination: map[string]interface{}{
// 			"page":        page,
// 			"total_pages": totalPages,
// 		},
// 	}

// 	return ctx.Status(fiber.StatusOK).JSON(response)
// }

func (n *NewsController) GetNews(ctx *fiber.Ctx) error {
	res, err := n.NewsUseCase.GetNews()
	if err != nil {
		response := model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}

		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response := model.Response{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data:       res,
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (n *NewsController) CreateNews(ctx *fiber.Ctx) error {
	var newsRequest model.News
	var response model.Response

	if err := ctx.BodyParser(&newsRequest); err != nil {
		response = model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		}

		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}

	if newsRequest.Title == "" || newsRequest.Content == "" {
		response = model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Title and content cannot be empty",
		}

		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}

	if err := n.NewsUseCase.CreateNews(newsRequest); err != nil {
		response = model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}

		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response = model.Response{
		StatusCode: http.StatusCreated,
		Message:    "News created successfully",
		Data:       newsRequest,
	}

	return ctx.Status(fiber.StatusCreated).JSON(response)
}

func (n *NewsController) GetNewsById(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		response := model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		}

		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}

	res, err := n.NewsUseCase.GetNewsById(id)
	if err != nil {
		response := model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}

		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response := model.Response{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data:       res,
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (n *NewsController) GetNewsByTitle(ctx *fiber.Ctx) error {
	title := ctx.Query("title")

	if title == "" {
		response := model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Title cannot be empty",
		}

		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}

	res, duration, err := n.NewsUseCase.GetNewsByTitle(title)
	if err != nil {
		response := model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}

		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response := model.Response{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data:       res,
		Duration:   duration.String(),
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}
