package paystack_transactions

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/theghostmc/paystack-go-sdk/pkg/paystack_client"
	"net/http"
	"net/url"
)

// InitializeTransaction represents the request body for initializing a transaction.
// It takes the client and the request data.
func InitializeTransaction(client *paystack_client.Client, reqData InitializeTransactionRequest) (*InitializeTransactionResponse, error) {
	url := "https://api.paystack.co/transaction/initialize"

	// Marshal the request body to JSON
	reqBody, err := json.Marshal(reqData)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %w", err)
	}

	// Create the HTTP request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	// Set headers
	req.Header.Set("Authorization", "Bearer "+client.APIKey)
	req.Header.Set("Content-Type", "application/json")

	// Perform the HTTP request
	resp, err := client.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to perform HTTP request: %w", err)
	}
	defer resp.Body.Close()

	// Decode the response body
	var respData InitializeTransactionResponse
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return nil, fmt.Errorf("failed to decode response body: %w", err)
	}

	return &respData, nil
}

// VerifyTransaction verifies the status of a transaction on Paystack.
// It takes the client and transaction reference as arguments and returns the transaction details.
func VerifyTransaction(client *paystack_client.Client, reference string) (*VerifyTransactionResponse, error) {
	url := fmt.Sprintf("https://api.paystack.co/transaction/verify/%s", reference)

	// Create the HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	// Set headers
	req.Header.Set("Authorization", "Bearer "+client.APIKey)

	// Perform the HTTP request
	resp, err := client.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to perform HTTP request: %w", err)
	}
	defer resp.Body.Close()

	// Decode the response body
	var respData VerifyTransactionResponse
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return nil, fmt.Errorf("failed to decode response body: %w", err)
	}

	return &respData, nil
}

// ListTransactions lists transactions carried out on your integration.
// It takes the client and optional query parameters as arguments and returns the list of transactions.
func ListTransactions(client *paystack_client.Client, queryParams map[string]string) (*ListTransactionsResponse, error) {
	baseURL := "https://api.paystack.co/transaction"

	// Parse the base URL
	url, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse base URL: %w", err)
	}

	// Add query parameters to the URL
	q := url.Query()
	for key, value := range queryParams {
		q.Add(key, value)
	}
	url.RawQuery = q.Encode()

	// Create the HTTP request
	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	// Set headers
	req.Header.Set("Authorization", "Bearer "+client.APIKey)

	// Perform the HTTP request
	resp, err := client.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to perform HTTP request: %w", err)
	}
	defer resp.Body.Close()

	// Decode the response body
	var respData ListTransactionsResponse
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return nil, fmt.Errorf("failed to decode response body: %w", err)
	}

	return &respData, nil
}

// FetchTransaction fetches the details of a transaction carried out on your integration.
// It takes the client and transaction ID as arguments and returns the transaction details.
func FetchTransaction(client *paystack_client.Client, transactionID uint64) (*FetchTransactionResponse, error) {
	url := fmt.Sprintf("https://api.paystack.co/transaction/%d", transactionID)

	// Create the HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	// Set headers
	req.Header.Set("Authorization", "Bearer "+client.APIKey)

	// Perform the HTTP request
	resp, err := client.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to perform HTTP request: %w", err)
	}
	defer resp.Body.Close()

	// Decode the response body
	var respData FetchTransactionResponse
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return nil, fmt.Errorf("failed to decode response body: %w", err)
	}

	return &respData, nil
}

// ChargeAuthorization charges an authorization on Paystack.
// It takes the client and request data as arguments and returns the response from Paystack.
func ChargeAuthorization(client *paystack_client.Client, reqData ChargeAuthorizationRequest) (*ChargeAuthorizationResponse, error) {
	url := "https://api.paystack.co/transaction/charge_authorization"

	// Marshal the request body to JSON
	reqBody, err := json.Marshal(reqData)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %w", err)
	}

	// Create the HTTP request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	// Set headers
	req.Header.Set("Authorization", "Bearer "+client.APIKey)
	req.Header.Set("Content-Type", "application/json")

	// Perform the HTTP request
	resp, err := client.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to perform HTTP request: %w", err)
	}
	defer resp.Body.Close()

	// Decode the response body
	var respData ChargeAuthorizationResponse
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return nil, fmt.Errorf("failed to decode response body: %w", err)
	}

	return &respData, nil
}
