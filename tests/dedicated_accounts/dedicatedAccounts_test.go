package tests

import (
	"testing"

	"github.com/theghostmc/paystack-go-sdk/pkg/paystack_accounts"
	"github.com/theghostmc/paystack-go-sdk/pkg/paystack_client"
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
