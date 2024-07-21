package utils

import (
	"Messenger/messenger"
	"errors"
	"fmt"
	"net/http"
	"unicode"
)

func hasLetter(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) {
			return true
		}
	}
	return false
}

func hasNumber(s string) bool {
	for _, r := range s {
		if unicode.IsNumber(r) {
			return true
		}
	}
	return false
}

// ValidateUsername check that username has length more than 3 and has both Letters and numbers
func ValidateUsername(username string) bool {
	if len(username) < 3 || !(hasLetter(username) && hasNumber(username)) {
		return false
	}
	return true
}

func ValidateFileId(fileId string) error {
	if len(fileId) == 0 {
		return errors.New("file-id is empty")
	}
	requestURL := fmt.Sprintf("http://127.0.0.1:8080/existsFile/%s", fileId)
	res, err := http.Get(requestURL)
	if err != nil {
		return errors.New("file Server doesn't up")
	}
	if res.StatusCode != 200 {
		return errors.New("invalid file-id")
	}
	return nil
}

type ContentType int

const (
	TEXT  ContentType = 0
	IMAGE ContentType = 1
	FILE  ContentType = 2
)

func GetContentType(content *messenger.Chat_Message_Content) (contentType ContentType, contentStr string) {
	switch v := content.GetContent().(type) {
	case *messenger.Chat_Message_Content_Text:
		return TEXT, v.Text
	case *messenger.Chat_Message_Content_Image:
		return IMAGE, v.Image
	case *messenger.Chat_Message_Content_File:
		return FILE, v.File
	}
	return 0, ""
}
