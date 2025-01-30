package ffmpeg_converter

type Video struct {
	Filename          string
	Processing        bool
	ProcessingSuccess bool
}

func newVideo(filename string, processing bool, processingSuccess bool) *Video {
	return &Video{filename, processing, processingSuccess}
}

func (s *Converter) GetVideoInstance(filename string) *Video {
	for i := range s.Videos {
		if s.Videos[i].Filename == filename {
			return &s.Videos[i]
		}
	}
	return nil
}

func (s *Converter) changeVideoStatues(filename string, processing bool, processingSuccess bool) bool {
	for i := range s.Videos {
		if s.Videos[i].Filename == filename {
			s.Videos[i].Processing = processing
			s.Videos[i].ProcessingSuccess = processingSuccess
			return true
		}
	}
	return false
}
