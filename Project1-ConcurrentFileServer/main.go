package main

import (
	"ConcurrentFileServer/core"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func downloadFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	handler := core.NewFileHandlerImpl()
	ctx := context.Background()
	fileId := r.PostFormValue("file_id")

	fmt.Println("fileId:", fileId)
	file, _, err := handler.DownloadFile(ctx, fileId)
	fmt.Println(err)
	fmt.Println(file)
	w.Write(file)
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	handler := core.NewFileHandlerImpl()
	ctx := context.Background()
	file, header, _ := r.FormFile("file")
	fmt.Println(header)
	fmt.Println(file)

	fileData, err := ioutil.ReadAll(file)
	fileId, err := handler.UploadFile(ctx, fileData, "text/plain")
	fmt.Println(err)

	data := struct {
		FileId string `json:"file_id"`
	}{
		FileId: fileId,
	}

	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		return
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/downloadFile", downloadFile)
	mux.HandleFunc("/uploadFile", uploadFile)
	mux.HandleFunc("/", hello)

	fmt.Printf("Server is running on 127.0.0.1:8080")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println("Server error:", err)
	}

}
