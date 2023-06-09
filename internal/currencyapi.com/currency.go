package currencyapi_com

import "time"

type ApiKey string

type CurrencyApiCom struct {
	ApiKey ApiKey
}

func (c *CurrencyApiCom) getCurrenciesByDate(currencies []string, date time.Time) {
	// https://api.currencyapi.com/v3/historical?apikey=iOfHhECk07obbE5XuGOIxZNhqgESdVMmSypw2LqT&currencies=EUR,RUB,JPY,USD&date=2014-01-01
	panic("TODO: implement me")
}

func (c *CurrencyApiCom) getCurrenciesFromTo(from, to string) {
	// https://api.currencyapi.com/v3/historical?apikey=iOfHhECk07obbE5XuGOIxZNhqgESdVMmSypw2LqT&base_currency=USD&currencies=EUR,RUB,JPY,USD&date=2014-01-01
	panic("TODO: implement me")
}
