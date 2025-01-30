package deleteFileById

import (
	"encoding/json"
	"ffmpeg_testtask/internal/app/API/middlewares"
	filestorage "ffmpeg_testtask/internal/app/fileStorage"
	httpstatuses "ffmpeg_testtask/pkg/httpStatuses"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

// @Summary      Delete file
// @Description  Delete file by ID
// @Tags         File
// @Produce      json
// @Param        id   path      string  true  "file ID"
// @Success      200  {object}  response
// @Router       /file/{id} [delete]
func HandleDeleteFileById(w http.ResponseWriter, r *http.Request) {
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

	logger.Message(fmt.Sprintf("[/file/%s -- DELETE]", file_id))

	fileStorage := filestorage.GetFileStorageInstance()
	if !fileStorage.IsExistingFile(file_id) {
		statusWriter.Write(w, httpstatuses.BadRequest)
		return
	}

	err := fileStorage.DeleteFile(file_id)
	if err != nil {
		statusWriter.Write(w, httpstatuses.InternalServerError)
		return
	}

	resp := newResponse(true)
	resp_json, err := json.Marshal(resp)
	if err != nil {
		statusWriter.Write(w, httpstatuses.InternalServerError)
		return
	}

	w.Write(resp_json)
}
