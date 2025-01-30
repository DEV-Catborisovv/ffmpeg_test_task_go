package ffmpeg_converter

import "sync"

type Converter struct {
	Videos []Video
}

func newConverter() *Converter { return &Converter{} }

var converterInstanceOnce sync.Once
var ConverterInstance *Converter

func GetConverterInstance() *Converter {
	converterInstanceOnce.Do(func() {
		ConverterInstance = newConverter()
	})
	return ConverterInstance
}
