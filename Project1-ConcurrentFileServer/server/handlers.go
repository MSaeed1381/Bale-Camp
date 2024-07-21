package server

import (
	"ConcurrentFileServer/utils"
	"bytes"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

func (server *Server) DownloadFile(w http.ResponseWriter, r *http.Request) {
	fileId := r.PostFormValue("file_id")
	if len(fileId) == 0 {
		JSONResponse(w, http.StatusBadRequest, Error{"file_id is empty"})
		return
	}

	file, _, err := server.fileHandler.DownloadFile(r.Context(), fileId)
	if err != nil {
		JSONResponse(w, http.StatusBadRequest, Error{err.Error()})
		return
	}

	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)

	part, err := writer.CreateFormFile("file", fileId)
	if _, err := part.Write(file); err != nil {
		JSONResponse(w, http.StatusInternalServerError, Error{err.Error()})
		return
	}

	if err := writer.Close(); err != nil {
		JSONResponse(w, http.StatusInternalServerError, Error{err.Error()})
		return
	}

	w.Header().Set("Content-Type", writer.FormDataContentType())
	w.Header().Set("Content-Length", strconv.Itoa(buffer.Len()))

	if _, err := w.Write(buffer.Bytes()); err != nil {
		JSONResponse(w, http.StatusInternalServerError, Error{err.Error()})
		return
	}
}

func (server *Server) UploadFile(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") == "application/x-www-form-urlencoded" {
		server.UploadURLFile(w, r)
		return
	}

	server.UploadFormDataFile(w, r)
}

func (server *Server) UploadFormDataFile(w http.ResponseWriter, r *http.Request) {
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		JSONResponse(w, http.StatusBadRequest, Error{err.Error()})
		return
	}

	extension := filepath.Ext(fileHeader.Filename)[1:]
	fileData, err := io.ReadAll(file)
	if err != nil {
		JSONResponse(w, http.StatusBadRequest, Error{err.Error()})
		return
	}

	fileId, err := server.fileHandler.UploadFile(r.Context(), fileData, utils.GetMineTypeByExtension(extension))
	if err != nil {
		JSONResponse(w, http.StatusBadRequest, Error{err.Error()})
		return
	}

	JSONResponse(w, http.StatusOK, PostFile{fileId})
}

func (server *Server) UploadURLFile(w http.ResponseWriter, r *http.Request) {
	url := r.PostFormValue("file")
	if len(url) == 0 {
		JSONResponse(w, http.StatusBadRequest, Error{"url is empty"})
	}

	response, err := http.Get(url)
	if err != nil {
		JSONResponse(w, http.StatusBadRequest, Error{err.Error()})
		return
	}
	defer response.Body.Close()

	extension := filepath.Ext(url)[1:]
	fileData, err := io.ReadAll(response.Body)
	if err != nil {
		JSONResponse(w, http.StatusBadRequest, Error{err.Error()})
		return
	}

	fileId, err := server.fileHandler.UploadFile(r.Context(), fileData, utils.GetMineTypeByExtension(extension))
	if err != nil {
		JSONResponse(w, http.StatusBadRequest, Error{err.Error()})
		return
	}

	JSONResponse(w, http.StatusOK, PostFile{fileId})
}

func (server *Server) Hello(w http.ResponseWriter, _ *http.Request) {
	JSONResponse(w, http.StatusOK, Success{Message: "Welcome to Upload/Download Server"})
}

func (server *Server) ExistsFile(w http.ResponseWriter, r *http.Request) {
	fileId := r.PathValue("file_id")

	if _, err := os.Stat(fmt.Sprintf("./files/%s", fileId)); errors.Is(err, os.ErrNotExist) {
		JSONResponse(w, http.StatusNotFound, Error{Error: "file not found"})
		return
	}

	JSONResponse(w, http.StatusOK, Success{Message: "file found!"})
}
