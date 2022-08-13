package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var access_token string = "333687e3236c2cf0a7f24a0d7832cf092a3499b5e1f70b66c8c511bf0b43dbec75063d53635611272cd92"

type vars struct {
	Type   string `json:"type"`
	Object map[string]map[string]interface{}
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

		//cities := map[string]string{}

		log.Println(params.Type)
		switch params.Type {
		case "message_new":
			if params.Object["message"]["text"] == "asd" {
				log.Println(fmt.Sprintf("%v", params.Object["message"]["text"]))
				_, err := http.Get("https://api.vk.com/method/messages.send?" + "user_id=" + fmt.Sprintf("%v", params.Object["message"]["from_id"]) + "&message=" + "хуй" + "&access_token=" + access_token)
				if err != nil {
					log.Println(err.Error())
					return
				}
				//get, err := http.Get("https://api.vk.com/method/database.getCities?v=5.52&access_token=" + access_token)
				if err != nil {
					//decoder = json.NewDecoder(get.Body)
					//decoder.Decode(&cities)
					//log.Println(cities)
					return
				}
			} else {

			}
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
