package model

//func init() {
//	RegisterModel(&Article{})
//}

type Tag struct {
	BaseModel
	Name string   `json:"name" gorm:"primaryKey"`
	Slug []string `json:"slug"`
}

// TableName return custom table name
func (u *Tag) TableName() string {
	return tableNamePrefix + "tag"
}

// ShortTableName return custom table name
func (u *Tag) ShortTableName() string {
	return "tag"
}

// PrimaryKey return custom primary key
func (u *Tag) PrimaryKey() string {
	return u.Name
}

// Index return custom index
func (u *Tag) Index() map[string]interface{} {
	index := make(map[string]interface{})
	if u.Name != "" {
		index["name"] = u.Name
	}

	return index
}
