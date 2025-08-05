package bank

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

const monoBaseURL = "https://api.monobank.ua"

type Transaction struct {
	ID              string `json:"id"`
	Time            int64  `json:"time"`
	Description     string `json:"description"`
	MCC             int    `json:"mcc"`
	OriginalMCC     int    `json:"originalMcc"`
	Amount          int64  `json:"amount"`
	OperationAmount int64  `json:"operationAmount"`
	CurrencyCode    int    `json:"currencyCode"`
	CommissionRate  int64  `json:"commissionRate"`
	CashbackAmount  int64  `json:"cashbackAmount"`
	BalanceAfter    int64  `json:"balance"`
	Hold            bool   `json:"hold"`
	ReceiptID       string `json:"receiptId"`
	CounterEdrpou   string `json:"counterEdrpou"`
	CounterIban     string `json:"counterIban"`
	CounterName     string `json:"counterName"`
}

func FetchRecentTransactions(accountID string, from time.Time) ([]Transaction, error) {
	token := os.Getenv("MONO_TOKEN")
	if token == "" {
		return nil, fmt.Errorf("MONO_TOKEN env variable not set")
	}

	url := fmt.Sprintf("%s/personal/statement/%s/%d", monoBaseURL, accountID, from.Unix())

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-Token", token)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("monobank API error: %s", resp.Status)
	}

	var transactions []Transaction
	if err := json.NewDecoder(resp.Body).Decode(&transactions); err != nil {
		return nil, err
	}

	return transactions, nil
}
