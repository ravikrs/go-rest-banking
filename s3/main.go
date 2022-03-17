package main

import (
	"github.com/ravikrs/go-rest-banking/s3/app"
	"github.com/ravikrs/go-rest-banking/s3/logger"
)

func main() {
	//log.Println("Starting the application")
	logger.Info("Starting the application")
	app.Start()
}
