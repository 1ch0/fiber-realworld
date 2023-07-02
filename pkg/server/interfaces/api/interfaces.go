package api

import (
	"github.com/1ch0/fiber-realworld/pkg/server/domain/service"
	jwtware "github.com/gofiber/contrib/jwt"
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

func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   jwtware.SigningKey{Key: []byte(service.SignedKey)},
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "error", "message": "Missing or malformed JWT", "data": nil})
	}
	return c.Status(fiber.StatusUnauthorized).
		JSON(fiber.Map{"status": "error", "message": "Invalid or expired JWT", "data": nil})
}

//func authTokenCheck(req *http.Request, res http.ResponseWriter) bool {
//	// support getting the token from the cookie
//	var tokenValue string
//	tokenHeader := req.Header.Get("Authorization")
//	if tokenHeader != "" {
//		splitted := strings.Split(tokenHeader, " ")
//		if len(splitted) != 2 {
//			bcode.ReturnHTTPError(req, res, bcode.ErrNotAuthorized)
//			return false
//		}
//		tokenValue = splitted[1]
//	}
//	if tokenValue == "" {
//		if strings.HasPrefix(req.URL.Path, "/view") {
//			tokenValue = req.URL.Query().Get("token")
//		}
//		if tokenValue == "" {
//			bcode.ReturnHTTPError(req, res, bcode.ErrNotAuthorized)
//			return false
//		}
//	}
//	token, err := service.ParseToken(tokenValue)
//	if err != nil {
//		bcode.ReturnHTTPError(req, res, err)
//		return false
//	}
//	if token.GrantType != service.GrantTypeAccess {
//		bcode.ReturnHTTPError(req, res, bcode.ErrNotAccessToken)
//		return false
//	}
//	newReq := req.WithContext(context.WithValue(req.Context(), &apis.CtxKeyUser, token.Username))
//	newReq = newReq.WithContext(context.WithValue(newReq.Context(), &apis.CtxKeyToken, tokenValue))
//	*req = *newReq
//	return true
//}
//
//// AuthTokenCheck Parse the token from the request
//func AuthTokenCheck(req *http.Request, res http.ResponseWriter, chain *utils.FilterChain) {
//	if authTokenCheck(req, res) {
//		chain.ProcessFilter(req, res)
//	}
//}
