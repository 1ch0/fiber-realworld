package service

import (
	"github.com/gofiber/fiber/v2"

	"github.com/1ch0/fiber-realworld/pkg/server/infrastructure/mongodb"
)

type UserService interface {
	Login(c *fiber.Ctx) error
}

type userServiceImpl struct {
	Mongo mongodb.MongoDB `inject:"mongo"`
}

func NewUserService() UserService {
	return &userServiceImpl{}
}

func (u *userServiceImpl) Login(c *fiber.Ctx) error {

	return nil
}
