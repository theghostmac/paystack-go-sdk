# Paystack Go SDK

## Transactions

### Partial Debit

```go
import "github.com/theghostmac/paystack-go-sdk/pkg/paystack_transactions"
import "github.com/theghostmac/paystack-go-sdk/pkg/paystack_client"

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
```
