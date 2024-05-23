package main

import (
	"log"

	"github.com/theghostmac/paystack-go-sdk/pkg/paystack_client"
	"github.com/theghostmac/paystack-go-sdk/pkg/paystack_customer"
)

func main() {
	client, err := paystack_client.NewClient("YOUR_SECRET_KEY")
	if err != nil {
		log.Fatalf("Failed to create Paystack client: %v", err)
	}

	reqData := paystack_customer.CreateCustomerRequest{
		Email:     "zero@sum.com",
		FirstName: "Zero",
		LastName:  "Sum",
		Phone:     "+2348123456789",
	}

	resp, err := paystack_customer.CreateCustomer(client, reqData)
	if err != nil {
		log.Fatalf("Failed to create customer: %v", err)
	}

	log.Printf("Create Customer response: %v", resp)
	log.Printf("Customer ID: %d", resp.Data.ID)
	log.Printf("Customer Code: %s", resp.Data.CustomerCode)
	log.Printf("Customer Email: %s", resp.Data.Email)
}
