package main

import (
	"ConcurrentFileServer/core"
	"ConcurrentFileServer/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strconv"
)

type Error struct {
	Error string `json:"error"`
}

type Success struct {
	Message string `json:"message"`
}

type PostFile struct {
	FileId string `json:"file_id"`
}

func downloadFile(w http.ResponseWriter, r *http.Request) {
	handler := core.NewFileHandlerImpl()
	fileId := r.PostFormValue("file_id")
	if len(fileId) == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		marshaled, _ := json.Marshal(Error{"file_id is empty"})
		w.Write(marshaled)
		return
	}

	file, _, err := handler.DownloadFile(r.Context(), fileId)
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

func uploadFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	handler := core.NewFileHandlerImpl()
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

	fileId, err := handler.UploadFile(r.Context(), fileData, utils.GetMineTypeByExtension(extension))
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

func hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	jsonResponse, _ := json.Marshal(Success{Message: "Welcome to Upload/Download Server"})
	_, err := w.Write(jsonResponse)
	if err != nil {
		return
	}
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /downloadFile", downloadFile)
	mux.HandleFunc("POST /uploadFile", uploadFile)
	mux.HandleFunc("GET /", hello)

	fmt.Println("Server is running on http://127.0.0.1:8080")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println("Server error:", err)
	}
}
