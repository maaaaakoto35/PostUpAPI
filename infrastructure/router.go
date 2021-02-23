package infrastructure

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/maaaaakoto35/PostUpAPI/interfaces/controllers"
)

// Init this func is initializing server.
func Init() {
	e := echo.New()

	userController := controllers.NewUserController(NewMySQLDb())

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/users", func(c echo.Context) error { return userController.GetUsers(c) })
	e.GET("/users/:id", func(c echo.Context) error { return userController.GetUser(c) })
	e.POST("/setup", func(c echo.Context) error { return userController.CreateUser(c) })
	e.POST("/users/:user_id", func(c echo.Context) error { return userController.UpdateUser(c) })
	e.DELETE("/users/:id", func(c echo.Context) error { return userController.DeleteUser(c) })

	e.Logger.Fatal(e.Start(":8080"))
}
