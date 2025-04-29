package di

import (
	"github.com/alistairfink/you-dont-need-a-framework/handlers"
)

func (c *Container) HttpServer() (*handlers.HttpServer, error) {
	if c.cache.httpServer == nil {
		exampleHandler, err := c.exampleHandler()
		if err != nil {
			return nil, err
		}

		c.cache.httpServer = handlers.NewHttpServer(
			c.env.HttpPort,
			[]handlers.Middleware{},
			exampleHandler,
		)
	}

	return c.cache.httpServer, nil
}

func (c *Container) exampleHandler() (*handlers.ExampleHandler, error) {
	if c.cache.exampleHandler == nil {
		c.cache.exampleHandler = handlers.NewExampleHandler()
	}

	return c.cache.exampleHandler, nil
}
