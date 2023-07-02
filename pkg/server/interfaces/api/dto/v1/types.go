package v1

import "time"

type LoginRequest struct {
	User struct {
		Email    string `form:"email" json:"email" binding:"exists,email"`
		Password string `form:"password" json:"password" binding:"exists,min=8,max=255"`
	} `json:"user"`
}

// LoginResponse is the response of login request
type LoginResponse struct {
	User LoginUser `json:"user"`
}

type LoginUser struct {
	Email string `json:"email"`
	Name  string `json:"username"`
	Bio   string `json:"bio"`
	Image string `json:"image"`
	Token string `json:"token"`
}

// RefreshTokenResponse is the response of refresh token request
type RefreshTokenResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type UserRequest struct {
	User struct {
		Name     string `form:"username" json:"username" binding:"exists,min=4,max=255"`
		Email    string `form:"email" json:"email" binding:"exists,email"`
		Password string `form:"password" json:"password" binding:"exists,min=8,max=255"`
		Bio      string `json:"bio"`
		Image    string `json:"image"`
	} `json:"user"`
}

type UserBase struct {
	Name  string `json:"username"`
	Email string `json:"email"`
}

type User struct {
	Email    string      `json:"email"`
	Token    string      `json:"token"`
	Username string      `json:"username"`
	Bio      string      `json:"bio"`
	Image    interface{} `json:"image"`
}

type Profile struct {
	Username  string `json:"username"`
	Bio       string `json:"bio"`
	Image     string `json:"image"`
	Following bool   `json:"following"`
}

type Article struct {
	Slug           string    `json:"slug"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	Body           string    `json:"body"`
	TagList        []string  `json:"tagList"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	Favorited      bool      `json:"favorited"`
	FavoritesCount int       `json:"favoritesCount"`
	Author         Author    `json:"author"`
}

type Author struct {
	Username  string `json:"username"`
	Bio       string `json:"bio"`
	Image     string `json:"image"`
	Following bool   `json:"following"`
}

type Comment struct {
	Id        int       `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Body      string    `json:"body"`
	Author    Author    `json:"author"`
}
