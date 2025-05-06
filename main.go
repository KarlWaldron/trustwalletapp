package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// 路由
	r.HandleFunc("/block-number", GetBlockNumberHandler).Methods("GET")
	r.HandleFunc("/block/{number}", GetBlockByNumberHandler).Methods("GET")

	log.Println("server setup，Listening port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}