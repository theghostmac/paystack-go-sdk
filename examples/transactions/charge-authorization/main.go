package main

import (
	"fmt"
	"github.com/theghostmac/paystack-go-sdk/pkg/paystack_client"
	"github.com/theghostmac/paystack-go-sdk/pkg/paystack_transactions"
	"log"
)

func main() {
	apiKey := "your_api_key"
	client, err := paystack_client.NewClient(apiKey)
	if err != nil {
		log.Fatal(err)
	}

	reqData := paystack_transactions.ChargeAuthorizationRequest{
		Amount:            "20000",
		Email:             "customer@email.com",
		AuthorizationCode: "AUTH_72btv547",
		Reference:         "unique_reference_12345",
		Currency:          "NGN",
		Metadata:          `{"custom_fields":[{"display_name":"Cart ID","variable_name": "cart_id","value": "8393"}]}`,
		Channels:          []string{"card"},
		SubAccount:        "ACCT_8f4s1eq7ml6rlzj",
		TransactionCharge: 100,
		Bearer:            "account",
		Queue:             true,
	}

	chargeResp, err := paystack_transactions.ChargeAuthorization(client, reqData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Charge Status: %s, Amount: %d\n", chargeResp.Data.Status, chargeResp.Data.Amount)
}
