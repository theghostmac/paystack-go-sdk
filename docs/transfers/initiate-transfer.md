# Paystack Go SDK

## Transfers

### Initiate Transfer

```go
import "github.com/theghostmc/paystack-go-sdk/pkg/paystack_transfers"
import "github.com/theghostmc/paystack-go-sdk/pkg/paystack_client"

func main() {
    apiKey := "your_api_key"
    client, err := paystack_client.NewClient(apiKey)
    if err != nil {
        log.Fatal(err)
    }

    reqData := paystack_transfers.InitiateTransferRequest{
        Source:    "balance",
        Amount:    3794800,
        Recipient: "RCP_gx2wn530m0i3w3m",
        Reason:    "Calm down",
    }

    resp, err := paystack_transfers.InitiateTransfer(client, reqData)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Transfer Code: %s\n", resp.Data.TransferCode)
    fmt.Printf("Status: %s\n", resp.Data.Status)
}
```