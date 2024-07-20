package server

import (
	"ConcurrentFileServer/utils"
	"bytes"
	"encoding/json"
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
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		marshaled, _ := json.Marshal(Error{"file_id is empty"})
		w.Write(marshaled)
		return
	}

	file, _, err := server.fileHandler.DownloadFile(r.Context(), fileId)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		marshaled, _ := json.Marshal(Error{err.Error()})
		w.Write(marshaled)
		return
	}

	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)

	part, err := writer.CreateFormFile("file", fileId)

	if _, err := part.Write(file); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		marshaled, _ := json.Marshal(Error{err.Error()})
		w.Write(marshaled)
		return
	}

	if err := writer.Close(); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		marshaled, _ := json.Marshal(Error{err.Error()})
		w.Write(marshaled)
		return
	}

	w.Header().Set("Content-Type", writer.FormDataContentType())
	w.Header().Set("Content-Length", strconv.Itoa(buffer.Len()))
	w.WriteHeader(http.StatusOK)

	if _, err := w.Write(buffer.Bytes()); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		marshaled, _ := json.Marshal(Error{err.Error()})
		w.Write(marshaled)
		return
	}

}

func (server *Server) UploadFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		marshaled, _ := json.Marshal(Error{err.Error()})
		w.Write(marshaled)
		return
	}

	extension := filepath.Ext(fileHeader.Filename)[1:]

	fileData, err := io.ReadAll(file)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		marshaled, _ := json.Marshal(Error{err.Error()})
		w.Write(marshaled)
		return
	}

	fileId, err := server.fileHandler.UploadFile(r.Context(), fileData, utils.GetMineTypeByExtension(extension))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		marshaled, _ := json.Marshal(Error{err.Error()})
		w.Write(marshaled)
		return
	}

	marshaledFileId, _ := json.Marshal(PostFile{fileId})

	w.WriteHeader(http.StatusOK)
	w.Write(marshaledFileId)
}

func (server *Server) Hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	jsonResponse, _ := json.Marshal(Success{Message: "Welcome to Upload/Download Server"})
	_, err := w.Write(jsonResponse)
	if err != nil {
		return
	}
}

func (server *Server) ExistsFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fileId := r.PathValue("file_id")

	if _, err := os.Stat(fmt.Sprintf("./files/%s", fileId)); errors.Is(err, os.ErrNotExist) {
		w.WriteHeader(http.StatusNotFound)
		jsonResponse, _ := json.Marshal(Error{Error: "file not found"})
		w.Write(jsonResponse)
		return
	}

	jsonResponse, _ := json.Marshal(Success{Message: "file found!"})
	_, err := w.Write(jsonResponse)
	if err != nil {
		return
	}
}
