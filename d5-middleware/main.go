package main

import (
	"d5-middleware/gin2"
	"log"
	"net/http"
	"time"
)

func onlyForV2() gin2.HandlerFunc {
	return func(c *gin2.Context) {
		t := time.Now()
		c.Fail(500, "Internal Server Error")
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func main() {
	r := gin2.New()
	r.Use(gin2.Logger())
	r.GET("/", func(c *gin2.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gin2</h1>")
	})

	v2 := r.Group("/v2")
	v2.Use(onlyForV2())
	{
		v2.GET("/hello/:name", func(c *gin2.Context) {
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
	}

	r.Run(":9999")
}
