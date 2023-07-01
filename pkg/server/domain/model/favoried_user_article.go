package model

type UserArticle struct {
	BaseModel
	UserName string `json:"username" `
	Slug     string `json:"slug"`
}

// TableName return custom table name
func (u *UserArticle) TableName() string {
	return tableNamePrefix + "user_article"
}

// ShortTableName return custom table name
func (u *UserArticle) ShortTableName() string {
	return "usr_art"
}

// PrimaryKey return custom primary key
func (u *UserArticle) PrimaryKey() string {
	return u.UserName + "_" + u.Slug
}

// Index return custom index
func (u *UserArticle) Index() map[string]interface{} {
	index := make(map[string]interface{})
	if u.UserName != "" {
		index["username"] = u.UserName
	}
	if u.Slug != "" {
		index["slug"] = u.Slug
	}

	return index
}
