package crypto

import (
	"cryptobot/internal/data"
	"encoding/json"
	"fmt"
	"net/http"
)

const bittrexApiURL = "https://api.bittrex.com/api/v1.1/public/getticker?market="

func GetBTC() (float64, error) {
	var priceBTC data.BittrexResponseCourse
	url := bittrexApiURL + "USD-BTC"
	fmt.Println(url)
	req, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	err = json.NewDecoder(req.Body).Decode(&priceBTC)
	if err != nil {
		return 0, err
	}
	fmt.Println(priceBTC)
	return priceBTC.Result.Last, nil
}

func GetETH() (float64, error) {
	var priceBTC data.BittrexResponseCourse
	url := bittrexApiURL + "USD-ETH"
	fmt.Println(url)
	req, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	err = json.NewDecoder(req.Body).Decode(&priceBTC)
	if err != nil {
		return 0, err
	}
	fmt.Println(priceBTC)
	return priceBTC.Result.Last, nil
}
