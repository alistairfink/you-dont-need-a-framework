package di

import (
	"github.com/alistairfink/you-dont-need-a-framework/handlers"
)

func (c *Container) HttpServer() (*handlers.HttpServer, error) {
	if c.cache.httpServer == nil {
		exampleHandler := c.exampleHandler()

		c.cache.httpServer = handlers.NewHttpServer(
			c.env.HttpPort,
			// In reverse order since these are wrapped in layers (like an onion)
			// We could change this to top down by changing the processing order in http_server.go
			// When a request is received it will go to loggerMiddleware, responseHeaderMiddleware, then the handler 
			[]handlers.Middleware{
				c.responseHeaderMiddleware(),
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

func (c *Container) responseHeaderMiddleware() handlers.Middleware {
	return handlers.NewResponseHeaderMiddleware()
}
