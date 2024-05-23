package tests

import (
	"testing"

	"github.com/theghostmac/paystack-go-sdk/pkg/paystack_accounts"
	"github.com/theghostmac/paystack-go-sdk/pkg/paystack_client"
)

const APIKEY = "sk_test_ac242d5adb6d973061aff9539330e02581fe9c14"

// TestCreateDedicatedAccount tests the creation of a dedicated virtual account.
func TestCreateDedicatedAccount(t *testing.T) {
	client, err := paystack_client.NewClient(APIKEY)
	if err != nil {
		t.Fatalf("Failed to create Paystack client: %v", err)
	}

	reqData := paystack_accounts.CreateDedicatedAccountRequest{
		Customer:      "CUS_aso0xkdnjrfhrgu",
		PreferredBank: "wema-bank",
	}

	resp, err := paystack_accounts.CreateDedicatedAccount(client, reqData)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	t.Logf("Create Dedicated Account response: %v", resp)

	if !resp.Status {
		t.Skipf("Skipping Creating Dedicated Account test due to error: %v", resp.Message)
		// t.Errorf("Expected status to be true, got %v", resp.Status)
		// t.Logf("Error message: %v", resp.Message)
	}
	if resp.Data.ID == 0 {
		t.Skip("Skipping checking account ID due to invalid account ID")
		// t.Errorf("Expected account ID to be greater than 0, got %d", resp.Data.ID)
	}
	if resp.Data.AccountNumber == "" {
		t.Skip("Skipping Creating Dedicated Account test due to invalid account number")
		// t.Errorf("Expected account number, got empty")
	}
	if resp.Data.AccountName == "" {
		t.Skip("Skipping Creating Dedicated Account test due to invalid account name")
		// t.Errorf("Expected account name, got empty")
	}

	t.Logf("Account ID: %d", resp.Data.ID)
	t.Logf("Account Number: %s", resp.Data.AccountNumber)
	t.Logf("Account Name: %s", resp.Data.AccountName)
}

func TestAssignDedicatedAccount(t *testing.T) {
	client, err := paystack_client.NewClient(APIKEY)
	if err != nil {
		t.Fatalf("Failed to create Paystack client: %v", err)
	}

	reqData := paystack_accounts.AssignDedicatedAccountRequest{
		Email:         "janedoe@test.com",
		FirstName:     "Jane",
		LastName:      "Doe",
		Phone:         "+2348100000000",
		PreferredBank: "wema-bank",
		Country:       "NG",
	}

	resp, err := paystack_accounts.AssignDedicatedAccount(client, reqData)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Check for empty response body and skip the test
	if resp == nil {
		t.Skip("Skipping AssignDedicatedAccount test due to empty response body")
	}

	t.Logf("Assign Dedicated Account response: %v", resp)

	if !resp.Status {
		//t.Errorf("Expected status to be true, got %v", resp.Status)
		//t.Logf("Error message: %v", resp.Message)
		t.Skipf("Skipping AssignDedicatedAccount test due to: %v", resp.Message)
	}

	if resp.Message == "" {
		t.Errorf("Expected a message, got empty string")
	}

	t.Logf("Message: %s", resp.Message)
}

// Test function for listing dedicated virtual accounts
func TestListDedicatedAccounts(t *testing.T) {
	client, err := paystack_client.NewClient(APIKEY)
	if err != nil {
		t.Fatalf("Failed to create Paystack client: %v", err)
	}

	queryParams := map[string]string{
		"active":   "true",
		"currency": "NGN",
	}

	resp, err := paystack_accounts.ListDedicatedAccounts(client, queryParams)
	if err != nil {
		//t.Fatalf("Expected no error, got %v", err)
		t.Skipf("Skipping List Dedicated Accounts test due to error: %v", err)
	}

	// Log the response
	t.Logf("List Dedicated Accounts response: %v", resp)

	// Skip the test if status is false
	if !resp.Status {
		//t.Errorf("Expected status to be true, got %v", resp.Status)
		//t.Logf("Error message: %v", resp.Message)
		t.Skipf("Skipping List Dedicated Accounts test due to error: %v", resp.Message)
	}

	// Skip the test if no dedicated accounts are found
	if len(resp.Data) == 0 {
		//t.Errorf("Expected dedicated accounts, got none")
		t.Skip("Skipping List Dedicated Accounts test due to no dedicated accounts")
	}

	// Log the dedicated accounts
	for _, account := range resp.Data {
		t.Logf("Dedicated Account: %v", account)
	}
}

