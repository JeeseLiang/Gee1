package main

import (
	"gee"
)

func test() gee.HandlerFunc {
	return func(c *gee.Context) {
		for i := 0; i < 10; i++ {
			c.String(200, "hello!\n")
		}
	}
}
func main() {
	r := gee.New()
	r.Use(gee.Logger())
	r.GET("/", func(c *gee.Context) {
		c.String(200, "<h1>Hello Gee!</h1>")
	})

	v1 := r.Group("/v1")
	v1.Use(test())
	{
		v1.GET("/test/:name", func(c *gee.Context) {
			c.String(200, "hello "+c.Param("name")+"\n")
		})
	}

	r.Run(":9999")

}
