package api

import (
	"github.com/1ch0/fiber-realworld/pkg/server/domain/service"
	apiv1 "github.com/1ch0/fiber-realworld/pkg/server/interfaces/api/dto/v1"
	"github.com/1ch0/fiber-realworld/pkg/server/utils/bcode"
	"github.com/1ch0/fiber-realworld/pkg/server/utils/log"
	"github.com/gofiber/fiber/v2"
)

type article struct {
	Article               service.ArticleService        `inject:""`
	AuthenticationService service.AuthenticationService `inject:""`
}

func NewArticleApi() Interface {
	return &article{}
}

// Register todo: add middleware for authentication
func (a *article) Register(app *fiber.App) {
	api := app.Group(versionPrefix)
	api.Post("/articles", a.CreateArticle)
	api.Get("/articles/:slug", a.GetArticle)
	api.Put("/articles/:slug", a.UpdateArticle)
	api.Get("/articles", a.GetArticles)
	api.Post("/articles/:slug/favorite", a.FavoriteArticle)
	api.Delete("/articles/:slug/favorite", a.UnFavoriteArticle)
}

func (a *article) CreateArticle(c *fiber.Ctx) error {
	var req *apiv1.ArticleRequest
	if err := c.BodyParser(&req); err != nil {
		return err
	}
	resp, err := a.Article.CreateArticle(c, req)
	if err != nil {
		return bcode.ReturnError(c, err)
	}
	return c.Status(fiber.StatusOK).JSON(resp)
}

func (a *article) GetArticle(c *fiber.Ctx) error {
	resp, err := a.Article.GetArticle(c)
	if err != nil {
		return bcode.ReturnError(c, err)
	}
	return c.Status(fiber.StatusOK).JSON(resp)
}

func (a *article) GetArticles(c *fiber.Ctx) error {
	resp, err := a.Article.GetArticles(c)
	if err != nil {
		return bcode.ReturnError(c, err)
	}
	return c.Status(fiber.StatusOK).JSON(resp)
}

func (a *article) FavoriteArticle(c *fiber.Ctx) error {
	resp, err := a.Article.FavoriteArticle(c)
	if err != nil {
		return bcode.ReturnError(c, err)
	}
	return c.Status(fiber.StatusOK).JSON(resp)
}

func (a *article) UnFavoriteArticle(c *fiber.Ctx) error {
	resp, err := a.Article.UnFavoriteArticle(c)
	if err != nil {
		return bcode.ReturnError(c, err)
	}
	return c.Status(fiber.StatusOK).JSON(resp)
}

func (a *article) UpdateArticle(c *fiber.Ctx) error {
	req := &apiv1.ArticleRequest{}
	if err := c.BodyParser(req); err != nil {
		log.Logger.Debugf("API UpdateArticle error: %v", err)
		return bcode.ReturnError(c, err)
	}
	resp, err := a.Article.UpdateArticle(c, req)
	if err != nil {
		return bcode.ReturnError(c, err)
	}
	return c.Status(fiber.StatusOK).JSON(resp)
}
