package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func handleReq(_ http.ResponseWriter, _ *http.Request) {
	fmt.Println("test")
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", handleReq)
	http.ListenAndServe(":8080", router)
}
