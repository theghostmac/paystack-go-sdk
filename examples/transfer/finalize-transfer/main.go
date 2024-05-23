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

	reqData := paystack_transfers.FinalizeTransferRequest{
		TransferCode: "TRF_vsyqdmlzble3uii",
		OTP:          "928783",
	}

	resp, err := paystack_transfers.FinalizeTransfer(client, reqData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Transfers Code: %s\n", resp.Data.TransferCode)
	fmt.Printf("Status: %s\n", resp.Data.Status)
}
