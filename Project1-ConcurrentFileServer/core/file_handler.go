package core

import (
	"ConcurrentFileServer/utils"
	"context"
	"errors"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strconv"
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
// copy from file from offset with length chunkSize to byte array (offset:offset+chunkSize) that is isolate form other goroutines
// wrapping error
func readChunk(offset int, chunkSize int, wg *sync.WaitGroup, file *os.File, byteArray *[]byte) (bool, error) {
	defer wg.Done()
	if _, err := file.ReadAt((*byteArray)[offset:offset+chunkSize], int64(offset)); err != nil {
		return false, errors.New("read chunk error")
	}
	return true, nil
}

// like read from memory, i will write in memory
// wrapping error
func writeChunk(offset int, chunkSize int, wg *sync.WaitGroup, file *os.File, byteArray *[]byte) (bool, error) {
	defer wg.Done()
	if _, err := file.WriteAt((*byteArray)[offset:offset+chunkSize], int64(offset)); err != nil {
		return false, errors.New("write chunkSize error")
	}
	return true, nil
}

func (f *FileHandlerImpl) UploadFile(ctx context.Context, file []byte, mimeType string) (string, error) {
	hashValue := utils.HashFileContent(file)
	encryptedHashValue, err := utils.Encrypt(strconv.FormatUint(hashValue, 10))
	if err != nil {
		return "", err
	}

	filename := fmt.Sprintf("%d:%s.%s", hashValue, encryptedHashValue, utils.GetExtensionByMimeType(mimeType))
	newFile, err := os.Create(filepath.Join(BASEPATH, filename))
	if err != nil {
		fmt.Println(err)
		return "", errors.New("create file error")
	}

	// close the file
	defer func(newFile *os.File) {
		if err := newFile.Close(); err != nil {
			panic("close file error")
		}
	}(newFile)

	// noWorkers goroutine for chunks and 1 for reminder work
	var NoWorkers = int(math.Min(float64(utils.GetNoWorker()), float64(len(file))))
	chunkSize := len(file) / NoWorkers
	reminder := len(file) % NoWorkers

	var wg sync.WaitGroup
	wg.Add(NoWorkers + 1)

	offset := 0
	temp := chunkSize
	for i := 0; i < NoWorkers+1; i++ {
		if i == NoWorkers {
			temp = reminder
		}

		go func() {
			_, err := writeChunk(offset, temp, &wg, newFile, &file)
			if err != nil {
				panic(err)
			}
		}()

		offset += chunkSize
	}

	wg.Wait()
	return filename, nil
}

func (f *FileHandlerImpl) DownloadFile(ctx context.Context, fileID string) ([]byte, string, error) {
	filePath := filepath.Join(BASEPATH, fileID)
	file, err := os.Open(filePath)
	if err != nil {
		return nil, "", errors.New("open file fail")
	}

	// wrap error (close the file)
	defer func(file *os.File) {
		if err := file.Close(); err != nil {
			panic("close file error")
		}
	}(file)

	stat, err := file.Stat()
	if err != nil {
		return nil, "", errors.New("stat file fail")
	}

	fileLength := int(stat.Size())
	byteArray := make([]byte, fileLength) // result

	var NoWorkers = int(math.Min(float64(utils.GetNoWorker()), float64(fileLength)))

	chunkSize := fileLength / NoWorkers
	reminder := fileLength % NoWorkers

	var wg sync.WaitGroup
	wg.Add(NoWorkers + 1)

	offset := 0
	temp := chunkSize
	for i := 0; i < NoWorkers+1; i++ {
		if i == NoWorkers {
			temp = reminder
		}

		go func() {
			_, err := readChunk(offset, temp, &wg, file, &byteArray)
			if err != nil {
				panic(err)
			}
		}()

		offset += chunkSize
	}

	wg.Wait()

	return byteArray, utils.GetMineTypeByExtension(filepath.Ext(fileID)[1:]), nil
}
