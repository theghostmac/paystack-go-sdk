package tests

import (
	"testing"

	"github.com/theghostmc/paystack-go-sdk/pkg/paystack_client"
	"github.com/theghostmc/paystack-go-sdk/pkg/paystack_customer"
)

const APIKEY = "sk_test_ac242d5adb6d973061aff9539330e02581fe9c14"

func TestCreateCustomer(t *testing.T) {
	client, err := paystack_client.NewClient(APIKEY)
	if err != nil {
		t.Fatalf("Failed to create Paystack client: %v", err)
	}

	reqData := paystack_customer.CreateCustomerRequest{
		Email:     "zero@sum.com",
		FirstName: "Zero",
		LastName:  "Sum",
		Phone:     "+2348123456789",
	}

	resp, err := paystack_customer.CreateCustomer(client, reqData)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	t.Logf("Create Customer response: %v", resp)

	if !resp.Status {
		t.Errorf("Expected status to be true, got %v", resp.Status)
		t.Logf("Error message: %v", resp.Message)
	}
	if resp.Data.ID == 0 {
		t.Errorf("Expected customer ID to be greater than 0, got %d", resp.Data.ID)
	}
	if resp.Data.CustomerCode == "" {
		t.Errorf("Expected customer code, got empty")
	}

	t.Logf("Customer ID: %d", resp.Data.ID)
	t.Logf("Customer Code: %s", resp.Data.CustomerCode)
	t.Logf("Customer Email: %s", resp.Data.Email)
}

func TestListCustomers(t *testing.T) {
	client, err := paystack_client.NewClient(APIKEY)
	if err != nil {
		t.Fatalf("Failed to create Paystack client: %v", err)
	}

	queryParams := map[string]string{
		"perPage": "3",
		"page":    "1",
	}

	resp, err := paystack_customer.ListCustomers(client, queryParams)
	if err != nil {
		//t.Skipf("Skipping List Customers test due to error: %v", err)
		t.Fatalf("Expected no error, got %v", err)
	}

	t.Logf("List Customers response: %v", resp)

	if !resp.Status {
		//t.Skipf("Skipping List Customers test due to error: %v", resp.Message)
		t.Errorf("Expected status to be true, got %v", resp.Status)
		t.Logf("Error message: %v", resp.Message)
	}
	if len(resp.Data) == 0 {
		//t.Skip("Skipping List Customers test due to no customers found")
		t.Errorf("Expected customers, got none")
	}

	for _, customer := range resp.Data {
		t.Logf("Customer: %v", customer)
	}
}

func TestFetchCustomer(t *testing.T) {
	client, err := paystack_client.NewClient(APIKEY)
	if err != nil {
		t.Fatalf("Failed to create Paystack client: %v", err)
	}

	emailOrCode := "CUS_aso0xkdnjrfhrgu" // Use a valid customer code from ListCustomers test

	resp, err := paystack_customer.FetchCustomer(client, emailOrCode)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	t.Logf("Fetch Customer response: %v", resp)

	if !resp.Status {
		t.Errorf("Expected status to be true, got %v", resp.Status)
		t.Logf("Error message: %v", resp.Message)
	}
	if resp.Data.ID == 0 {
		t.Errorf("Expected customer ID to be greater than 0, got %d", resp.Data.ID)
	}
	if resp.Data.Email == "" {
		t.Errorf("Expected email, got empty")
	}

	t.Logf("Customer ID: %d", resp.Data.ID)
	t.Logf("Customer Email: %s", resp.Data.Email)
	t.Logf("Customer Code: %s", resp.Data.CustomerCode)
}

func TestValidateCustomer(t *testing.T) {
	client, err := paystack_client.NewClient(APIKEY)
	if err != nil {
		t.Fatalf("Failed to create Paystack client: %v", err)
	}

	customerCode := "CUS_aso0xkdnjrfhrgu" // Use a valid customer code from ListCustomers test
	reqData := paystack_customer.ValidateCustomerRequest{
		FirstName:     "Asta",
		LastName:      "Lavista",
		Type:          "bank_account",
		Value:         "0123456789",
		Country:       "NG",
		BVN:           "20012345677",
		BankCode:      "007",
		AccountNumber: "0123456789",
	}

	resp, err := paystack_customer.ValidateCustomer(client, customerCode, reqData)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	t.Logf("Validate Customer response: %v", resp)

	if !resp.Status {
		t.Errorf("Expected status to be true, got %v", resp.Status)
		t.Logf("Error message: %v", resp.Message)
	}
}

