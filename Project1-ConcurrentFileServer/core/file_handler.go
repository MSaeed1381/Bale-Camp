package core

import (
	"ConcurrentFileServer/utils"
	"context"
	"errors"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// BASEPATH folder for downloading and uploading files
const BASEPATH = "./files"

type FileHandler interface {
	UploadFile(ctx context.Context, file []byte, mimeType string) (string, error)
	DownloadFile(ctx context.Context, fileID string) ([]byte, string, error)
}

type FileHandlerImpl struct {
}

func NewFileHandlerImpl() FileHandler {
	return &FileHandlerImpl{}
}

// function to read a chunk of data from file and copy it to shared memory
func readChunk(offset int, chunkSize int, wg *sync.WaitGroup, file *os.File, byteArray *[]byte) (bool, error) {
	defer wg.Done()
	// copy from file from offset with length chunkSize to byte array (offset:offset+chunkSize) that is isolate form other goroutines
	_, err := file.ReadAt((*byteArray)[offset:offset+chunkSize], int64(offset))

	if err != nil {
		return false, errors.New("read chunk error")
	}
	return true, nil
}

func writeChunk(offset int, chunk int, wg *sync.WaitGroup, file *os.File, byteArray *[]byte) (bool, error) {
	defer wg.Done()
	// like read from memory, i will write in memory
	_, err := file.WriteAt((*byteArray)[offset:offset+chunk], int64(offset))

	if err != nil {
		return false, errors.New("write chunk error")
	}
	return true, nil
}

func (f *FileHandlerImpl) UploadFile(ctx context.Context, file []byte, mimeType string) (string, error) {
	hashValue := utils.HashFileContent(file)
	filename := fmt.Sprintf("%d:%d.%s", hashValue, hashValue, utils.GetExtensionByMimeType(mimeType))
	name := filepath.Join(BASEPATH, filename)
	fmt.Println(name)
	newFile, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
		return "", errors.New("create file error")
	}

	defer func(newFile *os.File) {
		err := newFile.Close()
		if err != nil {
			panic("close file error")
		}
	}(newFile)

	const numberOfWorkers = 16
	var workers = int(math.Min(numberOfWorkers, float64(len(file))))

	chunkSize := len(file) / workers
	reminder := len(file) % workers

	var wg sync.WaitGroup
	wg.Add(workers + 1)

	offset := 0
	temp := chunkSize
	for i := 0; i < workers+1; i++ {
		if i == workers {
			temp = reminder
		}
		go writeChunk(offset, temp, &wg, newFile, &file)
		offset += chunkSize
	}

	wg.Wait()
	return filename, nil
}

func (f *FileHandlerImpl) DownloadFile(ctx context.Context, fileID string) ([]byte, string, error) {
	filePath := filepath.Join(BASEPATH, fileID)
	file, openFileErr := os.Open(filePath)
	if openFileErr != nil {
		return nil, "", errors.New("open file fail")
	}

	// wrap error
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic("close file error")
		}
	}(file)

	stat, err := file.Stat()
	if err != nil {
		return nil, "", errors.New("stat file fail")
	}

	fileLength := int(stat.Size())
	byteArray := make([]byte, fileLength) // result

	const numberOfWorkers = 16
	var workers = int(math.Min(numberOfWorkers, float64(fileLength)))

	chunks := fileLength / workers
	reminder := fileLength % workers

	var wg sync.WaitGroup
	wg.Add(workers + 1)

	offset := 0
	chunkSize := chunks
	for i := 0; i < workers+1; i++ {
		if i == workers {
			chunkSize = reminder
		}
		go readChunk(offset, chunkSize, &wg, file, &byteArray)
		offset += chunks
	}

	wg.Wait()

	splitFileID := strings.Split(fileID, ".")
	return byteArray, utils.GetMineTypeByExtension(splitFileID[len(splitFileID)-1]), nil
}
