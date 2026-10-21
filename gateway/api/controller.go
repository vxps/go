package api

import (
	"encoding/json"
	"errors"
	"gateway/mapper"
	"gateway/models"
	"ledger"
	"log"
	"net/http"
)

func GetBudgets(w http.ResponseWriter, r *http.Request) {
	budgets := ledger.GetBudgets()
	response := make([]models.BudgetResponse, 0, len(budgets))
	for _, budget := range budgets {
		response = append(response, mapper.ToBudgetResponse(budget))
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Printf("failed to encode response: %v", err)
	}

}

func SetBudget(w http.ResponseWriter, r *http.Request) {
	var request models.CreateBudgetRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Printf("failed to decode request: %v", err)
		WriteJsonError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	model := mapper.ToBudgetDomain(request)
	err = ledger.SetBudget(model)
	if err != nil {
		log.Printf("failed to set budget: %v", err)
		WriteJsonError(w, http.StatusBadRequest, "invalid budget data")
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(mapper.ToBudgetResponse(model))
	if err != nil {
		log.Printf("failed to encode response: %v", err)
	}
}

func GetTransaction(w http.ResponseWriter, r *http.Request) {

	transactions := ledger.ListTransactions()
	response := make([]models.TransactionResponse, 0, len(transactions))
	for _, tr := range transactions {
		response = append(response, mapper.ToTransactionResponse(tr))
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Printf("failed to encode response: %v", err)
	}
}

func SetTransaction(w http.ResponseWriter, r *http.Request) {
	var request models.CreateTransactionRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Printf("failed to decode request: %v", err)
		WriteJsonError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	model := mapper.ToTransactionDomain(request)
	created, err := ledger.AddTransaction(model)
	if err != nil {
		log.Printf("failed to add transaction: %v", err)
		if errors.Is(err, ledger.ErrBudgetExceeded) {
			WriteJsonError(w, http.StatusConflict, "budget exceeded for category")
		} else {
			WriteJsonError(w, http.StatusBadRequest, "invalid transaction data")
		}
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(mapper.ToTransactionResponse(created))
	if err != nil {
		log.Printf("failed to encode response: %v", err)
	}
}

func WriteJsonError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(map[string]string{"error": message})
	if err != nil {
		log.Printf("failed to encode error response: %v", err)
	}
}
