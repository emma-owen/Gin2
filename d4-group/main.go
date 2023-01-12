package main

import (
	"group/gin2"
	"net/http"
)

func main() {
	r := gin2.New()
	r.GET("/", func(c *gin2.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gin2，Index Page</h1>")
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *gin2.Context) {
			c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
		})

		v1.GET("/hello", func(c *gin2.Context) {
			c.String(http.StatusOK, "hello %s,you're at %s\n", c.Query("name"), c.Path)
		})
	}
	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *gin2.Context) {
			c.String(http.StatusOK, "hello %s,you're at %s\n", c.Param("name"), c.Path)
		})
		v2.GET("/login", func(c *gin2.Context) {
			c.JSON(http.StatusOK, gin2.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})
	}

	r.Run(":9999")
}
