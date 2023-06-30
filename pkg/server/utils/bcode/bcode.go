package bcode

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

var ErrServer = NewBcode(500, "The service has lapsed.")

// ErrForbidden check user perms failure
var ErrForbidden = NewBcode(403, "403 Forbidden")

// ErrUnauthorized check user auth failure
var ErrUnauthorized = NewBcode(401, "401 Unauthorized")

// ErrNotFound the request resource is not found
var ErrNotFound = NewBcode(404, "404 Not Found")

// ErrUpstreamNotFound the proxy upstream is not found
var ErrUpstreamNotFound = NewBcode(502, "Upstream not found")

var ErrInvalidRequest = NewBcode(400, "The request is invalid.")

type Bcode struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

func (b *Bcode) Error() string {
	return fmt.Sprintf("BusinessCode:%d Message:%s", b.Code, b.Message)
}

var bcodeMap map[int32]*Bcode

func NewBcode(businessCode int32, message string) *Bcode {
	if bcodeMap == nil {
		bcodeMap = make(map[int32]*Bcode)
	}
	if _, exit := bcodeMap[businessCode]; exit {
		panic("bcode business code is exist")
	}
	bcode := &Bcode{Code: businessCode, Message: message}
	bcodeMap[businessCode] = bcode
	return bcode
}

func New(c *fiber.Ctx, bcode *Bcode) error {
	return c.Status(fiber.StatusBadRequest).JSON(bcode)
}

func ReturnError(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusBadRequest).JSON(Bcode{Code: 400, Message: err.Error()})
}
