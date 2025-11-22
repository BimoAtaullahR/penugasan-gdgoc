package main

import (
	"github.com/BimoAtaullahR/penugasan-gdgoc/routes"
)

func main() {
	r := routes.SetupRoutes()
	r.Run(":8080")
}