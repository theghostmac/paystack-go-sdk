# Paystack Go SDK

## Dedicated Accounts

### Fetch Bank Providers

```go
import "github.com/theghostmc/paystack-go-sdk/pkg/paystack_accounts"
import "github.com/theghostmc/paystack-go-sdk/pkg/paystack_client"

func main() {
    apiKey := "your_api_key"
    client, err := paystack_client.NewClient(apiKey)
    if err != nil {
        log.Fatal(err)
    }

    resp, err := paystack_accounts.FetchBankProviders(client)
    if err != nil {
        log.Fatal(err)
    }

    for _, provider := range resp.Data {
        fmt.Printf("Bank Provider: %v\n", provider)
    }
}
```