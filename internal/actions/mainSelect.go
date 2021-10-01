package actions

import (
	"bytes"
	"cryptobot/internal/data"
	"encoding/json"
	"fmt"
	"net/http"
)

func SelectBTC(botURL string, chatID int) error {

	messageData := data.MessageRequest{ChatID: chatID, Text: "BTC Info"}
	buf, err := json.Marshal(messageData)
	if err != nil {
		return err
	}

	_, err = http.Post(botURL, "application/json", bytes.NewBuffer(buf))
	return nil
}

func SelectXMR(botURL string, chatID int) error {
	messageData := data.MessageRequest{ChatID: chatID, Text: "XMR Info"}
	buf, err := json.Marshal(messageData)
	if err != nil {
		return err
	}
	fmt.Println(buf)
	resp, err := http.Post(botURL, "application/json", bytes.NewBuffer(buf))
	var error = struct {
		Error interface{}
	}{}
	fmt.Println(resp.Body)
	err = json.NewDecoder(resp.Body).Decode(&error)
	if err != nil {
		fmt.Println(error)
	}
	fmt.Println(error)
	return nil
}

func SelectETH(botURL string, chatID int) error {
	messageData := data.MessageRequest{ChatID: chatID, Text: "ETH Info"}

	buf, err := json.Marshal(messageData)

	if err != nil {
		return err
	}

	_, err = http.Post(botURL, "application/json", bytes.NewBuffer(buf))
	return nil
}

func MainSelectError(botURL string, chatID int) error {
	messageData := data.MessageRequest{ChatID: chatID, Text: "Error. Select crypto"}
	buf, err := json.Marshal(messageData)
	if err != nil {
		return err
	}

	_, err = http.Post(botURL, "application/json", bytes.NewBuffer(buf))
	return nil
}
