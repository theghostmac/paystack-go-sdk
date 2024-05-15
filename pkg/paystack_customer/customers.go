package paystack_customer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/theghostmc/paystack-go-sdk/pkg/paystack_client"
	"net/http"
	"net/url"
)

// CreateCustomer creates a customer on your integration.
// It takes the client and request data as arguments and returns the response data and an error, if any.
func CreateCustomer(client *paystack_client.Client, reqData CreateCustomerRequest) (*CreateCustomerResponse, error) {
	url := "https://api.paystack.co/customer"

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

	var respData CreateCustomerResponse
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return nil, fmt.Errorf("failed to decode response body: %w", err)
	}

	return &respData, nil
}

// ListCustomers lists customers available on your integration.
// It takes the client and query parameters as arguments and returns the list of customers.
func ListCustomers(client *paystack_client.Client, queryParams map[string]string) (*ListCustomersResponse, error) {
	baseURL := "https://api.paystack.co/customer"
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
	var respData ListCustomersResponse
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return nil, fmt.Errorf("failed to decode response body: %w", err)
	}

	return &respData, nil
}

// FetchCustomer gets the details of a customer on your integration.
// It takes the client and customer email or code as arguments and returns the customer details.
func FetchCustomer(client *paystack_client.Client, emailOrCode string) (*FetchCustomerResponse, error) {
	url := fmt.Sprintf("https://api.paystack.co/customer/%s", emailOrCode)

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
	var respData FetchCustomerResponse
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return nil, fmt.Errorf("failed to decode response body: %w", err)
	}

	return &respData, nil
}

// UpdateCustomer updates a customer's details on your integration.
// It takes the client, customer code, and request data as arguments and returns the updated customer details.
func UpdateCustomer(client *paystack_client.Client, customerCode string, reqData UpdateCustomerRequest) (*UpdateCustomerResponse, error) {
	url := fmt.Sprintf("https://api.paystack.co/customer/%s", customerCode)

	// Marshal the request body to JSON
	reqBody, err := json.Marshal(reqData)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %w", err)
	}

	// Create the HTTP request
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(reqBody))
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
	var respData UpdateCustomerResponse
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return nil, fmt.Errorf("failed to decode response body: %w", err)
	}

	return &respData, nil
}

// ValidateCustomer validates a customer's identity on your integration.
// It takes the client, customer code, and request data as arguments and returns the validation response.
func ValidateCustomer(client *paystack_client.Client, customerCode string, reqData ValidateCustomerRequest) (*ValidateCustomerResponse, error) {
	url := fmt.Sprintf("https://api.paystack.co/customer/%s/identification", customerCode)

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
	var respData ValidateCustomerResponse
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return nil, fmt.Errorf("failed to decode response body: %w", err)
	}

	return &respData, nil
}
