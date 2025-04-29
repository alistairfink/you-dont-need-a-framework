package di

import (
	"github.com/alistairfink/you-dont-need-a-framework/handlers"
)

type containerCache struct {
	httpServer *handlers.HttpServer
}

type Container struct {
	cache containerCache
	env   Env
}

func NewContainer(env Env) (*Container, error) {
	return &Container{
		cache: containerCache{},
		env:   env,
	}, nil
}
