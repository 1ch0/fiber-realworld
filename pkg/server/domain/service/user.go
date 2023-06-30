package service

import (
	"github.com/1ch0/fiber-realworld/pkg/server/utils/bcode"
	"github.com/gofiber/fiber/v2"

	"github.com/1ch0/fiber-realworld/pkg/server/domain/model"

	"github.com/1ch0/fiber-realworld/pkg/server/infrastructure/mongodb"
	apisv1 "github.com/1ch0/fiber-realworld/pkg/server/interfaces/api/dto/v1"
)

type UserService interface {
	GetUser(c *fiber.Ctx, email string) (*model.User, error)
	GetUserByEmail(c *fiber.Ctx, email string) (*apisv1.UserBase, error)
	CreateUser(c *fiber.Ctx, req apisv1.CreateUserRequest) error
}

type userServiceImpl struct {
	Store mongodb.MongoDB `inject:"mongo"`
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

func (u *userServiceImpl) CreateUser(c *fiber.Ctx, req apisv1.CreateUserRequest) error {
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

	return c.JSON(apisv1.UserBase{
		Name:  user.Name,
		Email: user.Email,
	})
}
