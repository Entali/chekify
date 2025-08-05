package bank

import (
	_ "bytes"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"
)

var (
	baseURL     = "https://api.monobank.ua"
	baseURLLock sync.RWMutex
)

func OverrideBaseURL(url string) string {
	baseURLLock.Lock()
	defer baseURLLock.Unlock()
	old := baseURL
	baseURL = url
	return old
}

func TestFetchRecentTransactions(t *testing.T) {
	// Мокуємо відповідь сервера
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[
			{
				"id": "123",
				"time": 1753734919,
				"description": "FORA Київ",
				"mcc": 5411,
				"amount": -5000,
				"currencyCode": 980
			}
		]`))
	}))
	defer mockServer.Close()

	// Підміняємо базовий URL на наш тестовий сервер
	old := OverrideBaseURL(mockServer.URL)
	defer OverrideBaseURL(old)

	// Викликаємо функцію
	since := time.Now().Add(-time.Hour) // за останню годину
	transactions, err := FetchRecentTransactions("test-account", since)
	if err != nil {
		t.Fatalf("Fetch failed: %v", err)
	}

	if len(transactions) != 1 {
		t.Fatalf("Expected 1 transaction, got %d", len(transactions))
	}

	if transactions[0].Description != "FORA Київ" {
		t.Errorf("Unexpected description: %s", transactions[0].Description)
	}
}
