package v1

type LoginRequest struct {
	User struct {
		Email    string `form:"email" json:"email" binding:"exists,email"`
		Password string `form:"password" json:"password" binding:"exists,min=8,max=255"`
	} `json:"user"`
}

// LoginResponse is the response of login request
type LoginResponse struct {
	User         *UserBase `json:"user"`
	AccessToken  string    `json:"accessToken"`
	RefreshToken string    `json:"refreshToken"`
}

// RefreshTokenResponse is the response of refresh token request
type RefreshTokenResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type CreateUserRequest struct {
	User struct {
		Name     string `form:"username" json:"username" binding:"exists,min=4,max=255"`
		Email    string `form:"email" json:"email" binding:"exists,email"`
		Password string `form:"password" json:"password" binding:"exists,min=8,max=255"`
	} `json:"user"`
}

type UserBase struct {
	Name  string `json:"username" gorm:"primaryKey"`
	Email string `json:"email"`
}
