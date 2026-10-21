package main

import (
	"gateway/api"
	"log"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/budgets", api.GetBudgets)
	mux.HandleFunc("GET /api/transactions", api.GetTransaction)
	mux.HandleFunc("POST /api/budgets", api.SetBudget)
	mux.HandleFunc("POST /api/transactions", api.SetTransaction)

	if err := http.ListenAndServe(":8080", Log(mux)); err != nil {
		log.Fatal(err)
	}
}

func Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s", r.Method, r.URL.Path, time.Since(t))
	})
}
