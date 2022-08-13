package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type vars struct {
	Type   string `json:"type"`
	Object map[string]interface{}
}

func main() {
	log.Println("Запустился")
	router := mux.NewRouter()

	router.HandleFunc("/nikitos", func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)

		params := vars{}
		err := decoder.Decode(&params)
		if err != nil {
			log.Println(err.Error())
			return
		}

		log.Println(params.Type)
		switch params.Type {
		case "message_new":
			log.Println(params.Object)
		}

		w.WriteHeader(200)
		_, err = w.Write([]byte("ok"))
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
