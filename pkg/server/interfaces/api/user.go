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
	api := app.Group(versionPrefix + "/users")
	api.Post("", u.CreateUser)
	api.Post("/login", u.login)
}

func (u *user) CreateUser(c *fiber.Ctx) error {
	var req apiv1.CreateUserRequest
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
