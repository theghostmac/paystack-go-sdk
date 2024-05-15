# Paystack Go SDK

## Transfers

### List Transfers

```go
import "github.com/theghostmc/paystack-go-sdk/pkg/paystack_transfers"
import "github.com/theghostmc/paystack-go-sdk/pkg/paystack_client"

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

    resp, err := paystack_transfers.ListTransfers(client, queryParams)
    if err != nil {
        log.Fatal(err)
    }

    for _, transfer := range resp.Data {
        fmt.Printf("Transfer: %v\n", transfer)
    }
}
```