package main

import (
	"fmt"
	"database/sql"
	"net/http"
	"encoding/json"
	"time"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 6432
	user     = "coinbase"
	password = "dev"
	dbname   = "coinbase"
  )
  
  func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	  "password=%s dbname=%s sslmode=disable",
	  host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
	  panic(err)
	}
	defer db.Close()
  
	err = db.Ping()
	if err != nil {
	  panic(err)
	}
	

	dbDriver := &DbDriver{db: db}
	go dbDriver.populateTickers()

	fmt.Println(dbDriver.db)
	
  
	fmt.Println("Successfully connected!")
	r := getRouter()
	// r.Get("/", dbDriver.indexHandler)
	
	r.Get("/api/favourites/list", dbDriver.getFaves)

	http.ListenAndServe(":8080", r) //localhost
	// getProducts()

  }

func getProducts() []*Product {
	resp, err := http.Get("https://api.pro.coinbase.com//products")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	var products []*Product
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&products)
	if err != nil {
		panic(err)
	}
	return products
}


// get ticker
func (dbDriver *DbDriver) getTicker(id string) {
	// get data from coinbase api
	resp, err := http.Get("https://api.pro.coinbase.com/products/" + id + "/ticker")

	if err != nil {

	}
	defer resp.Body.Close()

	// create a var to hold resp of json data
	var ticker *Ticker // holds json data
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&ticker)
	if err != nil {
		panic(err)
	}
	tickerData := &TickerData{ID: id, Price: ticker.Price, Time: ticker.Time, Bid: ticker.Bid, Ask: ticker.Ask, Volume: ticker.Volume, Size: ticker.Size}
	dbDriver.refreshTickers(tickerData, id)

}

// refresh tickers
func (dbDriver *DbDriver) refreshTickers(tData *TickerData, id string) {
	row := dbDriver.db.QueryRow("select price from tickers where id = $1", id)

	switch err := row.Scan(&tData.ID); err {
	case sql.ErrNoRows:
		//fmt.Println("NO ID")
		_, err = dbDriver.db.Exec(`INSERT INTO tickers (id, price, time, bid, ask, volume, size) VALUES ($1, $2, $3, $4, $5, $6, $7);`, id, tData.Price, tData.Time, tData.Bid, tData.Ask, tData.Volume, tData.Size) // OK
		if err != nil {
			panic(err)

		}
	case nil:
		//fmt.Println()//"ID EXIST"
		_, err = dbDriver.db.Exec(`UPDATE tickers SET price = $2, time = $3, bid = $4, ask = $5, volume = $6, size = $7 WHERE id = $1;`, id, tData.Price, tData.Time, tData.Bid, tData.Ask, tData.Volume, tData.Size) // OK
		//fmt.Println(`INSERT INTO tickers (id, price, time, bid, ask, volume, size) VALUES ($1, $2, $3, $4, $5, $6, $7)`, id, tData.Price, tData.Time, tData.Bid, tData.Ask, tData.Volume, tData.Size)                 // OK

		if err != nil {
			panic(err)

		}
	default:
		panic(err)
	}

}


func (dbDriver *DbDriver) populateTickers() {
	// loop through products call get ticker
	n := 0
	for {
		go dbDriver.getTicker(getProducts()[n].ID)

		if n >= len(getProducts())-1 {
			n = 0
		} else {
			n++
		}
		time.Sleep(2 * time.Second)
	}
}








