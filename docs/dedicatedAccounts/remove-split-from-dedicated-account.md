# Paystack Go SDK

## Dedicated Accounts

### Remove Split from Dedicated Account

```go
import "github.com/theghostmc/paystack-go-sdk/pkg/paystack_accounts"
import "github.com/theghostmc/paystack-go-sdk/pkg/paystack_client"

func main() {
    apiKey := "your_api_key"
    client, err := paystack_client.NewClient(apiKey)
    if err != nil {
        log.Fatal(err)
    }

    reqData := paystack_accounts.RemoveSplitFromDedicatedAccountRequest{
        AccountNumber: "0033322211", // Replace with a valid dedicated virtual account number
    }

    resp, err := paystack_accounts.RemoveSplitFromDedicatedAccount(client, reqData)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Account ID: %d\n", resp.Data.ID)
    fmt.Printf("Account Number: %s\n", resp.Data.AccountNumber)
    fmt.Printf("Account Name: %s\n", resp.Data.AccountName)
}
```