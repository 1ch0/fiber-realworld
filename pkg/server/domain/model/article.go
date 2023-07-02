package model

//func init() {
//	RegisterModel(&Article{})
//}

type Article struct {
	BaseModel
	Slug           string   `json:"slug" gorm:"primaryKey"`
	Title          string   `json:"title"`
	Description    string   `json:"description"`
	Body           string   `json:"body"`
	TagList        []string `json:"tagList"`
	Favorited      bool     `json:"favorited"`
	FavoritesCount int      `json:"favoritesCount"`
}

// TableName return custom table name
func (u *Article) TableName() string {
	return tableNamePrefix + "article"
}

// ShortTableName return custom table name
func (u *Article) ShortTableName() string {
	return "article"
}

// PrimaryKey return custom primary key
func (u *Article) PrimaryKey() string {
	return u.Slug
}

// Index return custom index
func (u *Article) Index() map[string]interface{} {
	index := make(map[string]interface{})
	if u.Slug != "" {
		index["slug"] = u.Slug
	}

	return index
}
