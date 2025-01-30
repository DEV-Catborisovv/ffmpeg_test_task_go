package main

import (
	"ffmpeg_testtask/internal/app"
	"ffmpeg_testtask/internal/app/configs"
	"log"
)

// @title           Test task for Go developer
// @version         1.0
// @description     This solution for this test task: https://github.com/sdobrimutrom/python_test

// @host      localhost:8797
// @BasePath  /

func main() {
	application_config := configs.GetApplicationConfigInstance()

	application := app.NewApp()
	err := application.Init(application_config.API_SERVER_PORT)

	if err != nil {
		log.Fatalf("Can't run application:\n\n%s", err)
	}
}
