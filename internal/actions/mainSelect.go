package actions

import (
	"bytes"
	"cryptobot/internal/crypto"
	"cryptobot/internal/data"
	"encoding/json"
	"fmt"
	"net/http"
)

func SelectBTC(botURL string, chatID int) error {

	messageData := data.MessageRequest{ChatID: chatID}

	priceBTC, err := crypto.GetBTC()
	if err != nil {
		return err
	}
	rubCourse, err := crypto.GetRubCourse()
	if err != nil {
		fmt.Println("Error getting course:", err)
	}
	messageData.Text = fmt.Sprintf("BTC Price\n \t * USD: %f \n \t * RUB %f", priceBTC, priceBTC*rubCourse)
	buf, err := json.Marshal(messageData)
	if err != nil {
		return err
	}
	_, err = http.Post(botURL, "application/json", bytes.NewBuffer(buf))
	return nil
}

func SelectXMR(botURL string, chatID int) error {
	messageData := data.MessageRequest{ChatID: chatID}
	priceXMR, err := crypto.GetXMR()
	if err != nil {
		return err
	}
	rubCourse, err := crypto.GetRubCourse()
	if err != nil {
		fmt.Println("Error getting course:", err)
	}
	messageData.Text = fmt.Sprintf("XMR Price\n \t * USD: %f \n \t * RUB %f", priceXMR, priceXMR*rubCourse)
	buf, err := json.Marshal(messageData)
	if err != nil {
		return err
	}
	_, err = http.Post(botURL, "application/json", bytes.NewBuffer(buf))

	return nil
}

func SelectETH(botURL string, chatID int) error {
	messageData := data.MessageRequest{ChatID: chatID, Text: "ETH Info"}
	priceBTC, err := crypto.GetETH()
	if err != nil {
		return err
	}
	rubCourse, err := crypto.GetRubCourse()
	if err != nil {
		fmt.Println("Error getting course")
	}
	messageData.Text = fmt.Sprintf("ETH Price\n \t * USD: %f \n \t * RUB %f", priceBTC, priceBTC*rubCourse)
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
