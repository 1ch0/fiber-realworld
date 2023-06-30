package api

import (
	"github.com/1ch0/fiber-realworld/pkg/server/domain/service"
	"github.com/gofiber/fiber/v2"
)

type user struct {
	UserService service.UserService `inject:""`
}

func NewUserApi() Interface {
	return &user{}
}

func (u *user) Register(app *fiber.App) {
	api := app.Group(versionPrefix + "/users")
	api.Post("/login", u.login)
}

func (u *user) login(c *fiber.Ctx) error {
	if err := u.UserService.Login(c); err != nil {
		return err
	}
	return nil
}
