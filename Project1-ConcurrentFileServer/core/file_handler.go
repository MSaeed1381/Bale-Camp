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
// copy from file from offset with length chunkSize to byte array (offset:offset+chunkSize) that is isolating form other goroutines
// wrapping error
func readChunk(offset int, chunkSize int, wg *sync.WaitGroup, file *os.File, byteArray *[]byte) error {
	defer wg.Done()
	if _, err := file.ReadAt((*byteArray)[offset:offset+chunkSize], int64(offset)); err != nil {
		return errors.New("read chunk error")
	}
	return nil
}

// like read from memory will write in memory
// wrapping error
func writeChunk(offset int, chunkSize int, wg *sync.WaitGroup, file *os.File, byteArray *[]byte) error {
	defer wg.Done()
	if _, err := file.WriteAt((*byteArray)[offset:offset+chunkSize], int64(offset)); err != nil {
		return errors.New("write chunkSize error")
	}
	return nil
}

func (f *FileHandlerImpl) UploadFile(ctx context.Context, file []byte, mimeType string) (string, error) {
	hashValue, err := utils.HashFileContent(file)
	if err != nil {
		return "", err
	}

	encryptedHashValue, err := utils.Encrypt(strconv.FormatUint(hashValue, 10))
	if err != nil {
		return "", err
	}

	filename := fmt.Sprintf("%d:%s.%s", hashValue, encryptedHashValue, utils.GetExtensionByMimeType(mimeType))
	newFile, err := os.Create(filepath.Join(BASEPATH, filename))
	if err != nil {
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
	wg.Add(NoWorkers)

	offset := 0
	temp := chunkSize
	for i := 0; i < NoWorkers+1; i++ {
		if i == NoWorkers {
			temp = reminder
			if temp == 0 {
				break
			}
			wg.Add(1)
		}

		if err := writeChunk(offset, temp, &wg, newFile, &file); err != nil {
			return "", err
		}

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
	wg.Add(NoWorkers)

	offset := 0
	temp := chunkSize
	for i := 0; i < NoWorkers+1; i++ {
		if i == NoWorkers {
			temp = reminder
			if temp == 0 {
				break
			}
			wg.Add(1)
		}

		if err := readChunk(offset, temp, &wg, file, &byteArray); err != nil {
			return nil, "", err
		}

		offset += chunkSize
	}

	wg.Wait()

	return byteArray, utils.GetMineTypeByExtension(filepath.Ext(fileID)[1:]), nil
}
