package main

import (
	"fmt"
	"github.com/theghostmc/paystack-go-sdk/pkg/paystack_client"
	"github.com/theghostmc/paystack-go-sdk/pkg/paystack_transactions"
	"log"
)

func main() {
	apiKey := "your_api_key"
	client, err := paystack_client.NewClient(apiKey)
	if err != nil {
		log.Fatal(err)
	}

	reqData := paystack_transactions.PartialDebitRequest{
		AuthorizationCode: "AUTH_72btv547",
		Currency:          "NGN",
		Amount:            "20000",
		Email:             "customer@email.com",
	}

	resp, err := paystack_transactions.PartialDebit(client, reqData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Partial Debit Response: %+v\n", resp)
}
