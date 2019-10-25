package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	_ "github.com/lib/pq"
)

func getRouter() chi.Router{
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi Owen")
		
})

r.Get("/api/favourites/list", func(w http.ResponseWriter, r *http.Request) {
	// dbDriver.getFaves()
		
})



return r
}




// http.ListenAndServe(":8080")
