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

### Verify Transaction

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

### List Transactions
```go
import "github.com/theghostmac/paystack-go-sdk/pkg/paystack_transactions"
import "github.com/theghostmac/paystack-go-sdk/pkg/paystack_client"

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

    resp, err := paystack_transactions.ListTransactions(client, queryParams)
    if err != nil {
        log.Fatal(err)
    }
    for _, transaction := range resp.Data {
        fmt.Printf("Transaction ID: %d, Status: %s, Amount: %d\n", transaction.ID, transaction.Status, transaction.Amount)
    }
}
```

### Fetch Transaction

```go
import "github.com/theghostmac/paystack-go-sdk/pkg/paystack_transactions"
import "github.com/theghostmac/paystack-go-sdk/pkg/paystack_client"

func main() {
    apiKey := "your_api_key"
    client, err := paystack_client.NewClient(apiKey)
    if err != nil {
        log.Fatal(err)
    }

    transactionID := uint64(3795390315) // Example transaction ID

    resp, err := paystack_transactions.FetchTransaction(client, transactionID)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Transaction ID: %d, Status: %s, Amount: %d\n", resp.Data.ID, resp.Data.Status, resp.Data.Amount)
}
```