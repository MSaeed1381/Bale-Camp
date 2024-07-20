package utils

import (
	"crypto/aes"
	"crypto/cipher"
	cryptorand "crypto/rand"
	"encoding/base64"
	"errors"
	"github.com/joho/godotenv"
	"hash/fnv"
	"io"
	mathrand "math/rand"
	"os"
	"strconv"
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
		b[i] = letterRunes[mathrand.Intn(len(letterRunes))]
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

func Encrypt(plaintext string) (string, error) {
	aesCipher, err := aes.NewCipher([]byte(GetSecretKey()))
	if err != nil {
		return "", errors.New("error creating AES cipher")
	}

	gcm, err := cipher.NewGCM(aesCipher)
	if err != nil {
		return "", errors.New("error creating GCM")
	}

	// We need a 12-byte nonce for GCM (modifiable if you use cipher.NewGCMWithNonceSize())
	// A nonce should always be randomly generated for every encryption.
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(cryptorand.Reader, nonce); err != nil {
		return "", errors.New("error creating nonce")
	}

	// ciphertext here is actually nonce+ciphertext
	// So that when we decrypt, just knowing the nonce size
	// is enough to separate it from the ciphertext.
	ciphertext := gcm.Seal(nonce, nonce, []byte(plaintext), nil)

	encoded := base64.RawURLEncoding.EncodeToString(ciphertext)
	return encoded, nil
}

func Decrypt(ciphertext string) (string, error) {
	aesCipher, err := aes.NewCipher([]byte(GetSecretKey()))
	if err != nil {
		return "", errors.New("error creating AES cipher")
	}

	gcm, err := cipher.NewGCM(aesCipher)
	if err != nil {
		return "", errors.New("error creating GCM")
	}

	ciphertextBytes, _ := base64.RawURLEncoding.DecodeString(ciphertext)
	ciphertext = string(ciphertextBytes)
	// Since we know the ciphertext is actually nonce+ciphertext
	// And len(nonce) == NonceSize(). We can separate the two.
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	plaintext, err := gcm.Open(nil, []byte(nonce), []byte(ciphertext), nil)
	if err != nil {
		return "", errors.New("error decrypting plaintext")
	}
	return string(plaintext), nil
}

func GetEnv(key string) (string, bool) {
	err := godotenv.Load(".env")
	if err != nil {
		return "", false
	}
	return os.Getenv(key), true
}

func GetSecretKey() string {
	secret, ok := GetEnv("SECRET_KEY")
	if !ok { // defaulted value for non-production environment
		return "N1PCdw3M2B1TfJhoaY2mL736p2vCUc47"
	}
	return secret
}

func GetNoWorker() int64 {
	noWorkers, ok := GetEnv("NO_WORKERS")
	if !ok { // defaulted value for non-production environment
		return 16
	}
	parseInt, err := strconv.ParseInt(noWorkers, 10, 64)
	if err != nil {
		return 16
	}
	return parseInt
}
