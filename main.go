package main

import (
	"server/app"
	"server/logger"
)

func main() {
	logger.Info("Starting the application")
	app.Start()
}
