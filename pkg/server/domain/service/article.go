package service

import (
	"fmt"
	"github.com/1ch0/fiber-realworld/pkg/server/domain/model"
	"github.com/1ch0/fiber-realworld/pkg/server/infrastructure/mongodb"
	"github.com/1ch0/fiber-realworld/pkg/server/interfaces/api/dto/convert"
	apiv1 "github.com/1ch0/fiber-realworld/pkg/server/interfaces/api/dto/v1"
	"github.com/1ch0/fiber-realworld/pkg/server/utils/bcode"
	"github.com/1ch0/fiber-realworld/pkg/server/utils/log"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ArticleService interface {
	CreateArticle(c *fiber.Ctx, req *apiv1.ArticleRequest) (*apiv1.ArticleResponse, error)
	UpdateArticle(c *fiber.Ctx, req *apiv1.ArticleRequest) (*apiv1.ArticleResponse, error)
	GetArticle(c *fiber.Ctx) (*apiv1.ArticleResponse, error)
	GetArticles(c *fiber.Ctx) (*apiv1.ArticlesResponse, error)
	FavoriteArticle(c *fiber.Ctx) (*apiv1.ArticleResponse, error)
	UnFavoriteArticle(c *fiber.Ctx) (*apiv1.ArticleResponse, error)
}

type articleServiceImpl struct {
	Store                 mongodb.MongoDB       `inject:"mongo"`
	AuthenticationService AuthenticationService `inject:""`
}

func NewArticleService() ArticleService {
	return &articleServiceImpl{}
}

func (a *articleServiceImpl) CreateArticle(c *fiber.Ctx, req *apiv1.ArticleRequest) (*apiv1.ArticleResponse, error) {
	newUUID, err := uuid.NewUUID()
	if err != nil {
		return nil, bcode.ReturnError(c, err)
	}

	article := &model.Article{
		Slug:        newUUID.String(),
		Title:       req.Article.Title,
		Description: req.Article.Description,
		Body:        req.Article.Body,
		TagList:     req.Article.TagList,
	}

	if err := a.Store.Add(c.Context(), article); err != nil {
		return nil, bcode.ReturnError(c, err)
	}

	return &apiv1.ArticleResponse{Article: convert.ArticleModelToAPI(article)}, nil
}

func (a *articleServiceImpl) GetArticle(c *fiber.Ctx) (*apiv1.ArticleResponse, error) {
	slug := c.Params("slug")
	article := &model.Article{Slug: slug}
	if err := a.Store.Get(c.Context(), article); err != nil {
		return nil, bcode.ReturnError(c, err)
	}

	return &apiv1.ArticleResponse{Article: convert.ArticleModelToAPI(article)}, nil
}

func (a *articleServiceImpl) GetArticles(c *fiber.Ctx) (*apiv1.ArticlesResponse, error) {
	entities, err := a.Store.List(c.Context(), &model.Article{}, &mongodb.ListOptions{})
	if err != nil {
		return nil, err
	}
	var articles []apiv1.Article
	for _, entity := range entities {
		article := entity.(*model.Article)
		articles = append(articles, *convert.ArticleModelToAPI(article))
	}

	return &apiv1.ArticlesResponse{Articles: apiv1.Articles{Articles: articles, ArticleCount: len(articles)}}, nil
}

func (a *articleServiceImpl) FavoriteArticle(c *fiber.Ctx) (*apiv1.ArticleResponse, error) {
	user, err := a.AuthenticationService.GetCurrentUser(c)
	if err != nil {
		return nil, bcode.ReturnError(c, err)
	}
	if user == nil {
		return nil, bcode.ErrUnauthorized
	}

	slug := c.Params("slug")
	article := &model.Article{Slug: slug}
	if err := a.Store.Get(c.Context(), article); err != nil {
		return nil, bcode.ReturnError(c, err)
	}

	article.Favorited = true
	article.FavoritesCount++
	if err := a.Store.Put(c.Context(), article); err != nil {
		return nil, bcode.ReturnError(c, err)
	}

	ua := &model.UserArticle{UserName: user.User.Name, Slug: slug}
	if err := a.Store.Add(c.Context(), ua); err != nil {
		return nil, fmt.Errorf("article %s already favorited", article.Slug)
	}

	return &apiv1.ArticleResponse{Article: convert.ArticleModelToAPI(article)}, nil
}

func (a *articleServiceImpl) UnFavoriteArticle(c *fiber.Ctx) (*apiv1.ArticleResponse, error) {
	user, err := a.AuthenticationService.GetCurrentUser(c)
	if err != nil {
		return nil, bcode.ReturnError(c, err)
	}

	slug := c.Params("slug")
	article := &model.Article{Slug: slug}
	if err := a.Store.Get(c.Context(), article); err != nil {
		return nil, bcode.ReturnError(c, err)
	}

	article.Favorited = false
	article.FavoritesCount--
	if err := a.Store.Put(c.Context(), article); err != nil {
		return nil, bcode.ReturnError(c, err)
	}

	ua := &model.UserArticle{UserName: user.User.Name, Slug: slug}
	if err := a.Store.Delete(c.Context(), ua); err != nil {
		return nil, bcode.ReturnError(c, err)
	}

	return &apiv1.ArticleResponse{Article: convert.ArticleModelToAPI(article)}, nil
}

func (a *articleServiceImpl) UpdateArticle(c *fiber.Ctx, req *apiv1.ArticleRequest) (*apiv1.ArticleResponse, error) {
	slug := c.Params("slug")
	article := &model.Article{Slug: slug}
	if err := a.Store.Get(c.Context(), article); err != nil {
		log.Logger.Debugf("Update article %s", slug)
		return nil, bcode.ReturnError(c, err)
	}

	if req.Article.Title != "" {
		article.Title = req.Article.Title
	}
	if req.Article.Description != "" {
		article.Description = req.Article.Description
	}
	if req.Article.Body != "" {
		article.Body = req.Article.Body
	}
	if req.Article.TagList != nil {
		article.TagList = req.Article.TagList
	}
	if err := a.Store.Put(c.Context(), article); err != nil {
		return nil, bcode.ReturnError(c, err)
	}

	return &apiv1.ArticleResponse{Article: convert.ArticleModelToAPI(article)}, nil
}