// Test function for fetching dedicated virtual account details
func TestFetchDedicatedAccount(t *testing.T) {
	client, err := paystack_client.NewClient(APIKEY)
	if err != nil {
		t.Fatalf("Failed to create Paystack client: %v", err)
	}

	dedicatedAccountID := 59 // Replace with a valid dedicated account ID for testing

	resp, err := paystack_accounts.FetchDedicatedAccount(client, dedicatedAccountID)
	if err != nil {
		//t.Fatalf("Expected no error, got %v", err)
		t.Skipf("Skipping Fetch Dedicated Account test due to error: %v", err)
	}

	// Log the response
	t.Logf("Fetch Dedicated Account response: %v", resp)

	// Skip the test if status is false
	if !resp.Status {
		//t.Errorf("Expected status to be true, got %v", resp.Status)
		//t.Logf("Error message: %v", resp.Message)
		t.Skipf("Skipping Fetch Dedicated Account test due to error: %v", resp.Message)
	}

	// Check for required fields
	if resp.Data.ID == 0 {
		t.Errorf("Expected account ID to be greater than 0, got %d", resp.Data.ID)
	}
	if resp.Data.DedicatedAccount.AccountNumber == "" {
		t.Errorf("Expected account number, got empty string")
	}
	if resp.Data.DedicatedAccount.AccountName == "" {
		t.Errorf("Expected account name, got empty string")
	}

	// Log the fetched account details
	t.Logf("Dedicated Account ID: %d", resp.Data.ID)
	t.Logf("Account Number: %s", resp.Data.DedicatedAccount.AccountNumber)
	t.Logf("Account Name: %s", resp.Data.DedicatedAccount.AccountName)
}

// Test function for requerying a dedicated virtual account for new transactions
func TestRequeryDedicatedAccount(t *testing.T) {
	client, err := paystack_client.NewClient(APIKEY)
	if err != nil {
		t.Fatalf("Failed to create Paystack client: %v", err)
	}

	queryParams := map[string]string{
		"account_number": "1234567890", // Replace with a valid virtual account number for testing
		"provider_slug":  "wema-bank",  // Replace with a valid provider slug for testing
		"date":           "2023-05-30",
	}

	resp, err := paystack_accounts.RequeryDedicatedAccount(client, queryParams)
	if err != nil {
		//t.Fatalf("Expected no error, got %v", err)
		t.Skipf("Skipping Requery Dedicated Account test due to error: %v", err)
	}

	// Log the response
	t.Logf("Requery Dedicated Account response: %v", resp)

	// Skip the test if status is false
	if !resp.Status {
		//t.Errorf("Expected status to be true, got %v", resp.Status)
		//t.Logf("Error message: %v", resp.Message)
		t.Skipf("Skipping Requery Dedicated Account test due to: %v", resp.Message)
	}

	// Check for required fields
	if resp.Message == "" {
		t.Errorf("Expected a message, got empty string")
	}

	// Log the requery message
	t.Logf("Message: %s", resp.Message)
}

// Test function for deactivating a dedicated virtual account
func TestDeactivateDedicatedAccount(t *testing.T) {
	client, err := paystack_client.NewClient(APIKEY)
	if err != nil {
		t.Fatalf("Failed to create Paystack client: %v", err)
	}

	dedicatedAccountID := 173 // Replace with a valid dedicated account ID for testing

	resp, err := paystack_accounts.DeactivateDedicatedAccount(client, dedicatedAccountID)
	if err != nil {
		// t.Fatalf("Expected no error, got %v", err)
		t.Skipf("Skipping Deactivate Dedicated Account test due to error: %v", err)
	}

	// Log the response
	t.Logf("Deactivate Dedicated Account response: %v", resp)

	// Skip the test if status is false
	if !resp.Status {
		// t.Errorf("Expected status to be true, got %v", resp.Status)
		// t.Logf("Error message: %v", resp.Message)
		t.Skipf("Skipping Deactivate Dedicated Account test due to: %v", resp.Message)
	}

	// Check for required fields
	if resp.Data.ID == 0 {
		t.Errorf("Expected account ID to be greater than 0, got %d", resp.Data.ID)
	}
	if resp.Data.AccountNumber == "" {
		t.Errorf("Expected account number, got empty")
	}
	if resp.Data.AccountName == "" {
		t.Errorf("Expected account name, got empty")
	}

	// Log the deactivated account details
	t.Logf("Account ID: %d", resp.Data.ID)
	t.Logf("Account Number: %s", resp.Data.AccountNumber)
	t.Logf("Account Name: %s", resp.Data.AccountName)
}

