package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"reflect"
)

type KeyboardButton struct {
	Text            string `json:"text"`
	RequestContact  bool   `json:"request_contact"`
	RequestLocation bool   `json:"request_location"`
}

type InlineKeyboardButton struct {
	Text         string `json:"text"`
	CallbackData string `json:"callback_data"`
	Url          string `json:"url"`
}

type ReplyMarkup struct {
	InlineKeyboard [][]interface{} `json:"inline_keyboard"`
	Keyboard       [][]interface{} `json:"keyboard"`
	ResizeKeyboard bool            `json:"resize_keyboard"`
	OnTimeKeyboard bool            `json:"one_time_keyboard"`
	Selective      bool            `json:"selective"`
}

type SendMessage struct {
	ChatID      interface{}  `json:"chat_id"`
	Text        string       `json:"text"`
	ParseMode   string       `json:"parse_mode"`
	ReplyMarkup *ReplyMarkup `json:"reply_markup"`
}

func isEmptyArray(array [][]interface{}) bool {
	if len(array) == 0 || len(array[0]) == 0 {
		return true
	}
	return false
}

func typeAssertion(array *[][]interface{}, keyboard interface{}) error {
	for i := 0; i < len(*array); i++ {
		for j := 0; j < len((*array)[i]); j++ {
			switch (*array)[i][j].(type) {
			case string:
				(*array)[i][j] = (*array)[i][j].(string)
			case map[string]interface{}:
				keyboardType := reflect.TypeOf(keyboard)
				newKeyboard := reflect.New(keyboardType).Interface()

				m, _ := json.Marshal((*array)[i][j])
				err := json.Unmarshal(m, &newKeyboard)

				if err != nil {
					return err
				}

				(*array)[i][j] = reflect.ValueOf(newKeyboard).Elem().Interface()
			default:
				panic("Unsupported type")
			}
		}
	}
	return nil
}

func ReadSendMessageRequest(fileName string) (*SendMessage, error) {
	// read file as byte array
	smByteResult, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("unable to read JSON file")
	}

	// unmarshal json as SendMessage Structure
	var sm *SendMessage

	if err := json.Unmarshal(smByteResult, &sm); err != nil {
		return nil, errors.New("bad format, unable to marshal SendMessage JSON")
	}

	// validate chatID, Text
	if sm.ChatID == nil {
		return nil, errors.New("chat_id is empty")
	} else if sm.Text == "" {
		return nil, errors.New("text is empty")
	}

	// handling replyMarkup Structure
	if sm.ReplyMarkup == nil ||
		(isEmptyArray(sm.ReplyMarkup.InlineKeyboard) && isEmptyArray(sm.ReplyMarkup.Keyboard)) {
		return sm, nil
	} else if !isEmptyArray(sm.ReplyMarkup.InlineKeyboard) {
		err := typeAssertion(&sm.ReplyMarkup.InlineKeyboard, InlineKeyboardButton{})
		if err != nil {
			return nil, err
		}
	} else if !isEmptyArray(sm.ReplyMarkup.Keyboard) {
		err := typeAssertion(&sm.ReplyMarkup.Keyboard, KeyboardButton{})
		if err != nil {
			return nil, err
		}
	}
	return sm, nil
}

func main() {
	msg, err := ReadSendMessageRequest("src/input_sample2.json")
	fmt.Println(msg, err)
}
