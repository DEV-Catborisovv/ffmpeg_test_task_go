package file

import (
	"encoding/json"
	"ffmpeg_testtask/internal/app/API/middlewares"
	filestorage "ffmpeg_testtask/internal/app/fileStorage"
	httpstatuses "ffmpeg_testtask/pkg/httpStatuses"
	"io"
	"net/http"
)

// @Summary      Import file
// @Description  Add file to system
// @Tags         File
// @Produce      json
// @Param file formData file true "Upload file"
// @Success      200  {object}  responseJson
// @Router       /file/ [post]
func HandlePostFileHTTPMethod(w http.ResponseWriter, r *http.Request) {

	// Here i'm disable getting error because if i can't get middleware for error-writer i really have problems :)
	StatusWriter, _ := middlewares.GetMiddleware(middlewares.MiddlewareStatusWriter)
	Logger, _ := middlewares.GetMiddleware(middlewares.MiddlewareLogger)
	UuidGenerator, _ := middlewares.GetMiddleware(middlewares.MiddlewareUuidGenerator)

	// casting to types
	statusWriter, _ := StatusWriter.(*middlewares.StatusWriter)

	logger, ok := Logger.(*middlewares.LoggerMiddleWare)
	if !ok {
		statusWriter.Write(w, httpstatuses.InternalServerError)
		return
	}

	uuidGenerator, ok := UuidGenerator.(*middlewares.UuidGenerator)
	if !ok {
		statusWriter.Write(w, httpstatuses.InternalServerError)
		return
	}

	// -- LOGIC OF METHOD

	// -- Deserialization data from body

	body_file, err := io.ReadAll(r.Body)
	if err != nil || len(body_file) == 0 {
		statusWriter.Write(w, httpstatuses.BadRequest)
		return
	}

	logger.Message("[/file -- POST] Added new file")

	// -- Uuid generate

	err, uuidCode := uuidGenerator.Generate()
	if err != nil {
		statusWriter.Write(w, httpstatuses.InternalServerError)
		return
	}

	// -- Save file in the storage
	fsInstance := filestorage.GetFileStorageInstance()
	err = fsInstance.SaveFile(body_file, uuidCode)
	if err != nil {
		statusWriter.Write(w, httpstatuses.InternalServerError)
		return
	}

	// -- Make a response

	respStruct := newResponseJson(uuidCode)
	response, err := json.Marshal(&respStruct)
	if err != nil {
		statusWriter.Write(w, httpstatuses.InternalServerError)
		return
	}

	w.Write(response)
}
