- 在Gee中，分组控制路由的本质只是将多个具有公共前缀的路由合并成一个路由来处理，因此，分组控制路由的目的就是为了简化路由的配置，提高路由的可读性和可维护性。
- 虽然每个RouterGroup都有自己的Engine实例，但它们其实都是同一个，就是在`r := gee.New()`执行时初始化的那一个
```go
    type RouterGroup struct {
      prefix      string        // 路由组的前缀
      middlewares []HandlerFunc // 中间件
      engine      *Engine       // 该分组下的引擎实例
    }
```
