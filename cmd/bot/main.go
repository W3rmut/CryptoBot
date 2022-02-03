package main

import (
	"cryptobot/internal/actions"
	"cryptobot/internal/config"
	"cryptobot/internal/data"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

var botId string

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	config.InitConfig()
	botId = config.ResultConfig.ApiKeys.TelegramKey
	botUrl := fmt.Sprintf("https://api.telegram.org/bot%s/", botId)
	offset := 0
	for {
		resp, err := getUpdate(botUrl+"getUpdates", offset)
		if err != nil {
			log.Printf("Smth went wrong: %s", err)
		}

		for _, update := range resp {
			switch update.Message.Text {
			case "XMR":
				actions.SelectXMR(botUrl+"sendMessage", update.Message.Chat.ChatID)
			case "BTC":
				actions.SelectBTC(botUrl+"sendMessage", update.Message.Chat.ChatID)
			case "ETH":
				actions.SelectETH(botUrl+"sendMessage", update.Message.Chat.ChatID)
			default:
				actions.MainSelectError(botUrl+"sendMessage", update.Message.Chat.ChatID)
			}
			offset = update.UpdateId + 1
		}
	}

	return nil
}

func getUpdate(botUrl string, offset int) ([]data.Update, error) {

	resp, err := http.Get(botUrl + "?offset=" + strconv.Itoa(offset))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var updateResponse data.UpdateResponse
	err = json.Unmarshal(body, &updateResponse)
	if err != nil {
		return nil, err
	}

	return updateResponse.Result, nil
}
