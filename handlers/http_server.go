package handlers

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
)

type Handler interface {
	initializeRoutes() []route
}

type route struct {
	method      string
	path        string
	logResponse bool
	handler     func(w http.ResponseWriter, r *http.Request)
}

func (r *route) getRoute() string {
	return fmt.Sprintf("%s %s", r.method, r.path)
}

type HttpServer struct {
	mux        *http.ServeMux
	server     *http.Server
	address    string
	middleware []Middleware
}

type Middleware interface {
	Handler(route, func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request)
}

func NewHttpServer(
	httpPort string,
	middleware []Middleware,
	handlers ...Handler,
) *HttpServer {
	server := HttpServer{
		mux:        http.NewServeMux(),
		address:    fmt.Sprintf(":%s", httpPort),
		server:     nil,
		middleware: middleware,
	}

	slog.Info(fmt.Sprintf("Listening on %s", server.address))
	server.initializeHandlers(handlers)
	server.server = &http.Server{
		Addr:    server.address,
		Handler: server.mux,
	}
	return &server
}

func (h *HttpServer) ListenAndServe() {
	_ = h.server.ListenAndServe()
}

func (h *HttpServer) Shutdown(ctx context.Context) error {
	err := h.server.Shutdown(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (h *HttpServer) initializeHandlers(handlers []Handler) {
	for _, handler := range handlers {
		routes := handler.initializeRoutes()
		for _, handledRoute := range routes {
			slog.Info(handledRoute.getRoute())
			handlerFunc := handledRoute.handler
			for _, middleware := range h.middleware {
				handlerFunc = middleware.Handler(handledRoute, handlerFunc)
			}

			h.mux.HandleFunc(handledRoute.getRoute(), handlerFunc)
		}
	}
}
