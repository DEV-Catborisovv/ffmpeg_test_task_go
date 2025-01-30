package fileIdInfo

import (
	"encoding/json"
	"ffmpeg_testtask/internal/app/API/middlewares"
	"ffmpeg_testtask/internal/app/ffmpeg_converter"
	filestorage "ffmpeg_testtask/internal/app/fileStorage"
	httpstatuses "ffmpeg_testtask/pkg/httpStatuses"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

// @Summary      Get file status
// @Description  get file status by Id
// @Tags         File
// @Produce      json
// @Param        id   path      string  true  "file ID"
// @Success      200  {object}  response
// @Router       /file/{id} [get]
func HandleGetFileIdInfo(w http.ResponseWriter, r *http.Request) {
	StatusWriter, _ := middlewares.GetMiddleware(middlewares.MiddlewareStatusWriter)
	Logger, _ := middlewares.GetMiddleware(middlewares.MiddlewareLogger)

	// Casting to types
	statusWriter, _ := StatusWriter.(*middlewares.StatusWriter)

	logger, ok := Logger.(*middlewares.LoggerMiddleWare)
	if !ok {
		statusWriter.Write(w, httpstatuses.InternalServerError)
		return
	}

	// -- LOGIC OF METHOD

	vars := mux.Vars(r)
	file_id := vars["id"]

	logger.Message(fmt.Sprintf("[/file/%s -- GET]", file_id))

	if !filestorage.GetFileStorageInstance().IsExistingFile(file_id) {
		statusWriter.Write(w, httpstatuses.BadRequest)
		return
	}

	// Work with data
	converter := ffmpeg_converter.GetConverterInstance()
	videoInstance := converter.GetVideoInstance(file_id)

	resp := &response{}
	if videoInstance == nil {
		resp = newResponse(file_id, file_id, false, false)
	} else {
		resp = newResponse(videoInstance.Filename, videoInstance.Filename, videoInstance.Processing, videoInstance.ProcessingSuccess)
	}

	resp_json, err := json.Marshal(resp)

	if err != nil {
		statusWriter.Write(w, httpstatuses.InternalServerError)
		return
	}

	w.Write(resp_json)
}
