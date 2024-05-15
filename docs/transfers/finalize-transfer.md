# Paystack Go SDK

## Transfers

### Finalize Transfer

```go
import "github.com/theghostmc/paystack-go-sdk/pkg/paystack_transfers"
import "github.com/theghostmc/paystack-go-sdk/pkg/paystack_client"

func main() {
    apiKey := "your_api_key"
    client, err := paystack_client.NewClient(apiKey)
    if err != nil {
        log.Fatal(err)
    }

    reqData := paystack_transfers.FinalizeTransferRequest{
        TransferCode: "TRF_vsyqdmlzble3uii",
        OTP:          "928783",
    }

    resp, err := paystack_transfers.FinalizeTransfer(client, reqData)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Transfer Code: %s\n", resp.Data.TransferCode)
    fmt.Printf("Status: %s\n", resp.Data.Status)
}
```