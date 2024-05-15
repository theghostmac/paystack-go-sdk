package main

import (
	"fmt"
	"log"

	"github.com/theghostmc/paystack-go-sdk/pkg/paystack_client"
	"github.com/theghostmc/paystack-go-sdk/pkg/paystack_transfers"
)

func main() {
	apiKey := "your_api_key"
	client, err := paystack_client.NewClient(apiKey)
	if err != nil {
		log.Fatal(err)
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
		log.Fatal(err)
	}

	for _, transfer := range resp.Data {
		fmt.Printf("Transfer Reference: %s\n", transfer.Reference)
		fmt.Printf("Status: %s\n", transfer.Status)
	}
}
