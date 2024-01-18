package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func writeEntry(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	defer r.Body.Close()

	var data map[string]interface{}
	json.Unmarshal(body, &data)

	var value = data["value"].(int)

	log.Println("Received value:", value)

	if value < 1000 {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/write/", writeEntry).Methods("POST")
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":6000", handler))
}
