package main

import (
	"fmt"
	"log"

	"github.com/theghostmc/paystack-go-sdk/pkg/paystack_client"
	"github.com/theghostmc/paystack-go-sdk/pkg/paystack_customer"
)

func main() {
	apiKey := "your_api_key"
	client, err := paystack_client.NewClient(apiKey)
	if err != nil {
		log.Fatal(err)
	}

	reqData := paystack_customer.SetCustomerRiskActionRequest{
		Customer:   "CUS_aso0xkdnjrfhrgu", // Use a valid customer code
		RiskAction: "allow",               // or "deny"
	}

	resp, err := paystack_customer.SetCustomerRiskAction(client, reqData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Customer Code: %s\n", resp.Data.CustomerCode)
	fmt.Printf("Risk Action: %s\n", resp.Data.RiskAction)
}
