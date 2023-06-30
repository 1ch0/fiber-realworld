package server

import (
	"context"
	"fmt"

	"github.com/1ch0/fiber-realworld/pkg/server/infrastructure/mongodb"

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
	mongo         mongodb.MongoDB
	cfg           config.Config
}

func New(cfg config.Config) Server {
	return &restServer{
		webContainer:  newFiber(),
		beanContainer: container.NewContainer(),
		cfg:           cfg,
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
	// infrastructure
	if err := r.beanContainer.ProvideWithName("RestServer", r); err != nil {
		return fmt.Errorf("fail to provides the RestServer bean to the container: %w", err)
	}

	var mongo mongodb.MongoDB
	if mongo, err = mongodb.New(context.Background(), r.cfg.Mongo); err != nil {
		return fmt.Errorf("fail to init mongodb %w", err)
	}
	r.mongo = mongo
	if err := r.beanContainer.ProvideWithName("mongo", r.mongo); err != nil {
		return fmt.Errorf("fail to provides the mongodb bean to the container: %w", err)
	}

	// domain
	if err := r.beanContainer.Provides(service.InitServiceBean(r.cfg)...); err != nil {
		return fmt.Errorf("fail to provides the service bean to the container: %w", err)
	}

	// interfaces
	if err := r.beanContainer.Provides(api.InitAPIBean()...); err != nil {
		return fmt.Errorf("fail to provides the api bean to the container: %w", err)
	}

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
