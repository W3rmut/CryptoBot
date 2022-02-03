package crypto

import (
	"bytes"
	"cryptobot/internal/data"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"golang.org/x/text/encoding/charmap"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const bittrexApiURL = "https://api.bittrex.com/api/v1.1/public/getticker?market="
const cbrXMLApiURL = "http://www.cbr.ru/scripts/XML_dynamic.asp?"
const USDCode = "R01235"
const coinGeckoURL = "https://api.coingecko.com/api/v3/simple/price?ids=monero&vs_currencies=usd"

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

func GetXMR() (float64, error) {
	var priceXMR data.CoinGeckoResponse
	url := coinGeckoURL
	fmt.Println(url)
	req, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	err = json.NewDecoder(req.Body).Decode(&priceXMR)
	if err != nil {
		return 0, err
	}
	fmt.Println(priceXMR)
	return priceXMR.Monero.Usd, nil
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

func GetRubCourse() (float64, error) {
	var courseData data.ResponseRUB
	url := cbrXMLApiURL
	nowData := time.Now()

	dateParam := nowData.Format("02/01/2006")

	url = url + fmt.Sprintf("date_req1=%s&date_req2=%s&VAL_NM_RQ=%s", dateParam, dateParam, USDCode)
	fmt.Println("test:", url)
	req, err := http.Get(url)
	if err != nil {
		return 0, err
	}

	dataRead, err := ioutil.ReadAll(req.Body)
	reader := bytes.NewReader(dataRead)
	decoder := xml.NewDecoder(reader)
	fmt.Println(string(dataRead))
	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		return charmap.Windows1251.NewDecoder().Reader(input), nil
	}
	err = decoder.Decode(&courseData)
	if err != nil {
		return 0, err
	}
	fmt.Println("test:", courseData)
	fmt.Println("type: ", reflect.TypeOf(courseData.Record.Value))
	courseData.Record.Value = strings.Replace(courseData.Record.Value, ",", ".", -1)
	result, err := strconv.ParseFloat(courseData.Record.Value, 64)
	fmt.Println(result)
	return result, nil
}
