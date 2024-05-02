package main

import (
	"github.com/ddiogoo/ddiogoo/auth_server/internal/start"
)

func main() {
	app := start.StartApplication()
	app.Run(":8000")
}
