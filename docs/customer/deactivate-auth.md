# Paystack Go SDK

## Customers

### Deactivate Authorization

```go
import "github.com/theghostmc/paystack-go-sdk/pkg/paystack_customer"
import "github.com/theghostmc/paystack-go-sdk/pkg/paystack_client"

func main() {
    apiKey := "your_api_key"
    client, err := paystack_client.NewClient(apiKey)
    if err != nil {
        log.Fatal(err)
    }

    reqData := paystack_customer.DeactivateAuthorizationRequest{
        AuthorizationCode: "AUTH_72btv547",
    }

    resp, err := paystack_customer.DeactivateAuthorization(client, reqData)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Response: %v\n", resp)
}
```