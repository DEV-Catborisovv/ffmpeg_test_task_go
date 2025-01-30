package main

import (
	"ffmpeg_testtask/internal/app/ffmpeg_converter"
	"testing"
	"time"
)

func TestConverting(t *testing.T) {
	converter := ffmpeg_converter.GetConverterInstance()

	startTime := time.Now()
	err := converter.StartConvertation("", 800, 600)
	duration := time.Since(startTime)
	t.Logf("Time for convertation: %v", duration)

	if err != nil {
		t.Error(err)
	}
}
