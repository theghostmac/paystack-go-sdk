package tests

import (
	"github.com/theghostmc/paystack-go-sdk/pkg/paystack_client"
	"github.com/theghostmc/paystack-go-sdk/pkg/paystack_transactions"
	"testing"
)

const APIKEY = "sk_test_ac242d5adb6d973061aff9539330e02581fe9c14"

func TestInitializeTransaction(t *testing.T) {
	client, err := paystack_client.NewClient(APIKEY)
	if err != nil {
		t.Fatalf("Failed to create Paystack client: %v", err)
	}

	reqData := paystack_transactions.InitializeTransactionRequest{
		Amount: 20000,
		Email:  "customer@email.com",
	}

	resp, err := paystack_transactions.InitializeTransaction(client, reqData)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	txnReference := resp.Data.Reference
	t.Logf("Transaction reference: %v", txnReference)

	if !resp.Status {
		t.Errorf("Expected status to be true, got %v", resp.Status)
	}
	if resp.Data.AuthorizationURL == "" {
		t.Errorf("Expected authorization_url, got empty")
	}

	// Log the transaction ID for use in other tests
	fetchResp, err := paystack_transactions.VerifyTransaction(client, txnReference)
	if err != nil {
		t.Fatalf("Failed to verify transaction: %v", err)
	}
	t.Logf("Transaction ID: %d", fetchResp.Data.ID)
}

func TestVerifyTransaction(t *testing.T) {
	client, err := paystack_client.NewClient(APIKEY)
	if err != nil {
		t.Fatalf("Failed to create Paystack client: %v", err)
	}

	reference := "xd5zqbqilr" // Gotten from initialize transaction test.

	resp, err := paystack_transactions.VerifyTransaction(client, reference)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if !resp.Status {
		t.Errorf("Expected status to be true, got %v", resp.Status)
	}
	if resp.Data.ID == 0 {
		t.Errorf("Expected transaction ID, got 0")
	}
}

func TestListTransactions(t *testing.T) {
	client, err := paystack_client.NewClient(APIKEY)
	if err != nil {
		t.Fatalf("Failed to create Paystack client: %v", err)
	}

	queryParams := map[string]string{
		"perPage": "2",
		"page":    "1",
	}

	resp, err := paystack_transactions.ListTransactions(client, queryParams)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if !resp.Status {
		t.Errorf("Expected status to be true, got %v", resp.Status)
	}
	if len(resp.Data) == 0 {
		t.Errorf("Expected transactions, got none")
	}
}

func TestFetchTransaction(t *testing.T) {
	client, err := paystack_client.NewClient(APIKEY)
	if err != nil {
		t.Fatalf("Failed to create Paystack client: %v", err)
	}

	transactionID := uint64(3795390315) // Replace with the transaction ID logged from TestInitializeTransaction

	resp, err := paystack_transactions.FetchTransaction(client, transactionID)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if !resp.Status {
		t.Errorf("Expected status to be true, got %v", resp.Status)
	}
	if resp.Data.ID == 0 {
		t.Errorf("Expected transaction ID, got 0")
	}
}
