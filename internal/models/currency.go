package models

import "time"

type CurrencyId int64

type Currency struct {
	ID    CurrencyId `json:"id"`
	Slug  string     `json:"slug"`
	Value float64    `json:"value"`
	Date  time.Time  `json:"date"`
	Base  string     `json:"base"`
}

type CreateCurrencyDto struct {
	Slug  string
	Value float64
	Date  time.Time
	Base  string
}
