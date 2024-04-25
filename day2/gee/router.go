package gee

import (
	"log"
	"net/http"
)

//将路由相关的方法提取出来

type Router struct {
	handlers map[string]HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		handlers: make(map[string]HandlerFunc),
	}
}

func (r *Router) addRouter(method, pattern string, handler HandlerFunc) {
	log.Printf("Route %4s - %4s", method, pattern)
	key := method + "-" + pattern
	r.handlers[key] = handler
}

func (r *Router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if h, ok := r.handlers[key]; ok {
		h(c)
	} else {
		http.Error(c.Writer, "404 NOT FOUND\n", 404)
	}
}
