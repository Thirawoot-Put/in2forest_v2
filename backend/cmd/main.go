package main

import (
	"thirawoot/in2forest_shop_backend/internal/infras/server"
)

func main() {
	app := server.AppServer()

	app.Start()
}
