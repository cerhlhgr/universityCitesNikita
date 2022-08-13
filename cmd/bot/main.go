package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	log.Println("Запустился")
	router := mux.NewRouter()

	router.HandleFunc("/nikitos", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.PostFormValue("type"))

		w.WriteHeader(200)
		_, err := w.Write([]byte("2d085acf"))
		if err != nil {
			log.Println(err.Error())
			return
		}
	})

	server := http.Server{
		Addr:    ":8123",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Println(err.Error())
		return
	}
}
