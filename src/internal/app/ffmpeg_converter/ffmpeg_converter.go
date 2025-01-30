package ffmpeg_converter

import (
	filestorage "ffmpeg_testtask/internal/app/fileStorage"
	"fmt"
	ffmpeg_go "github.com/u2takey/ffmpeg-go"
	"path/filepath"
)

func (s *Converter) StartConvertation(filename string, width int, height int) error {
	go func() {
		filePath := filepath.Join(filestorage.FileStoragePath, filename)
		outputFilePath := filepath.Join(filestorage.FileStoragePath, "converted/", fmt.Sprintf("%s.mp4", filename))

		var isChanged bool
		isChanged = s.changeVideoStatues(filename, true, false)
		if !isChanged {
			newvideo := newVideo(filename, true, false)
			s.Videos = append(s.Videos, *newvideo)
		}

		err := ffmpeg_go.Input(filePath, ffmpeg_go.KwArgs{"ss": "1"}).
			Output(outputFilePath, ffmpeg_go.KwArgs{"s": fmt.Sprintf("%dx%d", width, height)}).
			OverWriteOutput().ErrorToStdOut().Run()

		if err != nil {
			isChanged = s.changeVideoStatues(filename, false, false)
		} else {
			isChanged = s.changeVideoStatues(filename, false, true)
		}
	}()

	return nil
}
