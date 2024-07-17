package utils

import (
	"github.com/joho/godotenv"
	"hash/fnv"
	"math/rand"
	"os"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func GetExtensionByMimeType(mimeType string) string {
	switch mimeType {
	case "image/png":
		return "png"
	case "image/jpeg":
		return "jpg"
	case "text/plain":
		return "txt"
	default:
		return ""
	}
}

func GetMineTypeByExtension(extension string) string {
	switch extension {
	case "png":
		return "image/png"
	case "jpg":
		return "image/jpeg"
	case "txt":
		return "text/plain"
	default:
		return ""
	}
}

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func HashFileContent(file []byte) uint64 { // TODO
	hash := fnv.New64a()
	_, err := hash.Write(file)
	if err != nil {
		return 0
	}
	return hash.Sum64()
}

func GetEnv(key string) (string, bool) {
	err := godotenv.Load(".env")
	if err != nil {
		return "", false
	}

	return os.Getenv(key), true
}
