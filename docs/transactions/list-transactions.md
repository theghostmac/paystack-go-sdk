# Paystack Go SDK

## Transactions

### Initialize Transaction

```go
import "github.com/theghostmac/paystack-go-sdk/pkg/paystack_transactions"
import "github.com/theghostmac/paystack-go-sdk/pkg/paystack_client"

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

    resp, err := paystack_transactions.InitializeTransaction(client, reqData)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(resp.Data.AuthorizationURL)
}
```

## Verify Transaction
```go
import "github.com/theghostmac/paystack-go-sdk/pkg/paystack_transactions"
import "github.com/theghostmac/paystack-go-sdk/pkg/paystack_client"

func main() {
    apiKey := "your_api_key"
    client, err := paystack_client.NewClient(apiKey)
    if err != nil {
        log.Fatal(err)
    }

    reference := "transaction_reference"

    resp, err := paystack_transactions.VerifyTransaction(client, reference)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(resp.Data.Status)
}
```

## List Transactions

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

	queryParams := map[string]string{
		"perPage": "2",
		"page":    "1",
	}

	listResp, err := paystack_transactions.ListTransactions(client, queryParams)
	if err != nil {
		log.Fatal(err)
	}
	for _, transaction := range listResp.Data {
		fmt.Printf("Transaction ID: %d, Status: %s, Amount: %d\n", transaction.ID, transaction.Status, transaction.Amount)
	}
}
```