package api

import (
	"github.com/gofiber/fiber/v2"
)

const versionPrefix = "/api"

type Interface interface {
	Register(app *fiber.App)
}

var registeredAPI []Interface

// RegisterAPI Register API handler
func RegisterAPI(ws ...Interface) {
	registeredAPI = append(registeredAPI, ws...)
}

func GetRegisteredAPI() []Interface {
	return registeredAPI
}

func InitAPIBean() []interface{} {
	RegisterAPI(NewHelloApi(), NewUserApi(), NewArticleApi())

	var beans []interface{}
	for i := range registeredAPI {
		beans = append(beans, registeredAPI[i])
	}
	beans = append(beans)
	return beans
}
