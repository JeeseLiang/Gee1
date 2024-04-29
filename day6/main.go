package main

import (
	"gee"
)

func main() {
	r := gee.New()
	r.Use(gee.Logger())
	r.LoadHTMLGlob("templates/*")
	//相当于把模板html中所有用到assets的地方替换成./static的真实路径形式
	r.Static("/assets", "./static")
	r.GET("/", func(c *gee.Context) {
		c.HTML(200, "index.html", nil)
	})
	r.Run(":9999")
}
