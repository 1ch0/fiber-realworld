package bcode

import (
	"errors"

	"github.com/gofiber/fiber/v2"

	"github.com/1ch0/fiber-realworld/pkg/server/infrastructure/mongodb"
)

func MatchDBErr(err error, c *fiber.Ctx) error {
	if errors.Is(err, mongodb.ErrRecordExist) {
		return New(c, ErrRecordExist)
	}

	switch err {
	case mongodb.ErrRecordExist:
		return New(c, ErrRecordExist)
	default:
		return err
	}
}

var (
	ErrRecordExist = NewBcode(10001, "Record already exists")
)
