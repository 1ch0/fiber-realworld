package convert

import (
	"github.com/1ch0/fiber-realworld/pkg/server/domain/model"
	apiv1 "github.com/1ch0/fiber-realworld/pkg/server/interfaces/api/dto/v1"
)

func ArticleModelToAPI(article *model.Article) *apiv1.Article {
	return &apiv1.Article{
		Slug:           article.Slug,
		Title:          article.Title,
		Description:    article.Description,
		Body:           article.Body,
		TagList:        article.TagList,
		CreatedAt:      article.CreateTime,
		UpdatedAt:      article.UpdateTime,
		Favorited:      article.Favorited,
		FavoritesCount: article.FavoritesCount,
	}
}
