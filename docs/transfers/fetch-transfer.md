# Paystack Go SDK

## Transfers

### Fetch Transfer

```go
import "github.com/theghostmc/paystack-go-sdk/pkg/paystack_transfers"
import "github.com/theghostmc/paystack-go-sdk/pkg/paystack_client"

func main() {
    apiKey := "your_api_key"
    client, err := paystack_client.NewClient(apiKey)
    if err != nil {
        log.Fatal(err)
    }

    idOrCode := "TRF_2x5j67tnnw1t98k" // Use a valid transfer ID or code

    resp, err := paystack_transfers.FetchTransfer(client, idOrCode)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Transfer: %v\n", resp.Data)
}
```