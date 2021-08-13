package main

import (
	"fmt"
	"os"
	"net/http"

	"github.com/labstack/echo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type User struct {
	gorm.Model
	Name string
}

func main() {
	connectionString := os.Getenv("CONNECTION_STRING")
	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB.AutoMigrate(&User{})
	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = ":80"
	}

	e := echo.New()
	e.GET("/", sayHi)
	e.GET("/greeting/:name", greeting)
	e.GET("/user", getUser)
	e.Start(port)
}

func sayHi(c echo.Context) error {
	return c.HTML(200, "<h1 style='color:#8FBC8F'>Hi</h1><p style='color:#2F4F4F'>I'm Riska</p>")
}

func greeting(c echo.Context) error {
	name := c.Param("name")
	return c.String(200, fmt.Sprintf("Hello %s", name))
}

func getUser(c echo.Context) error {
	var users []User
	if err := DB.Find(&users).Error; err != nil {
		fmt.Println(err)
	}
	return c.JSON(http.StatusOK, users)
}