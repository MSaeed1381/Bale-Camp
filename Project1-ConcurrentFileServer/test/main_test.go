package test

import (
	"ConcurrentFileServer/core"
	"ConcurrentFileServer/pkg"
	"ConcurrentFileServer/utils"
	"bytes"
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func TestUpload(t *testing.T) {
	var (
		data           = []byte("ali")
		mimeType       = "text/plain"
		filesDirectory = "../files"
		handler        = core.NewFileHandlerImpl()
		ctx            = context.Background()
	)

	fileId, err := handler.UploadFile(ctx, data, mimeType)
	fmt.Println(fileId)
	assert.Nil(t, err)
	assert.NotEmpty(t, fileId)

	dir, err := os.ReadDir(filesDirectory)
	assert.Nil(t, err)
	assert.Len(t, dir, 1)
	assert.False(t, dir[0].IsDir())
	expectedFileName := fileId //fmt.Sprintf("%s.%s", fileId, utils.GetExtensionByMimeType(mimeType))
	assert.Equal(t, expectedFileName, dir[0].Name())
	file, err := os.Open(fmt.Sprintf("%s/%s", filesDirectory, dir[0].Name()))
	assert.Nil(t, err)
	buf := make([]byte, len(data))
	read, err := file.Read(buf)
	assert.Nil(t, err)
	assert.Equal(t, len(data), read)
	assert.True(t, bytes.Equal(data, buf))
}

func TestDownload(t *testing.T) {
	var (
		data = []byte("ali")

		mimeType       = "text/plain"
		filesDirectory = "../files"
		fileId         = "tmp"
		handler        = core.NewFileHandlerImpl()
		ctx            = context.Background()
	)

	create, err := os.Create(fmt.Sprintf("%s/%s.%s", filesDirectory, fileId, utils.GetExtensionByMimeType(mimeType)))
	assert.Nil(t, err)

	write, err := create.Write(data)
	assert.Nil(t, err)
	assert.NotEmpty(t, write)

	filename := fmt.Sprintf("%s.%s", fileId, utils.GetExtensionByMimeType(mimeType))
	fmt.Println(filename)
	resultFile, resultMimeType, err := handler.DownloadFile(ctx, filename)
	assert.NotEmpty(t, resultMimeType)

	fmt.Println(resultFile)
	fmt.Println(data)
	assert.True(t, bytes.Equal(resultFile, data))
}

func TestUploadAndDownloadScenario(t *testing.T) {
	var (
		data     = []byte("ali")
		mimeType = "text/plain"
		handler  = core.NewFileHandlerImpl()
		ctx      = context.Background()
	)

	fileId, err := handler.UploadFile(ctx, data, mimeType)
	assert.Nil(t, err)
	assert.NotEmpty(t, fileId)

	file, downloadMimeType, err := handler.DownloadFile(ctx, fileId)
	assert.Nil(t, err)
	assert.NotEmpty(t, downloadMimeType)
	assert.Equal(t, mimeType, downloadMimeType)
	assert.True(t, bytes.Equal(file, data))
}

func TestUploadAndDownloadConcurrent(t *testing.T) {
	var (
		workerPool = pkg.NewWorkerPool(25)
		handler    = core.NewFileHandlerImpl()
		mimeType   = "text/plain"
	)

	ticker := time.NewTicker(100 * time.Millisecond)
	timer := time.NewTimer(5 * time.Second)

	condition := true
	for condition {
		select {
		case <-timer.C:
			condition = false
		case <-ticker.C:
			workerPool.SubmitJob(func() {
				ctx := context.Background()
				data := []byte(utils.RandStringRunes(1000))
				fileId, err := handler.UploadFile(ctx, data, mimeType)
				assert.Nil(t, err)
				time.Sleep(1 * time.Second)
				file, s, err := handler.DownloadFile(ctx, fileId)
				assert.Nil(t, err)
				assert.NotEmpty(t, s)
				assert.NotEmpty(t, file)
				assert.True(t, bytes.Equal(file, data))
			})
		}
	}

}
