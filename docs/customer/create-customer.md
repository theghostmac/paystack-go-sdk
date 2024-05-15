# Paystack Go SDK

## Customers

### Create Customer

```go
import "github.com/theghostmc/paystack-go-sdk/pkg/paystack_customer"
import "github.com/theghostmc/paystack-go-sdk/pkg/paystack_client"

func main() {
    apiKey := "your_api_key"
    client, err := paystack_client.NewClient(apiKey)
    if err != nil {
        log.Fatal(err)
    }

    reqData := paystack_customer.CreateCustomerRequest{
        Email:     "zero@sum.com",
        FirstName: "Zero",
        LastName:  "Sum",
        Phone:     "+2348123456789",
    }

    resp, err := paystack_customer.CreateCustomer(client, reqData)
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Customer ID: %d", resp.Data.ID)
    log.Printf("Customer Code: %s", resp.Data.CustomerCode)
    log.Printf("Customer Email: %s", resp.Data.Email)
}
```