- 为Context添加储存中间件的功能，使得Context不仅可以作用在请求上，也可以作用在响应上。
  ```go
    func (c *Context)A(){
       //part1
       c.Next()
       //part2
	   }
  ```
  - 其中，part1和part2分别代表中间件的执行顺序，part1在中间件执行之前，part2在中间件执行之后。
  - 这样，Context就具备了储存中间件的功能，并且可以作用在请求和响应上。
- 为Context添加中间件的执行顺序的功能，使得中间件可以按照顺序执行。
---
在main.go中测试框架时，突发奇想想试试如果在中间件中使用协程并发处理会怎么样。
```go
func test() gee.HandlerFunc {
	return func(c *gee.Context) {
        var wg sync.WaitGroup
		for i := 0; i < 10; i++ {
            wg.Add(1)
            go func() {
                defer wg.Done()
                for i := 0; i < 100; i++ {
                    c.String(200, "hello world\n")
                }
            }()
        }
        wg.Wait()
    }   
}
```
结果发现，出现了许多没有解决的问题，比如：
1. `c.String(200, "hello world\n")`这行代码会多次调用String，每次调用都会写入和发送状态码，当状态码被重复发送时会发生错误。
2. 在Go的内置库中，Header是以map形式存储的，而map不是并发安全的，如果有多个协程试图同时对map写入数据，就会发生报错
3. 对第二条而言，即使在写入数据时加锁，也无法解决并发问题，原因暂时没有完全理清。
--- 
于是决定暂时放弃并发处理，继续使用单协程先测试框架的功能。
```go
func test() gee.HandlerFunc {
	return func(c *gee.Context) {
		for i := 0; i < 10; i++ {
			c.String(200, "hello!\n")
		}
	}
}
```
结果跟预料一样，框架的功能正常运行。