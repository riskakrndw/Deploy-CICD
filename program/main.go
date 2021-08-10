package main

import (
	"fmt"
	"os"

	"github.com/labstack/echo"
)

func main() {
	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = ":80"
	}

	e := echo.New()
	e.GET("/", sayHi)
	e.GET("/:name", greeting)
	e.Start(port)
}

func sayHi(c echo.Context) error {
	return c.HTML(200, "<h1 style='color:#8FBC8F'>Hi</h1><p style='color:#2F4F4F'>I'm Riska</p>")
}

func greeting(c echo.Context) error {
	name := c.Param("name")
	return c.String(200, fmt.Sprintf("Hello %s", name))
}