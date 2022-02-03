package data

type BittrexResponseCourse struct {
	Success bool   `json:"success"`
	Result  Result `json:"result"`
}
type Result struct {
	Bid  float64 `json:"Bid"`
	Ask  float64 `json:"Ask"`
	Last float64 `json:"Last"`
}

type CoinGeckoResponse struct {
	Monero struct {
		Usd float64 `json:"usd"`
	} `json:"monero"`
}
