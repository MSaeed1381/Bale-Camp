package main

import (
	"ConcurrentFileServer/core"
	"context"
	"fmt"
)

func main() {
	var (
		data     = []byte("ali")
		mimeType = "text/plain"
		//filesDirectory = "../files"
		handler = core.NewFileHandlerImpl()
		ctx     = context.Background()
	)

	fileId, err := handler.UploadFile(ctx, data, mimeType)

	fmt.Println(fileId, err)
	//fmt.Println(fileId)
	//assert.Nil(t, err)
	//assert.NotEmpty(t, fileId)
	//
	//dir, err := os.ReadDir(filesDirectory)
	//assert.Nil(t, err)
	//assert.Len(t, dir, 1)
	//assert.False(t, dir[0].IsDir())
	//expectedFileName := fileId //fmt.Sprintf("%s.%s", fileId, utils.GetExtensionByMimeType(mimeType))
	//assert.Equal(t, expectedFileName, dir[0].Name())
	//file, err := os.Open(fmt.Sprintf("%s/%s", filesDirectory, dir[0].Name()))
	//assert.Nil(t, err)
	//buf := make([]byte, len(data))
	//read, err := file.Read(buf)
	//assert.Nil(t, err)
	//assert.Equal(t, len(data), read)
	//assert.True(t, bytes.Equal(data, buf))
}
