package v1

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

type ArticlesResponse struct {
	Articles Articles `json:"articles"`
}

type Articles struct {
	ArticleCount int       `json:"articleCount"`
	Articles     []Article `json:"articles"`
}
