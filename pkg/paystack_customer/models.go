package paystack_customer

// Customer represents a customer in the Paystack system.
type Customer struct {
	Integration  int         `json:"integration"`
	FirstName    string      `json:"first_name"`
	LastName     string      `json:"last_name"`
	Email        string      `json:"email"`
	Phone        string      `json:"phone"`
	Metadata     interface{} `json:"metadata"`
	Domain       string      `json:"domain"`
	CustomerCode string      `json:"customer_code"`
	RiskAction   string      `json:"risk_action"`
	ID           int         `json:"id"`
	CreatedAt    string      `json:"createdAt"`
	UpdatedAt    string      `json:"updatedAt"`
}

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

// ListCustomersResponse represents the response from the Paystack API for listing customers.
type ListCustomersResponse struct {
	Status  bool       `json:"status"`
	Message string     `json:"message"`
	Data    []Customer `json:"data"`
	Meta    struct {
		Next     string `json:"next"`
		Previous string `json:"previous"`
		PerPage  int    `json:"perPage"`
	} `json:"meta"`
}
