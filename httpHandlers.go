package main

import (
	"fmt"
	"net/http"
    "encoding/json"
    "time"
    "database/sql"
	// "net/http"
	// "encoding/json"

    "github.com/go-chi/chi"
    // "github.com/gofrs/uuid"
)


func (dbDriver *DbDriver) setGet() {
    r := getRouter()
    r.Post("/api/favourites/list", dbDriver.getFaves)
    http.ListenAndServe(":8080", r)
}

// Get Products Function
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

// Populate Ticker Table Function
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

// Grabs ticker data
func (dbDriver *DbDriver) getTicker(id string) {
	resp, err := http.Get("https://api.pro.coinbase.com/products/" + id + "/ticker") // this gets the data from the API
	
	if err != nil {
	}
	defer resp.Body.Close()
	var ticker *Ticker // this variable holds the ticker that was fetched from the API
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&ticker)
	if err != nil {
		panic(err)
	}
	tickerData := &TickerData{ID: id, Price: ticker.Price, Time: ticker.Time, Bid: ticker.Bid, Ask: ticker.Ask, Volume: ticker.Volume, Size: ticker.Size}
	dbDriver.refreshTickers(tickerData, id)
}

// Refreshes the ticker data
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
func (db *DbDriver) getFaves(w http.ResponseWriter, r *http.Request) {
	var (
        userID     UserID
        tickerList []ShortTicker
        tickerId   string
        price      string
        time       string
        bid        string
        ask        string
        volume     string
        size       string
    )
    decoder := json.NewDecoder(r.Body)
    err := decoder.Decode(&userID)
    // get favourites from db
    rows, err := db.db.Query(`SELECT tickers.id, tickers.price, tickers.time, tickers.bid, tickers.ask, tickers.volume, tickers.size 
                                    FROM tickers 
                                    INNER JOIN user_favourites ON tickers.Id=user_favourites.coin_id
                                    WHERE user_favourites.user_id=$1;`, userID.UID)
if err != nil {
    fmt.Println(err)
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
}

defer rows.Close()
for rows.Next() {
    err := rows.Scan(
        &tickerId, &price, &time, &bid, &ask, &volume, &size,
    )
    if err != nil {
        fmt.Println(err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    t := ShortTicker{TickerId: tickerId, Price: price, Time: time, Bid: bid, Ask: ask, Volume: volume, Size: size}
    fmt.Println(t)
    tickerList = append(tickerList, t)
}
w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(tickerList)
}

func getRouter() chi.Router{
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi Owen")
		
})
return r
}

func (db *DbDriver) tglFave(w http.ResponseWriter, r *http.Request) {

}

