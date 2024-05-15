# Paystack Go SDK

## Dedicated Accounts

### Requery Dedicated Account

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
        "account_number": "1234567890", // Replace with a valid virtual account number for testing
        "provider_slug":  "wema-bank",  // Replace with a valid provider slug for testing
        "date":           "2023-05-30",
    }

    resp, err := paystack_accounts.RequeryDedicatedAccount(client, queryParams)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Message: %s\n", resp.Message)
}
```