package paystack_customer

// CreateCustomerRequest represents the request body for creating a customer.
type CreateCustomerRequest struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Phone     string `json:"phone,omitempty"`
}

// CreateCustomerResponse represents the response from the Paystack API for creating a customer.
type CreateCustomerResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Email           string      `json:"email"`
		Integration     int         `json:"integration"`
		Domain          string      `json:"domain"`
		CustomerCode    string      `json:"customer_code"`
		ID              int         `json:"id"`
		Identified      bool        `json:"identified"`
		Identifications interface{} `json:"identifications"`
		CreatedAt       string      `json:"createdAt"`
		UpdatedAt       string      `json:"updatedAt"`
	} `json:"data"`
}
