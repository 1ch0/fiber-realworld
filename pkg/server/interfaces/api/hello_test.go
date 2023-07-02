package api

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestHealthzApi(t *testing.T) {
	api := NewHelloApi()
	app := fiber.New()
	api.Register(app)

	req := httptest.NewRequest("GET", "/healthz", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Error(err)
	}
	t.Logf("resp: %#v", resp)
}
