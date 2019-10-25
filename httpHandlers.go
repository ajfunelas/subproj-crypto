package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

// func (db *DbDriver) indexHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "index")
// }

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
    var userIDPlaceHolder string = "1"
    // get favourites from db
    rows, err := db.db.Query(`SELECT tickers.id, tickers.price, tickers.time, tickers.bid, tickers.ask, tickers.volume, tickers.size 
                                    FROM tickers 
                                    INNER JOIN user_favourites ON tickers.Id=user_favourites.coin_id
                                    WHERE user_favourites.user_id=$1;`, userIDPlaceHolder)
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


