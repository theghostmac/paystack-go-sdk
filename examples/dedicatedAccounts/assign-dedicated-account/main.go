package main

import (
	"fmt"
	"log"

	"github.com/theghostmc/paystack-go-sdk/pkg/paystack_accounts"
	"github.com/theghostmc/paystack-go-sdk/pkg/paystack_client"
)

func main() {
	apiKey := "your_api_key"
	client, err := paystack_client.NewClient(apiKey)
	if err != nil {
		log.Fatal(err)
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
		log.Fatal(err)
	}
	fmt.Printf("Response: %v\n", resp)
}
