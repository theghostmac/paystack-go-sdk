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

Verify the transaction:
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