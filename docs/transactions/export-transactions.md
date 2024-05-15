# Paystack Go SDK

## Transactions

### Export Transactions

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
        "perPage": "50",
        "page":    "1",
    }

    resp, err := paystack_transactions.ExportTransactions(client, queryParams)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Export Path: %s\n", resp.Data.Path)
}
```
