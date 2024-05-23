package main

import (
	"fmt"
	"log"

	"github.com/theghostmac/paystack-go-sdk/pkg/paystack_client"
	"github.com/theghostmac/paystack-go-sdk/pkg/paystack_transfers"
)

func main() {
	apiKey := "your_api_key"
	client, err := paystack_client.NewClient(apiKey)
	if err != nil {
		log.Fatal(err)
	}

	reference := "ref_demo" // Use a valid transfer reference

	resp, err := paystack_transfers.VerifyTransfer(client, reference)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Transfer: %v\n", resp.Data)
}
