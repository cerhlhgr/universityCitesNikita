package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Form)

		w.WriteHeader(200)
		w.Write([]byte("ok"))
	})

	server := http.Server{
		Addr:    "2.59.42.22:8000",
		Handler: router,
	}

	server.ListenAndServe()
}
