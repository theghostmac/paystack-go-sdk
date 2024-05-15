package paystack_transfers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/theghostmc/paystack-go-sdk/pkg/paystack_client"
	"net/http"
)

// InitiateTransfer initiates a transfer to send money to customers.
func InitiateTransfer(client *paystack_client.Client, reqData InitiateTransferRequest) (*InitiateTransferResponse, error) {
	url := "https://api.paystack.co/transfer"

	// Encode the request body
	jsonData, err := json.Marshal(reqData)
	if err != nil {
		return nil, fmt.Errorf("failed to encode request body: %w", err)
	}

	// Create the HTTP request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
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
	var respData InitiateTransferResponse
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return nil, fmt.Errorf("failed to decode response body: %w", err)
	}

	return &respData, nil
}

// FinalizeTransfer finalizes an initiated transfer.
func FinalizeTransfer(client *paystack_client.Client, reqData FinalizeTransferRequest) (*FinalizeTransferResponse, error) {
	url := "https://api.paystack.co/transfer/finalize_transfer"

	// Encode the request body
	jsonData, err := json.Marshal(reqData)
	if err != nil {
		return nil, fmt.Errorf("failed to encode request body: %w", err)
	}

	// Create the HTTP request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
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
	var respData FinalizeTransferResponse
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return nil, fmt.Errorf("failed to decode response body: %w", err)
	}

	return &respData, nil
}

// InitiateBulkTransfer initiates a bulk transfer.
func InitiateBulkTransfer(client *paystack_client.Client, reqData InitiateBulkTransferRequest) (*InitiateBulkTransferResponse, error) {
	url := "https://api.paystack.co/transfer/bulk"

	// Encode the request body
	jsonData, err := json.Marshal(reqData)
	if err != nil {
		return nil, fmt.Errorf("failed to encode request body: %w", err)
	}

	// Create the HTTP request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
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
	var respData InitiateBulkTransferResponse
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return nil, fmt.Errorf("failed to decode response body: %w", err)
	}

	return &respData, nil
}
