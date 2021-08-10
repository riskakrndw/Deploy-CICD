package main

import (
	"fmt"
	"os"

	"github.com/labstack/echo"
)

func main() {
	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = ":8080"
	}

	e := echo.New()
	e.GET("/", sayHi)
	e.GET("/:name", greeting)
	e.Start(port)
}

func sayHi(c echo.Context) error {
	return c.String(200, "Hi")
}

func greeting(c echo.Context) error {
	name := c.Param("name")
	return c.String(200, fmt.Sprintf("Hello %s", name))
}