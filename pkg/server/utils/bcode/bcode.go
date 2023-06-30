package bcode

import "github.com/gofiber/fiber/v2"

var ErrServer = fiber.NewError(500, "The service has lapsed.")

// ErrForbidden check user perms failure
var ErrForbidden = fiber.NewError(403, "403 Forbidden")

// ErrUnauthorized check user auth failure
var ErrUnauthorized = fiber.NewError(401, "401 Unauthorized")

// ErrNotFound the request resource is not found
var ErrNotFound = fiber.NewError(404, "404 Not Found")

// ErrUpstreamNotFound the proxy upstream is not found
var ErrUpstreamNotFound = fiber.NewError(502, "Upstream not found")
