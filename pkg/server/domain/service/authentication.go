/*
Copyright 2022 The KubeVela Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"github.com/1ch0/fiber-realworld/pkg/server/domain/model"
	"github.com/1ch0/fiber-realworld/pkg/server/infrastructure/mongodb"
	apisv1 "github.com/1ch0/fiber-realworld/pkg/server/interfaces/api/dto/v1"
	"github.com/1ch0/fiber-realworld/pkg/server/utils/bcode"
)

const (
	jwtIssuer = "fiber-issuer"
	// GrantTypeAccess is the grant type for access token
	GrantTypeAccess = "access"
	// GrantTypeRefresh is the grant type for refresh token
	GrantTypeRefresh = "refresh"
)

// signedKey is the signed key of JWT
var signedKey = "fiber-realworld"

// AuthenticationService is the service of authentication
type AuthenticationService interface {
	Login(c *fiber.Ctx, loginReq apisv1.LoginRequest) error
	RefreshToken(c *fiber.Ctx, refreshToken string) (*apisv1.RefreshTokenResponse, error)
	login(c *fiber.Ctx, req apisv1.LoginRequest) (*apisv1.UserBase, error)
}

type authenticationServiceImpl struct {
	UserService UserService     `inject:""`
	Store       mongodb.MongoDB `inject:"mongo"`
}

// NewAuthenticationService new authentication service
func NewAuthenticationService() AuthenticationService {
	return &authenticationServiceImpl{}
}

func (a *authenticationServiceImpl) Login(c *fiber.Ctx, req apisv1.LoginRequest) error {
	userBase, err := a.login(c, req)
	if err != nil {
		return err
	}
	accessToken, err := a.generateJWTToken(userBase.Email, GrantTypeAccess, time.Hour)
	if err != nil {
		return err
	}
	refreshToken, err := a.generateJWTToken(userBase.Email, GrantTypeAccess, time.Hour*24)
	if err != nil {
		return err
	}
	return c.JSON(apisv1.LoginResponse{
		User:         userBase,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}

func (a *authenticationServiceImpl) login(c *fiber.Ctx, req apisv1.LoginRequest) (*apisv1.UserBase, error) {
	if req.User.Email == "" || req.User.Password == "" {
		return nil, bcode.New(c, bcode.ErrInvalidRequest)
	}
	userBase, err := a.UserService.GetUser(c, req.User.Email)
	if err != nil {
		return nil, err
	}

	if err := CompareHashWithPassword(userBase.Password, req.User.Password); err != nil {
		return nil, bcode.ReturnError(c, err)
	}

	return &apisv1.UserBase{Name: userBase.Name, Email: userBase.Email}, nil
}

func (a *authenticationServiceImpl) generateJWTToken(email, grantType string, expireDuration time.Duration) (string, error) {
	claims := model.CustomClaims{
		Email:     email,
		GrantType: grantType,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    jwtIssuer,
			NotBefore: jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expireDuration)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(signedKey))
}

func (a *authenticationServiceImpl) RefreshToken(c *fiber.Ctx, refreshToken string) (*apisv1.RefreshTokenResponse, error) {
	claim, err := ParseToken(refreshToken)
	if err != nil {
		if errors.Is(err, bcode.ErrTokenExpired) {
			return nil, bcode.ErrRefreshTokenExpired
		}
		return nil, err
	}
	if claim.GrantType == GrantTypeRefresh {
		accessToken, err := a.generateJWTToken(claim.Email, GrantTypeRefresh, time.Hour)
		if err != nil {
			return nil, err
		}
		refreshToken, err = a.generateJWTToken(claim.Email, GrantTypeRefresh, time.Hour*24)
		if err != nil {
			return nil, err
		}
		return &apisv1.RefreshTokenResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		}, nil
	}
	return nil, err
}

func ParseToken(tokenString string) (*model.CustomClaims, error) {
	// 解析令牌字符串
	token, err := jwt.ParseWithClaims(tokenString, &model.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(signedKey), nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %v", err)
	}

	if claims, ok := token.Claims.(*model.CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, nil
}

func GeneratePasswordHash(s string) (string, error) {
	if s == "" {
		return "", fmt.Errorf("password is empty")
	}
	hashed, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashed), nil
}

func CompareHashWithPassword(hash, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return errors.New("the password is inconsistent with the user")
	}
	return err
}