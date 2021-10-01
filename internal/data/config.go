package data

type Config struct {
	ApiKeys ApiKeysCFG `toml:"api-keys"`
}

type ApiKeysCFG struct {
	TelegramKey string
	BinanceKey  string
}
