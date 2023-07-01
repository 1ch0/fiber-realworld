package model

import "github.com/golang-jwt/jwt/v5"

func init() {
	RegisterModel(&User{})
}

// DefaultAdminUserAlias default admin user alias
const DefaultAdminUserAlias = "Administrator"

// RoleAdmin admin role
const RoleAdmin = "admin"

// User is the model of user
type User struct {
	BaseModel
	Name     string `json:"username" gorm:"primaryKey"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
	Bio      string `json:"bio"`
	Image    string `json:"image"`
}

// TableName return custom table name
func (u *User) TableName() string {
	return tableNamePrefix + "user"
}

// ShortTableName return custom table name
func (u *User) ShortTableName() string {
	return "usr"
}

// PrimaryKey return custom primary key
func (u *User) PrimaryKey() string {
	return u.Email
}

// Index return custom index
func (u *User) Index() map[string]interface{} {
	index := make(map[string]interface{})
	if u.Name != "" {
		index["name"] = u.Name
	}
	if u.Email != "" {
		index["email"] = u.Email
	}

	return index
}

// CustomClaims is the custom claims
type CustomClaims struct {
	Email     string `json:"email"`
	GrantType string `json:"grantType"`
	jwt.RegisteredClaims
}
