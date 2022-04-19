package main

func main() {
	e := echo.New()

	e.POST("/login", controller.login)
	e.POST("/auth", controller.auth)

	e.POST("/log/masuk", controller.productMasuk)

	e.Start(":8080")
}