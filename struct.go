package main

import 
"database/sql"

type DbDriver struct {
	db *sql.DB
}

type Ticker struct {
	Coin_ID string `json: trade_id"`
	Price string `json: "price"`
	Time    string `json: "size"`
	Bid     string `json: "time"`
	Ask     string `json: "bid"`
	Volume  string `json: "ask"`
	Size    string `json: "volume"`
}
type TickerData struct {
	ID     string
	Price  string
	Time   string
	Bid    string
	Ask    string
	Volume string
	Size   string
}

// coins - query data
type Product struct {
	ID             string `json: "id"`
	BaseCurrency   string `json: "base_currency"`
	QuoteCurrency  string `json: "quote_currency"`
	BaseMinSize    string `json: "base_min_size"`
	BaseMaxSize    string `json: "base_max_size"`
	QuoteIncrement string `json: "quote_increment"`
	BaseIncrement  string `json: "base_increment"`
	DisplayName    string `json: "display_name"`
	MinMarketFunds string `json: "min_market_funds"`
	MaxMarketFunds string `json: "max_market_funds"`
	MarginEnabled  bool   `json: "margin_enabled"`
	PostOnly       bool   `json: "post_only"`
	LimitOnly      bool   `json: "limit_only"`
	CancelOnly     bool   `json: "cancel_only"`
	Status         string `json: "status"`
	StatusMessage  string `json: "status_message"`
}

type ShortTicker struct {
	TickerId string
	Price  string
	Time   string
	Bid    string
	Ask    string
	Volume string
	Size   string
}

type UserID struct {
	UID string `json: "uid"`
}