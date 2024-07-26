package server

import (
	"ConcurrentFileServer/core"
	"fmt"
	"net/http"
)

type FileServer interface {
	Hello(w http.ResponseWriter, r *http.Request)
	DownloadFile(http.ResponseWriter, *http.Request)
	UploadFile(http.ResponseWriter, *http.Request)
	ExistsFile(http.ResponseWriter, *http.Request)
	Serve(addr string) error
}

type Server struct {
	mux         *http.ServeMux
	fileHandler core.FileHandler
}

func NewFileServer() FileServer {
	return &Server{mux: http.NewServeMux(), fileHandler: core.NewFileHandlerImpl()}
}

func (server *Server) Serve(addr string) error {
	server.mux.HandleFunc("POST /downloadFile", server.DownloadFile)
	server.mux.HandleFunc("POST /uploadFile", server.UploadFile)
	server.mux.HandleFunc("GET /existsFile/{file_id}", server.ExistsFile)
	server.mux.HandleFunc("GET /", server.Hello)

	fmt.Printf("Server is running on %s", addr)

	if err := http.ListenAndServe(addr, server.mux); err != nil {
		return err
	}

	return nil
}
