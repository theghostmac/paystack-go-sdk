# Paystack Go SDK

## Transactions

### Initialize Transaction
Here's how you can use the updated InitializeTransaction function with the Client struct from the paystack_client package:
```go
package main

import (
	"fmt"
	"log"

	"github.com/theghostmac/paystack-go-sdk/pkg/paystack_client"
	"github.com/theghostmac/paystack-go-sdk/pkg/paystack_transactions"
)

func main() {
	apiKey := "your_api_key"
	client, err := paystack_client.NewClient(apiKey)
	if err != nil {
		log.Fatal(err)
	}

	reqData := paystack_transactions.InitializeTransactionRequest{
		Amount: 20000,
		Email:  "customer@email.com",
	}

	initResp, err := paystack_transactions.InitializeTransaction(client, reqData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Authorization URL:", initResp.Data.AuthorizationURL)
}
```