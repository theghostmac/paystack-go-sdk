# Paystack Go SDK

## Dedicated Accounts

### List Dedicated Accounts

```go
import "github.com/theghostmc/paystack-go-sdk/pkg/paystack_accounts"
import "github.com/theghostmc/paystack-go-sdk/pkg/paystack_client"

func main() {
    apiKey := "your_api_key"
    client, err := paystack_client.NewClient(apiKey)
    if err != nil {
        log.Fatal(err)
    }

    queryParams := map[string]string{
        "active": "true",
        "currency": "NGN",
    }

    resp, err := paystack_accounts.ListDedicatedAccounts(client, queryParams)
    if err != nil {
        log.Fatal(err)
    }

    for _, account := range resp.Data {
        fmt.Printf("Dedicated Account: %v\n", account)
    }
}
```