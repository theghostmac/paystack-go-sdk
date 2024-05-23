package main

import (
	"fmt"
	"log"

	"github.com/theghostmac/paystack-go-sdk/pkg/paystack_client"
	"github.com/theghostmac/paystack-go-sdk/pkg/paystack_customer"
)

func main() {
	apiKey := "your_api_key"
	client, err := paystack_client.NewClient(apiKey)
	if err != nil {
		log.Fatal(err)
	}

	customerCode := "CUS_aso0xkdnjrfhrgu" // Use a valid customer code
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
		log.Fatal(err)
	}
	fmt.Printf("Validation Status: %t\n", resp.Status)
	fmt.Printf("Message: %s\n", resp.Message)
}
