package gee

//相比day1，调整了路由的注册方法，简化代码
//day4,添加路由分组控制功能

import (
	"log"
	"net/http"
	"strings"
	"time"
)

// HandlerFunc 为Handler统一一个别名在框架中使用
type HandlerFunc func(*Context)

// RouterGroup 路由分组也可以在Trie中完成,这里选择独立出来
type RouterGroup struct {
	prefix      string        // 路由组的前缀
	middlewares []HandlerFunc // 中间件
	engine      *Engine       // 该分组下的引擎实例
}

type Engine struct {
	*RouterGroup // 相当于继承自RouterGroup,可以使用它的方法
	router       *Router
	groups       []*RouterGroup // 这个http服务控制下的所有路由组
}

func (g *RouterGroup) Group(prefix string) *RouterGroup {
	engine := g.engine
	newGroup := &RouterGroup{
		prefix: g.prefix + prefix,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

func (g *RouterGroup) Use(middlewares ...HandlerFunc) {
	g.middlewares = append(g.middlewares, middlewares...)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := NewContext(w, req)
	//通过URL前缀简单判断哪些中间件应被添加执行
	for _, group := range e.groups {
		if len(group.prefix) > 0 && strings.HasPrefix(req.URL.Path, group.prefix) {
			c.handlers = append(c.handlers, group.middlewares...)
		}
	}
	//
	e.router.handle(c)
}

func (e *Engine) Run(addr string) {
	http.ListenAndServe(addr, e)
}

// 重写ADD GET POST方法,用RouterGroup来完成

func (g *RouterGroup) ADD(method, pattern string, handler HandlerFunc) {
	pat := g.prefix + pattern //这里要注意加上组的前缀才是完整分组url
	log.Printf("Route %4s - %s", method, pat)
	g.engine.router.addRouter(method, pat, handler)
}

func (g *RouterGroup) GET(pattern string, handler HandlerFunc) {
	g.ADD("GET", pattern, handler)
}

func (g *RouterGroup) POST(pattern string, handler HandlerFunc) {
	g.ADD("POST", pattern, handler)
}

// Logger 提供一个日志中间件
func Logger() HandlerFunc {
	return func(c *Context) {
		tim := time.Now()
		c.Next()
		log.Printf("[%d] %s in %v", c.StatusCode, c.Req.RequestURI, time.Since(tim))
	}
}

func New() *Engine {
	engine := &Engine{router: NewRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}