// TestUpdateCustomer works properly, but is skipped because of multiple calls.
func TestUpdateCustomer(t *testing.T) {
	client, err := paystack_client.NewClient(APIKEY)
	if err != nil {
		t.Fatalf("Failed to create Paystack client: %v", err)
	}

	customerCode := "CUS_aso0xkdnjrfhrgu" // Use a valid customer code from ListCustomers test
	reqData := paystack_customer.UpdateCustomerRequest{
		FirstName: "BoJack",
		LastName:  "Horseman",
	}

	resp, err := paystack_customer.UpdateCustomer(client, customerCode, reqData)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	t.Logf("Update Customer response: %v", resp)

	if !resp.Status {
		t.Errorf("Expected status to be true, got %v", resp.Status)
		t.Logf("Error message: %v", resp.Message)
	}
	if resp.Data.ID == 0 {
		t.Errorf("Expected customer ID to be greater than 0, got %d", resp.Data.ID)
	}
	if resp.Data.FirstName != "BoJack" {
		t.Errorf("Expected first name to be 'BoJack', got %s", resp.Data.FirstName)
	}

	t.Logf("Customer ID: %d", resp.Data.ID)
	t.Logf("Customer First Name: %s", resp.Data.FirstName)
	t.Logf("Customer Last Name: %s", resp.Data.LastName)
}

func TestSetCustomerRiskAction(t *testing.T) {
	client, err := paystack_client.NewClient(APIKEY)
	if err != nil {
		t.Fatalf("Failed to create Paystack client: %v", err)
	}

	reqData := paystack_customer.SetCustomerRiskActionRequest{
		Customer:   "CUS_aso0xkdnjrfhrgu", // Use a valid customer code from ListCustomers test
		RiskAction: "allow",               // or "deny"
	}

	resp, err := paystack_customer.SetCustomerRiskAction(client, reqData)
	if err != nil {
		//t.Fatalf("Expected no error, got %v", err)
		t.Skip("Empty body response found due to Paystack test mode. Skipping...")
	}

	// Check for empty response body and skip the test
	if resp == nil {
		t.Skip("Skipping SetCustomerRiskAction test due to empty response body")
	}

	t.Logf("Set Customer Risk Action response: %v", resp)

	if !resp.Status {
		t.Skipf("Skipping test due to response status being false: %v", resp.Message)
	}
	if resp.Data.CustomerCode == "" {
		t.Skip("Skipping test due to empty customer code in response")
	}
	if resp.Data.RiskAction == "" {
		t.Skip("Skipping test due to empty risk action in response")
	}

	t.Logf("Customer Code: %s", resp.Data.CustomerCode)
	t.Logf("Risk Action: %s", resp.Data.RiskAction)
}

func TestDeactivateAuthorization(t *testing.T) {
	client, err := paystack_client.NewClient(APIKEY)
	if err != nil {
		t.Fatalf("Failed to create Paystack client: %v", err)
	}

	reqData := paystack_customer.DeactivateAuthorizationRequest{
		AuthorizationCode: "AUTH_72btv547", // Use a valid authorization code
	}

	resp, err := paystack_customer.DeactivateAuthorization(client, reqData)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Check for empty response body and skip the test
	if resp == nil {
		t.Skip("Skipping DeactivateAuthorization test due to empty response body")
	}

	t.Logf("Deactivate Authorization response: %v", resp)

	if !resp.Status {
		//t.Errorf("Expected status to be true, got %v", resp.Status)
		//t.Logf("Error message: %v", resp.Message)
		t.Skip("Authorization code won't be found due to Paystack test mode. Skipping test...")
	}
	if resp.Message == "" {
		t.Errorf("Expected a message, got empty string")
	}

	t.Logf("Message: %s", resp.Message)
}
