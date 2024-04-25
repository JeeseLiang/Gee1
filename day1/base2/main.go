package main

import (
	"fmt"
	"net/http"
)

/*
从base1中的http.ListenAndServe(":8888", nil)中发现
第二个参数是Handler接口
只要实现了Handler接口，就可以将http请求交给该实例处理
*/

type Engine struct{}

/*
定义一个Engine结构体，并实现ServeHTTP方法，这样就可以将Engine实例作为Handler处理http请求
通过这种方法，我们可以统一将http请求的入口交由我们自定义，走出实现Web框架的第一步
*/
func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)

	case "/hello":
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] : %q\n", k, v)
		}
	default:
		fmt.Fprintf(w, "404 NOT FOUND : %s\n", req.URL)
	}
}

func main() {
	e := new(Engine)
	http.ListenAndServe(":8888", e)
}
