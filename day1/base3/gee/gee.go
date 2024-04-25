package gee

import (
	"fmt"
	"net/http"
)

// HandlerFunc 为Handler统一一个别名在框架中使用
type HandlerFunc func(http.ResponseWriter, *http.Request)

// Engine 实现ServeHTTP方法
type Engine struct {
	//路由表，将静态地址和静态路由绑定在Engine实例中，当用户
	//调用get等方法时，从路由表中查询
	router map[string]HandlerFunc
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := e.router[key]; ok {
		handler(w, req)
	} else {
		w.WriteHeader(404) //设置404返回码状态
		fmt.Fprintf(w, "404 NOT FOUND : %s\n", req.URL)
	}
}

// ADD 为GET等方法提供一个统一的添加进路由表的方法
func (e *Engine) ADD(method, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	e.router[key] = handler
}

func (e *Engine) GET(pattern string, handler HandlerFunc) {
	e.ADD("GET", pattern, handler)
}

func (e *Engine) POST(pattern string, handler HandlerFunc) {
	e.ADD("POST", pattern, handler)
}

func (e *Engine) Run(addr string) {
	http.ListenAndServe(addr, e)
}

func New() *Engine {
	return &Engine{map[string]HandlerFunc{}}
}
