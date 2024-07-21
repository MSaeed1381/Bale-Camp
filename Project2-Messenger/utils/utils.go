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

func ValidateFileId(fileId string) (bool, error) {
	if len(fileId) == 0 { // TODO other validation
		return false, nil
	}

	requestURL := fmt.Sprintf("http://127.0.0.1:8080/existsFile/%s", fileId)
	res, err := http.Get(requestURL)
	if err != nil {
		return false, errors.New("file Server doesn't up")
	}
	if res.StatusCode != 200 {
		return false, nil
	}
	return true, nil
}

func ContentType(content *messenger.Chat_Message_Content) (contentType string, str string) {
	switch v := content.GetContent().(type) {
	case *messenger.Chat_Message_Content_Text:
		return "text", v.Text
	case *messenger.Chat_Message_Content_Image:
		return "image", v.Image
	case *messenger.Chat_Message_Content_File:
		return "file", v.File
	}
	return "", ""
}
