package di

import (
	"github.com/alistairfink/you-dont-need-a-framework/handlers"
)

func (c *Container) HttpServer() (*handlers.HttpServer, error) {
	if c.cache.httpServer == nil {
		exampleHandler := c.exampleHandler()

		c.cache.httpServer = handlers.NewHttpServer(
			c.env.HttpPort,
			[]handlers.Middleware{
				c.loggerMiddleware(),
			},
			exampleHandler,
		)
	}

	return c.cache.httpServer, nil
}

func (c *Container) exampleHandler() *handlers.ExampleHandler {
	return handlers.NewExampleHandler()
}

func (c *Container) loggerMiddleware() handlers.Middleware {
	return handlers.NewLoggerMiddleware()
}
