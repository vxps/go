package main

import (
	"fmt"
	"log"
	"net/http"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusOK)

	_, err := w.Write([]byte("pong"))
	if err != nil {
		log.Printf("failed to write response: %v", err)
	}
}

func main() {
	http.HandleFunc("/ping", pingHandler)

	fmt.Println("Gateway service started on :8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}