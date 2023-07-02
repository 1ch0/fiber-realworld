package service

import (
	"github.com/1ch0/fiber-realworld/pkg/server/utils/bcode"
	"github.com/1ch0/fiber-realworld/pkg/server/utils/log"
	"github.com/gofiber/fiber/v2"
	"time"

	"github.com/1ch0/fiber-realworld/pkg/server/domain/model"

	"github.com/1ch0/fiber-realworld/pkg/server/infrastructure/mongodb"
	apisv1 "github.com/1ch0/fiber-realworld/pkg/server/interfaces/api/dto/v1"
)

type UserService interface {
	GetUser(c *fiber.Ctx, email string) (*model.User, error)
	GetUserByEmail(c *fiber.Ctx, email string) (*apisv1.UserBase, error)
	CreateUser(c *fiber.Ctx, req apisv1.UserRequest) error
	UpdateUser(c *fiber.Ctx, req apisv1.UserRequest) error
}

type userServiceImpl struct {
	Store                 mongodb.MongoDB       `inject:"mongo"`
	AuthenticationService AuthenticationService `inject:""`
}

func NewUserService() UserService {
	return &userServiceImpl{}
}

func (u *userServiceImpl) GetUser(c *fiber.Ctx, email string) (*model.User, error) {
	user := &model.User{
		Email: email,
	}
	if err := u.Store.Get(c.Context(), user); err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userServiceImpl) GetUserByEmail(c *fiber.Ctx, email string) (*apisv1.UserBase, error) {
	user := &model.User{
		Email: email,
	}
	if err := u.Store.Get(c.Context(), user); err != nil {
		return nil, err
	}
	return &apisv1.UserBase{
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (u *userServiceImpl) CreateUser(c *fiber.Ctx, req apisv1.UserRequest) error {
	hash, err := GeneratePasswordHash(req.User.Password)
	if err != nil {
		return bcode.ReturnError(c, err)
	}

	user := &model.User{
		Name:     req.User.Name,
		Email:    req.User.Email,
		Password: hash,
	}

	if err := u.Store.Add(c.Context(), user); err != nil {
		return bcode.MatchDBErr(err, c)
	}

	token, err := u.AuthenticationService.GenerateJWTToken(user.Email, GrantTypeAccess, time.Hour)
	if err != nil {
		log.Logger.Errorf("create user api generate token error: %v", err)
	}
	return c.JSON(apisv1.LoginResponse{
		User: apisv1.LoginUser{
			Email: user.Email,
			Name:  req.User.Name,
			Bio:   req.User.Bio,
			Image: req.User.Image,
			Token: token,
		}})

}

func (u *userServiceImpl) UpdateUser(c *fiber.Ctx, req apisv1.UserRequest) error {
	token, err := u.AuthenticationService.GetToken(c)
	if err != nil {
		return bcode.ReturnError(c, err)
	}
	_, err = ParseToken(token)
	if err != nil {
		return bcode.ReturnError(c, err)
	}

	user := &model.User{
		Email: req.User.Email,
	}
	if err := u.Store.Get(c.Context(), user); err != nil {
		return bcode.MatchDBErr(err, c)
	}

	if err := u.Store.Put(c.Context(), &model.User{
		Email:    user.Email,
		Password: user.Password,
		Name:     req.User.Name,
		Bio:      req.User.Bio,
		Image:    req.User.Image,
	}); err != nil {
		return bcode.MatchDBErr(err, c)
	}

	return c.JSON(apisv1.LoginResponse{
		User: apisv1.LoginUser{
			Email: user.Email,
			Name:  req.User.Name,
			Bio:   req.User.Bio,
			Image: req.User.Image,
			Token: token,
		}})
}
