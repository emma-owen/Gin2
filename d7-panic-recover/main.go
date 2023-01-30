package main

import (
	"d7-panic-recovery/gin2"
	"net/http"
)

func main() {
	r := gin2.Default()
	r.GET("/", func(c *gin2.Context) {
		c.String(http.StatusOK, "Hello Geektutu\n")
	})

	r.GET("/panic", func(c *gin2.Context) {
		names := []string{"geektutu"}
		c.String(http.StatusOK, names[100])
	})

	r.Run(":9999")
}
