package api

import (
	"fmt"
	"net/http"

	_ "ffmpeg_testtask/docs"
	"ffmpeg_testtask/internal/app/API/methods/DELETE/deleteFileById"
	"ffmpeg_testtask/internal/app/API/methods/GET/fileIdInfo"
	"ffmpeg_testtask/internal/app/API/methods/PATCH/fileId"
	"ffmpeg_testtask/internal/app/API/methods/POST/file"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

type ApiServer struct {
	serverPort int
	router     *mux.Router
}

func NewApiServer(port int) *ApiServer {
	return &ApiServer{serverPort: port}
}

// Setup routes
func (s *ApiServer) setupRoutes() {
	s.router.HandleFunc("/file/", file.HandlePostFileHTTPMethod).Methods("POST")
	s.router.HandleFunc("/file/{id}", fileId.HandlePatchFileById).Methods("PATCH")
	s.router.HandleFunc("/file/{id}", fileIdInfo.HandleGetFileIdInfo).Methods("GET")
	s.router.HandleFunc("/file/{id}", deleteFileById.HandleDeleteFileById).Methods("DELETE")
}

// Init server method
func (s *ApiServer) Init() error {
	s.router = mux.NewRouter()

	s.setupRoutes()

	// Настройка CORS
	corsOptions := handlers.AllowedOrigins([]string{"*"})
	corsHeaders := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})
	corsMethods := handlers.AllowedMethods([]string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"})

	// Добавление маршрута для Swagger
	s.router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL(fmt.Sprintf("http://localhost:%d/swagger/doc.json", s.serverPort)),
	))

	// Обернуть маршрутизатор в CORS middleware
	handler := handlers.CORS(corsOptions, corsHeaders, corsMethods)(s.router)

	return http.ListenAndServe(fmt.Sprintf(":%d", s.serverPort), handler)
}
