package main

import (
	"gin2"
	"net/http"
)

func main() {
	r := gin2.New()
	r.GET("/", func(c *gin2.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gin2</h1>")
	})

	r.GET("/hello", func(c *gin2.Context) {
		c.String(http.StatusOK, "hello %s,you're at %s\n", c.Query("name"), c.Path)
	})

	r.GET("/hello/:name", func(c *gin2.Context) {
		c.String(http.StatusOK, "hello %s,you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *gin2.Context) {
		c.JSON(http.StatusOK, gin2.H{
			"filepath": c.Param("filepath"),
		})
	})

	r.Run(":9999")
}
