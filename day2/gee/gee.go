package gee

//相比day1，调整了路由的注册方法，简化代码

import (
	"net/http"
)

// HandlerFunc 为Handler统一一个别名在框架中使用
type HandlerFunc func(*Context)

// Engine 实现ServeHTTP方法
type Engine struct {
	//路由表，将静态地址和静态路由绑定在Engine实例中，当用户
	//调用get等方法时，从路由表中查询
	//router map[string]HandlerFunc
	//路由已经封装
	router *Router
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := NewContext(w, req)
	e.router.handle(c)
}

// ADD 为GET等方法提供一个统一的添加进路由表的方法
func (e *Engine) ADD(method, pattern string, handler HandlerFunc) {
	e.router.addRouter(method, pattern, handler)
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
	return &Engine{router: NewRouter()}
}
