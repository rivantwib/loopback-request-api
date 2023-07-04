package main

import (
	"log"
	"fmt"
	"net/http"
	"strings"
	
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)


func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/ip", GetIPAddressHandler).Methods("GET")

	// Starting Server
	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}


func GetIPAddressHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ip := strings.Split(r.RemoteAddr, ":")[0]

	response := fmt.Sprintf("Your IP address: %s\n", ip)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}