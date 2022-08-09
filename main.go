package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.PostFormValue("type"))

		w.WriteHeader(200)
		w.Write([]byte("ok"))
	})

	server := http.Server{
		Addr:    "2.59.42.22:80",
		Handler: router,
	}

	server.ListenAndServe()
}
