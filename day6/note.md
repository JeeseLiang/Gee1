- 这一节主要是实现对模板Template的封装，主要都是在对`Go`内置的模板引擎进行封装
- 实现到这发现在`Gee`框架中，各个模块之间的联系有点过于紧密，导致代码的可能比较混乱，例如
```go
type Context struct {
	// origin objects
	Writer http.ResponseWriter
	Req    *http.Request
	// request info
	Method string
	Path   string
	Params map[string]string
	// response info
	StatusCode int
	// middleware
	handlers []HandlerFunc
	idx      int
	// 让Context也可以访问到engine
	engine *Engine
}

type RouterGroup struct {
    prefix      string        // 路由组的前缀
    middlewares []HandlerFunc // 中间件
    engine      *Engine       // 该分组下的引擎实例
}

type Engine struct {
    router*RouterGroup // 相当于继承自RouterGroup,可以使用它的方法
    router       *Router
    groups       []*RouterGroup // 这个http服务控制下的所有路由组
    //封装渲染引擎
    htmlTemplates *template.Template
    funcMap       template.FuncMap
}
```
- 这个结构之间有过强的耦合性，导致代码如果需要修改时会异常困难
- 写完这一节后，个人认为`Gee`框架的结构设计其实可以优化，比如将`Context`和`Engine`分离，让`Context`只负责处理请求相关的逻辑，而`Engine`只负责路由和渲染相关的逻辑，降低耦合度。