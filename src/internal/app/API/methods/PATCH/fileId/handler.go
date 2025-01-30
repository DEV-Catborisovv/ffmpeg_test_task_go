package fileId

import (
	"encoding/json"
	"ffmpeg_testtask/internal/app/API/middlewares"
	"ffmpeg_testtask/internal/app/ffmpeg_converter"
	filestorage "ffmpeg_testtask/internal/app/fileStorage"
	httpstatuses "ffmpeg_testtask/pkg/httpStatuses"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

// @Summary      Start convert file
// @Description  Start convert file to ffmpeg
// @Tags         File
// @Produce      json
// @Param 		input 	body 	request 	true 	"data for ffmpeg"
// @Param        id   path      string  true  "file ID"
// @Success      200  {object}  response
// @Router       /file/{id} [patch]
func HandlePatchFileById(w http.ResponseWriter, r *http.Request) {
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

	// -- Deserialization data from body

	vars := mux.Vars(r)
	file_id := vars["id"]

	body_request, err := io.ReadAll(r.Body)
	if err != nil {
		statusWriter.Write(w, httpstatuses.BadRequest)
		return
	}
	logger.Message(fmt.Sprintf("[/file/%s -- PATCH]", file_id))

	request := newRequest()
	err = json.Unmarshal(body_request, &request)
	if err != nil {
		statusWriter.Write(w, httpstatuses.BadRequest)
		return
	}

	// Work with files
	if !filestorage.GetFileStorageInstance().IsExistingFile(file_id) {
		statusWriter.Write(w, httpstatuses.BadRequest)
		return
	}

	// Convert video
	ffmpegConverterInstance := ffmpeg_converter.GetConverterInstance()
	ffmpegConverterInstance.StartConvertation(file_id, request.Width, request.Height)

	resp := newResponse(true)
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		statusWriter.Write(w, httpstatuses.InternalServerError)
		return
	}
	w.Write(jsonResp)
}
