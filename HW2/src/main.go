package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
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
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
	Keyboard       [][]KeyboardButton       `json:"keyboard"`
	ResizeKeyboard bool                     `json:"resize_keyboard"`
	OnTimeKeyboard bool                     `json:"one_time_keyboard"`
	Selective      bool                     `json:"selective"`
}

type SendMessage struct {
	ChatID      interface{} `json:"chat_id"`
	Text        string      `json:"text"`
	ParseMode   string      `json:"parse_mode"`
	ReplyMarkup interface{} `json:"reply_markup"`
}

func ReadSendMessageRequest(fileName string) (*SendMessage, error) {
	// read file as byte array
	smByteResult, readFileErr := os.ReadFile(fileName)
	if readFileErr != nil {
		fmt.Printf("Unable to read JSON file due to %s\n", readFileErr)
		return nil, errors.New("unable to read JSON file")
	}

	// unmarshal json as SendMessage Structure
	var sm *SendMessage
	unmarshalError := json.Unmarshal(smByteResult, &sm)
	if unmarshalError != nil {
		fmt.Printf("Unable to marshal SendMessage JSON due to %s\n", unmarshalError)
		return nil, errors.New("unable to marshal SendMessage JSON")
	}

	// validate chatID, Text
	if sm.ChatID == nil {
		return nil, errors.New("chat_id is empty")
	} else if sm.Text == "" {
		return nil, errors.New("text is empty")
	}

	// handling replyMarkup Structure
	if sm.ReplyMarkup == nil {
		return sm, nil
	}

	var rm *ReplyMarkup
	rmByteArr, _ := json.Marshal(sm.ReplyMarkup) // ignore replyMarkup marshal error due to Marshaling SendMessage
	unmarshalErr := json.Unmarshal(rmByteArr, &rm)
	if unmarshalErr != nil {
		return nil, unmarshalErr
	}

	sm.ReplyMarkup = *rm
	return sm, nil
}