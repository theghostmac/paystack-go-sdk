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

	reqData := paystack_customer.DeactivateAuthorizationRequest{
		AuthorizationCode: "AUTH_72btv547",
	}

	resp, err := paystack_customer.DeactivateAuthorization(client, reqData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Response: %v\n", resp)
}
