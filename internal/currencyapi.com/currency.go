package currencyapi_com

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/regimentor/currency-calc/pkg"
)

type ApiKey string

type CurrencyApiCom struct {
	ApiKey ApiKey
}

func NewCurrencyApiCom(apiKey ApiKey) *CurrencyApiCom {
	return &CurrencyApiCom{ApiKey: apiKey}
}

type ResponseCurrenciesData struct {
	Code  string  `json:"code"`
	Value float64 `json:"value"`
}
type CurrenciesComResponse struct {
	Meta struct {
		LastUpdatedAt time.Time `json:"last_updated_at"`
	} `json:"meta"`
	Data map[string]ResponseCurrenciesData `json:"data"`
}

func (c *CurrencyApiCom) GetCurrenciesByDate(currencies []string, date time.Time) (*CurrenciesComResponse, error) {
	// https://api.currencyapi.com/v3/historical?apikey=APIKEY&currencies=EUR,RUB,JPY,USD&date=2014-01-01

	log.Printf("CurrencyApiCom.GetCurrenciesByDate get currencies by date %v", date)

	year, month, day := date.Date()
	url := fmt.Sprintf("https://api.currencyapi.com/v3/historical?apikey=%s&currencies=%s&date=%d-%d-%d",
		c.ApiKey, strings.Join(currencies, ","), year, month, day)

	bytes, err := pkg.HttpRequest("GET", url)
	if err != nil {
		return nil, fmt.Errorf("request due err: %v", err)
	}

	var response *CurrenciesComResponse
	if err := json.Unmarshal(bytes, &response); err != nil {
		return nil, fmt.Errorf("unmarshal response due err: %v", err)
	}

	return response, nil
}

func (c *CurrencyApiCom) getCurrenciesFromTo(base string, currencies []string, date time.Time) (*CurrenciesComResponse, error) {
	// https://api.currencyapi.com/v3/historical?apikey=APIKEY&base_currency=USD&currencies=EUR,RUB,JPY,USD&date=2014-01-01
	year, month, day := date.Date()
	url := fmt.Sprintf("https://api.currencyapi.com/v3/historical?apikey=%s&base_currency=%s&currencies=%s&date=%d-%d-%d",
		c.ApiKey, base, strings.Join(currencies, ","), year, month, day)

	bytes, err := pkg.HttpRequest("GET", url)
	if err != nil {
		return nil, fmt.Errorf("request due err: %v", err)
	}

	var response *CurrenciesComResponse
	if err := json.Unmarshal(bytes, &response); err != nil {
		return nil, fmt.Errorf("unmarshal response due err: %v", err)
	}

	return response, nil
}
