package server

import (
	"context"
	"fmt"

	"github.com/1ch0/fiber-realworld/pkg/server/domain/service"
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/1ch0/fiber-realworld/pkg/server/config"
	"github.com/1ch0/fiber-realworld/pkg/server/interfaces/api"
	"github.com/1ch0/fiber-realworld/pkg/server/utils/container"
)

type Server interface {
	Run(ctx context.Context, errChan chan error) error
}

type restServer struct {
	webContainer  *fiber.App
	beanContainer *container.Container
	cfg           config.Config
}

func New(cfg config.Config) Server {
	return &restServer{
		webContainer: newFiber(),
		cfg:          cfg,
	}
}

func newFiber() *fiber.App {
	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true, // 区分大小写
		StrictRouting: true, // 严格路由
		ServerHeader:  "Fiber Template",
		AppName:       "fiber template",
		JSONEncoder:   sonic.Marshal,
		JSONDecoder:   sonic.Unmarshal,
	})
	app.Use(recover.New())
	app.Use(logger.New(logger.Config{TimeFormat: "2006-01-02 15:04:05", TimeZone: "Asia/Shanghai"}))
	return app
}

func (r *restServer) buildIoCContainer() (err error) {
	r.beanContainer = container.NewContainer()

	if err = r.beanContainer.Populate(); err != nil {
		return fmt.Errorf("fail to provides the event bean to the container: %w", err)
	}
	return nil
}

func (r *restServer) RegisterAPIRoute() {
	api.InitAPIBean()
	for _, handler := range api.GetRegisteredAPI() {
		handler.Register(r.webContainer)
	}
}

func (r *restServer) Run(ctx context.Context, errChan chan error) error {
	// build the Ioc Container
	if err := r.buildIoCContainer(); err != nil {
		return err
	}
	// init database
	if err := service.InitData(ctx); err != nil {
		return fmt.Errorf("fail to init database %w", err)
	}

	r.RegisterAPIRoute()
	return r.webContainer.Listen(r.cfg.Server.BindAddr)
}
