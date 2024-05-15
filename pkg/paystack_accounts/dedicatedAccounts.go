package paystack_accounts

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/theghostmc/paystack-go-sdk/pkg/paystack_client"
	"net/http"
)

// CreateDedicatedAccount creates a dedicated virtual account for an existing customer.
func CreateDedicatedAccount(client *paystack_client.Client, reqData CreateDedicatedAccountRequest) (*CreateDedicatedAccountResponse, error) {
	url := "https://api.paystack.co/dedicated_account"

	reqBody, err := json.Marshal(reqData)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+client.APIKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to perform HTTP request: %w", err)
	}
	defer resp.Body.Close()

	var respData CreateDedicatedAccountResponse
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return nil, fmt.Errorf("failed to decode response body: %w", err)
	}

	return &respData, nil
}
