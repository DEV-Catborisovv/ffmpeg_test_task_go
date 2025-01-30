package app

import api "ffmpeg_testtask/internal/app/API"

type App struct{}

func NewApp() *App {
	return &App{}
}

// Run application method
func (s *App) Init(API_SERVER_PORT int) error {

	// Init API SERVER
	err := api.NewApiServer(API_SERVER_PORT).Init()
	if err != nil {
		return err
	}

	return nil
}
