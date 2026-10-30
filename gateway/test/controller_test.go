package test

import (
	"encoding/json"
	"gateway/api"
	"gateway/models"
	"ledger"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestBudgetHttp(t *testing.T) {
	ledger.Reset()
	t.Cleanup(ledger.Reset)

	body := `{"category":"food","limit":5000,"period":"monthly"}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/api/budgets",
		strings.NewReader(body),
	)
	w := httptest.NewRecorder()
	api.SetBudget(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("expected status %d, got %d", http.StatusCreated, w.Code)
	}

	var postResponse models.BudgetResponse
	err1 := json.NewDecoder(w.Body).Decode(&postResponse)
	if err1 != nil {
		t.Fatalf("failed to decode response: %v", err1)
	}
	if w.Header().Get("Content-Type") != "application/json; charset=utf-8" {
		t.Errorf("unexpected Content-Type: %s", w.Header().Get("Content-Type"))
	}
	if postResponse.Category != "food" || postResponse.Limit != 5000 {
		t.Errorf("unexpected POST response: %+v", postResponse)
	}

	getReq := httptest.NewRequest(
		http.MethodGet,
		"/api/budgets",
		nil,
	)
	getW := httptest.NewRecorder()
	api.GetBudgets(getW, getReq)
	if getW.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, getW.Code)
	}
	var getResponse []models.BudgetResponse
	err := json.NewDecoder(getW.Body).Decode(&getResponse)

	if err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	if len(getResponse) != 1 {
		t.Fatalf("expected 1 budget, got %d", len(getResponse))
	}
	if getResponse[0].Category != "food" || getResponse[0].Limit != 5000 {
		t.Fatalf("incorrect data %s - %d", getResponse[0].Category, getResponse[0].Limit)
	}
	contentType := getW.Header().Get("Content-Type")
	if contentType != "application/json; charset=utf-8" {
		t.Errorf("unexpected Content-Type: %s", contentType)
	}
}

func TestNegativeHttp(t *testing.T) {
	t.Run("invalid tests", func(t *testing.T) {
		body := `{"category":"food","limit":`
		req := httptest.NewRequest(
			http.MethodPost,
			"/api/budgets",
			strings.NewReader(body),
		)
		w := httptest.NewRecorder()
		api.SetBudget(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("expected status 400, got %d", w.Code)
		}
	})
}

func TestTransactionsHttp(t *testing.T) {
	ledger.Reset()
	t.Cleanup(ledger.Reset)

	transactionBody := `{
		"amount": 3000,
		"category": "food",
		"description": "ланч",
		"date": "2025-09-10"
		}`

	body := `{"category":"food","limit":5000,"period":"monthly"}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/api/budgets",
		strings.NewReader(body),
	)
	w := httptest.NewRecorder()
	api.SetBudget(w, req)
	if w.Code != http.StatusCreated {
		t.Fatalf("error setting budget")
	}
	t.Run("ok test", func(t *testing.T) {

		tReq := httptest.NewRequest(
			http.MethodPost,
			"/api/transactions",
			strings.NewReader(transactionBody),
		)
		newW := httptest.NewRecorder()
		api.SetTransaction(newW, tReq)
		if newW.Code != http.StatusCreated {
			t.Fatalf("expected status %d, got %d", http.StatusCreated, newW.Code)
		}

		req2 := httptest.NewRequest(
			http.MethodGet,
			"/api/transactions",
			nil,
		)
		w2 := httptest.NewRecorder()
		api.GetTransaction(w2, req2)
		if w2.Code != http.StatusOK {
			t.Fatalf("expected status %d, got %d", http.StatusOK, w2.Code)
		}
		var getResponse []models.TransactionResponse
		err := json.NewDecoder(w2.Body).Decode(&getResponse)
		if err != nil {
			t.Fatalf("failed to decode response: %v", err)
		}
		if len(getResponse) != 1 {
			t.Fatalf("expected 1 transaction, got %d", len(getResponse))
		}
		if getResponse[0].Amount != 3000 || getResponse[0].Category != "food" {
			t.Fatalf("incorrect data %s - %d", getResponse[0].Category, getResponse[0].Amount)
		}
	})

	var errorResponse struct {
		Error string `json:"error"`
	}
	t.Run("exceeded", func(t *testing.T) {
		reqNew := httptest.NewRequest(
			http.MethodPost,
			"/api/transactions",
			strings.NewReader(transactionBody),
		)
		wNew := httptest.NewRecorder()
		api.SetTransaction(wNew, reqNew)
		if wNew.Code != http.StatusConflict {
			t.Fatalf("expected status %d, got %d", http.StatusConflict, wNew.Code)
		}
		err := json.NewDecoder(wNew.Body).Decode(&errorResponse)
		if err != nil {
			t.Fatalf("failed to decode error response: %v", err)
		}
		if errorResponse.Error != "budget exceeded for category" {
			t.Error("diff error")
		}

	})

	t.Run("bad json", func(t *testing.T) {
		body3 := `{"category":"food","li`

		req3 := httptest.NewRequest(
			http.MethodPost,
			"/api/transactions",
			strings.NewReader(body3),
		)
		w3 := httptest.NewRecorder()
		api.SetTransaction(w3, req3)
		if w3.Code != http.StatusBadRequest {
			t.Fatalf("expected status %d, got %d", http.StatusBadRequest, w3.Code)
		}

	})
}
