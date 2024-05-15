# Paystack Go SDK

## Dedicated Accounts

### Create Dedicated Account

```go
import "github.com/theghostmc/paystack-go-sdk/pkg/paystack_accounts"
import "github.com/theghostmc/paystack-go-sdk/pkg/paystack_client"

func main() {
    apiKey := "your_api_key"
    client, err := paystack_client.NewClient(apiKey)
    if err != nil {
        log.Fatal(err)
    }

    reqData := paystack_accounts.CreateDedicatedAccountRequest{
        Customer:      "CUS_xxxxxxxxx",
        PreferredBank: "wema-bank",
    }

    resp, err := paystack_accounts.CreateDedicatedAccount(client, reqData)
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Account ID: %d", resp.Data.ID)
    log.Printf("Account Number: %s", resp.Data.AccountNumber)
    log.Printf("Account Name: %s", resp.Data.AccountName)
}
```
