package v1

import "time"

type ArticleRequest struct {
	Article struct {
		Title       string   `json:"title,omitempty"`
		Description string   `json:"description,omitempty"`
		Body        string   `json:"body,omitempty"`
		TagList     []string `json:"tagList,omitempty"`
	} `json:"article"`
}

type ArticleResponse struct {
	Article *Article `json:"article"`
}

type Article struct {
	Slug           string    `json:"slug" gorm:"primaryKey"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	Body           string    `json:"body"`
	TagList        []string  `json:"tagList"`
	Author         string    `json:"author"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	Favorited      bool      `json:"favorited"`
	FavoritesCount int       `json:"favoritesCount"`
}

type ArticlesResponse struct {
	Articles Articles `json:"articles"`
}

type Articles struct {
	ArticleCount int       `json:"articleCount"`
	Articles     []Article `json:"articles"`
}
