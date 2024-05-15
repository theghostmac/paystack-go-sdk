# Paystack Go SDK

## Transfers

### Verify Transfer

```go
import "github.com/theghostmc/paystack-go-sdk/pkg/paystack_transfers"
import "github.com/theghostmc/paystack-go-sdk/pkg/paystack_client"

func main() {
    apiKey := "your_api_key"
    client, err := paystack_client.NewClient(apiKey)
    if err != nil {
        log.Fatal(err)
    }

    reference := "ref_demo" // Use a valid transfer reference

    resp, err := paystack_transfers.VerifyTransfer(client, reference)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Transfer: %v\n", resp.Data)
}
```