package routes

import (
	c "problem2/controller"
	"github.com/labstack/echo"
)
func New() *echo.Echo {
	e := echo.New()
	// user routing
	e.GET("/users", c.GetUsersController)
	e.GET("/users/:id", c.GetUserController)
	e.POST("/users", c.CreateUserController)
	// e.DELETE("/users/:id", c.DeleteUserController)
	e.PUT("/users/:id", c.UpdateUserController)
   
	
	e.GET("/books", c.GetUsersController)
	e.GET("/books/:id", c.GetUserController)
	e.POST("/books", c.CreateUserController)
	// e.DELETE("/books/:id", c.DeleteUserController)
	e.PUT("/books/:id", c.UpdateUserController)
   

	return e

  }