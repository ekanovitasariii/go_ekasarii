package main

import (
	"problem2/config"
	"problem2/routes"
)

func main(){
	config.InitDB()
	e := routes.New()
	e.Start(":8000")
}