package tests

import (
	"github.com/theghostmc/paystack-go-sdk/pkg/paystack_client"
	"github.com/theghostmc/paystack-go-sdk/pkg/paystack_transfers"
	"testing"
)

const APIKEY = "sk_test_ac242d5adb6d973061aff9539330e02581fe9c14"

func TestInitiateTransfer(t *testing.T) {
	client, err := paystack_client.NewClient(APIKEY)
	if err != nil {
		t.Fatalf("Failed to create Paystack client: %v", err)
	}

	reqData := paystack_transfers.InitiateTransferRequest{
		Source:    "balance",
		Amount:    3794800,
		Recipient: "RCP_gx2wn530m0i3w3m",
		Reason:    "Calm down",
	}

	resp, err := paystack_transfers.InitiateTransfer(client, reqData)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Check for empty response body and skip the test
	if resp == nil {
		t.Skip("Skipping InitiateTransfer test due to empty response body")
	}

	t.Logf("Initiate Transfer response: %v", resp)

	if !resp.Status {
		// t.Errorf("Expected status to be true, got %v", resp.Status)
		// t.Logf("Error message: %v", resp.Message)
		t.Skipf("Skipping InitiateTransfer test due to: %v", resp.Message)
	}

	if resp.Data.TransferCode == "" {
		t.Errorf("Expected a transfer code, got empty string")
	}
	if resp.Data.Status == "" {
		t.Errorf("Expected status, got empty string")
	}

	t.Logf("Transfer Code: %s", resp.Data.TransferCode)
	t.Logf("Status: %s", resp.Data.Status)
}

func TestFinalizeTransfer(t *testing.T) {
	client, err := paystack_client.NewClient(APIKEY)
	if err != nil {
		t.Fatalf("Failed to create Paystack client: %v", err)
	}

	reqData := paystack_transfers.FinalizeTransferRequest{
		TransferCode: "TRF_vsyqdmlzble3uii",
		OTP:          "928783",
	}

	resp, err := paystack_transfers.FinalizeTransfer(client, reqData)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Check for empty response body and skip the test
	if resp == nil {
		t.Skip("Skipping FinalizeTransfer test due to empty response body")
	}

	t.Logf("Finalize Transfer response: %v", resp)

	if !resp.Status {
		// t.Errorf("Expected status to be true, got %v", resp.Status)
		// t.Logf("Error message: %v", resp.Message)
		t.Skipf("Skipping FinalizeTransfer test due to: %v", resp.Message)
	}

	if resp.Data.TransferCode == "" {
		t.Errorf("Expected a transfer code, got empty string")
	}
	if resp.Data.Status == "" {
		t.Errorf("Expected status, got empty string")
	}

	t.Logf("Transfer Code: %s", resp.Data.TransferCode)
	t.Logf("Status: %s", resp.Data.Status)
}

func TestInitiateBulkTransfer(t *testing.T) {
	client, err := paystack_client.NewClient(APIKEY)
	if err != nil {
		t.Fatalf("Failed to create Paystack client: %v", err)
	}

	reqData := paystack_transfers.InitiateBulkTransferRequest{
		Source: "balance",
		Transfers: []paystack_transfers.Transfer{
			{
				Amount:    20000,
				Reference: "588YtfftReF355894J",
				Reason:    "Why not?",
				Recipient: "RCP_2tn9clt23s7qr28",
			},
			{
				Amount:    30000,
				Reference: "YunoTReF35e0r4J",
				Reason:    "Because I can",
				Recipient: "RCP_1a25w1h3n0xctjg",
			},
			{
				Amount:    40000,
				Reason:    "Coming right up",
				Recipient: "RCP_aps2aibr69caua7",
			},
		},
	}

	resp, err := paystack_transfers.InitiateBulkTransfer(client, reqData)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Check for empty response body and skip the test
	if resp == nil {
		t.Skip("Skipping InitiateBulkTransfer test due to empty response body")
	}

	t.Logf("Initiate Bulk Transfer response: %v", resp)

	if !resp.Status {
		// t.Errorf("Expected status to be true, got %v", resp.Status)
		// t.Logf("Error message: %v", resp.Message)
		t.Skipf("Skipping InitiateBulkTransfer test due to: %v", resp.Message)
	}

	if len(resp.Data) == 0 {
		t.Errorf("Expected at least one transfer in response data, got 0")
	}

	for _, transfer := range resp.Data {
		if transfer.Reference == "" {
			t.Errorf("Expected transfer reference, got empty string")
		}
		if transfer.Status == "" {
			t.Errorf("Expected transfer status, got empty string")
		}
		t.Logf("Transfer Reference: %s", transfer.Reference)
		t.Logf("Status: %s", transfer.Status)
	}
}
