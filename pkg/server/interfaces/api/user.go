package api

import (
	"github.com/1ch0/fiber-realworld/pkg/server/domain/service"
	apiv1 "github.com/1ch0/fiber-realworld/pkg/server/interfaces/api/dto/v1"
	"github.com/gofiber/fiber/v2"
)

type user struct {
	UserService           service.UserService           `inject:""`
	AuthenticationService service.AuthenticationService `inject:""`
}

func NewUserApi() Interface {
	return &user{}
}

func (u *user) Register(app *fiber.App) {
	api := app.Group(versionPrefix)
	api.Post("users", u.CreateUser)
	api.Post("/users/login", u.login)
	api.Get("/user", u.CurrentUser)
	api.Put("/user", u.UpdateUser)
}

func (u *user) CreateUser(c *fiber.Ctx) error {
	var req apiv1.UserRequest
	if err := c.BodyParser(&req); err != nil {
		return err
	}
	return u.UserService.CreateUser(c, req)
}

func (u *user) login(c *fiber.Ctx) error {
	var req apiv1.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return err
	}
	return u.AuthenticationService.Login(c, req)
}

func (u *user) CurrentUser(c *fiber.Ctx) error {
	return u.AuthenticationService.CurrentUser(c)
}

func (u *user) UpdateUser(c *fiber.Ctx) error {
	var req apiv1.UserRequest
	if err := c.BodyParser(&req); err != nil {
		return err
	}
	return u.UserService.UpdateUser(c, req)
}
