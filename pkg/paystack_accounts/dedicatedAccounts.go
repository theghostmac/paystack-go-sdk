package paystack_accounts

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/theghostmc/paystack-go-sdk/pkg/paystack_client"
	"net/http"
	"net/url"
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

// AssignDedicatedAccount assigns a dedicated virtual account to a customer.
// It takes the client and request data as arguments and returns the assignment response.
func AssignDedicatedAccount(client *paystack_client.Client, reqData AssignDedicatedAccountRequest) (*AssignDedicatedAccountResponse, error) {
	url := "https://api.paystack.co/dedicated_account/assign"

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
	var respData AssignDedicatedAccountResponse
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return nil, fmt.Errorf("failed to decode response body: %w", err)
	}

	return &respData, nil
}

// ListDedicatedAccounts lists dedicated virtual accounts available on your integration.
// It takes the client and query parameters as arguments and returns the list of dedicated virtual accounts.
func ListDedicatedAccounts(client *paystack_client.Client, queryParams map[string]string) (*ListDedicatedAccountsResponse, error) {
	baseURL := "https://api.paystack.co/dedicated_account"
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse base URL: %w", err)
	}

	// Add query parameters to URL
	q := u.Query()
	for key, value := range queryParams {
		q.Set(key, value)
	}
	u.RawQuery = q.Encode()

	// Create the HTTP request
	req, err := http.NewRequest("GET", u.String(), nil)
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
	var respData ListDedicatedAccountsResponse
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return nil, fmt.Errorf("failed to decode response body: %w", err)
	}

	return &respData, nil
}

// FetchDedicatedAccount fetches the details of a dedicated virtual account on your integration.
// It takes the client and dedicated account ID as arguments and returns the account details.
func FetchDedicatedAccount(client *paystack_client.Client, dedicatedAccountID int) (*FetchDedicatedAccountResponse, error) {
	url := fmt.Sprintf("https://api.paystack.co/dedicated_account/%d", dedicatedAccountID)

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
	var respData FetchDedicatedAccountResponse
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return nil, fmt.Errorf("failed to decode response body: %w", err)
	}

	return &respData, nil
}

// RequeryDedicatedAccount requeries a dedicated virtual account for new transactions.
// It takes the client and query parameters as arguments and returns the requery status.
func RequeryDedicatedAccount(client *paystack_client.Client, queryParams map[string]string) (*RequeryDedicatedAccountResponse, error) {
	baseURL := "https://api.paystack.co/dedicated_account/requery"
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse base URL: %w", err)
	}

	// Add query parameters to URL
	q := u.Query()
	for key, value := range queryParams {
		q.Set(key, value)
	}
	u.RawQuery = q.Encode()

	// Create the HTTP request
	req, err := http.NewRequest("GET", u.String(), nil)
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
	var respData RequeryDedicatedAccountResponse
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return nil, fmt.Errorf("failed to decode response body: %w", err)
	}

	return &respData, nil
}

// DeactivateDedicatedAccount deactivates a dedicated virtual account on your integration.
// It takes the client and the dedicated account ID as arguments and returns the deactivation status.
func DeactivateDedicatedAccount(client *paystack_client.Client, dedicatedAccountID int) (*DeactivateDedicatedAccountResponse, error) {
	url := fmt.Sprintf("https://api.paystack.co/dedicated_account/%d", dedicatedAccountID)

	// Create the HTTP request
	req, err := http.NewRequest("DELETE", url, nil)
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
	var respData DeactivateDedicatedAccountResponse
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return nil, fmt.Errorf("failed to decode response body: %w", err)
	}

	return &respData, nil
}
