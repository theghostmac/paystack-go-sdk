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
