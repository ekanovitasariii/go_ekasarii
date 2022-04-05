package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("saya eka novita sari.")

	e := echo.New()
	e.GET("/", func(ctx echo.Context) error {
		name := ctx.Param("name")
		message := ctx.Param("*")

		data := fmt.Sprint("Hello %s, I love message for you: %s", name, message)
		
		return ctx.String(http.StatusOK, data)
	})
	e.Start(":1000")
}
