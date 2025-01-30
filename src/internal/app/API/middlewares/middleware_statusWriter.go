package middlewares

import (
	httpstatuses "ffmpeg_testtask/pkg/httpStatuses"
	"fmt"
	"net/http"
)

type StatusWriter struct {
	Middleware
}

func NewStatusWriterMiddleware() *StatusWriter {
	return &StatusWriter{}
}

func (s *StatusWriter) Write(writer http.ResponseWriter, httpStatus httpstatuses.HTTPStatus) {
	writer.WriteHeader(httpStatus.Code)
	writer.Write([]byte(fmt.Sprintf("{ \"error\": \"%s\" }", httpStatus.Message)))
}