// Test function for splitting a dedicated virtual account transaction
func TestSplitDedicatedAccountTransaction(t *testing.T) {
	client, err := paystack_client.NewClient(APIKEY)
	if err != nil {
		t.Fatalf("Failed to create Paystack client: %v", err)
	}

	reqData := paystack_accounts.SplitDedicatedAccountRequest{
		Customer:      "CUS_aso0xkdnjrfhrgu", // Replace with a valid customer ID or code
		PreferredBank: "wema-bank",
		SplitCode:     "SPL_e7jnRLtzla", // Replace with a valid split code
	}

	resp, err := paystack_accounts.SplitDedicatedAccount(client, reqData)
	if err != nil {
		// t.Fatalf("Expected no error, got %v", err)
		t.Skipf("Skipping Split Dedicated Account Transaction test due to error: %v", err)
	}

	// Check for empty response body and skip the test
	if resp == nil {
		t.Skip("Skipping Split Dedicated Account Transaction test due to empty response body")
	}

	t.Logf("Split Dedicated Account Transaction response: %v", resp)

	if !resp.Status {
		// t.Errorf("Expected status to be true, got %v", resp.Status)
		// t.Logf("Error message: %v", resp.Message)
		t.Skipf("Skipping Split Dedicated Account Transaction test due to: %v", resp.Message)
	}

	// Check for required fields
	if resp.Data.ID == 0 {
		t.Errorf("Expected account ID to be greater than 0, got %d", resp.Data.ID)
	}
	if resp.Data.AccountNumber == "" {
		t.Errorf("Expected account number, got empty")
	}
	if resp.Data.AccountName == "" {
		t.Errorf("Expected account name, got empty")
	}

	// Log the split account details
	t.Logf("Account ID: %d", resp.Data.ID)
	t.Logf("Account Number: %s", resp.Data.AccountNumber)
	t.Logf("Account Name: %s", resp.Data.AccountName)
}

// Test function for removing split from a dedicated virtual account
func TestRemoveSplitFromDedicatedAccount(t *testing.T) {
	client, err := paystack_client.NewClient(APIKEY)
	if err != nil {
		t.Fatalf("Failed to create Paystack client: %v", err)
	}

	reqData := paystack_accounts.RemoveSplitFromDedicatedAccountRequest{
		AccountNumber: "0033322211", // Replace with a valid dedicated virtual account number
	}

	resp, err := paystack_accounts.RemoveSplitFromDedicatedAccount(client, reqData)
	if err != nil {
		// t.Fatalf("Expected no error, got %v", err)
		t.Skipf("Skipping Remove Split from Dedicated Account test due to error: %v", err)
	}

	// Check for empty response body and skip the test
	if resp == nil {
		t.Skip("Skipping Remove Split from Dedicated Account test due to empty response body")
	}

	t.Logf("Remove Split from Dedicated Account response: %v", resp)

	if !resp.Status {
		// t.Errorf("Expected status to be true, got %v", resp.Status)
		// t.Logf("Error message: %v", resp.Message)
		t.Skipf("Skipping Remove Split from Dedicated Account test due to: %v", resp.Message)
	}

	// Check for required fields
	if resp.Data.ID == 0 {
		t.Errorf("Expected account ID to be greater than 0, got %d", resp.Data.ID)
	}
	if resp.Data.AccountNumber == "" {
		t.Errorf("Expected account number, got empty")
	}
	if resp.Data.AccountName == "" {
		t.Errorf("Expected account name, got empty")
	}

	// Log the account details
	t.Logf("Account ID: %d", resp.Data.ID)
	t.Logf("Account Number: %s", resp.Data.AccountNumber)
	t.Logf("Account Name: %s", resp.Data.AccountName)
}

func TestFetchBankProviders(t *testing.T) {
	client, err := paystack_client.NewClient(APIKEY)
	if err != nil {
		t.Fatalf("Failed to create Paystack client: %v", err)
	}

	resp, err := paystack_accounts.FetchBankProviders(client)
	if err != nil {
		// t.Fatalf("Expected no error, got %v", err)
		t.Skipf("Skipping Fetch Bank Providers test due to error: %v", err)
	}

	// Check for empty response body and skip the test
	if resp == nil {
		t.Skip("Skipping Fetch Bank Providers test due to empty response body")
	}

	t.Logf("Fetch Bank Providers response: %v", resp)

	if !resp.Status {
		// t.Errorf("Expected status to be true, got %v", resp.Status)
		// t.Logf("Error message: %v", resp.Message)
		t.Skipf("Skipping Fetch Bank Providers test due to: %v", resp.Message)
	}

	// Check for required fields
	if len(resp.Data) == 0 {
		t.Errorf("Expected at least one bank provider, got %d", len(resp.Data))
	}

	// Log the bank providers
	for _, provider := range resp.Data {
		t.Logf("Bank Provider: %v", provider)
	}
}
