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
