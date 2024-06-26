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

// FetchCustomerResponse represents the response from the Paystack API for fetching a customer.
type FetchCustomerResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Transactions   []interface{} `json:"transactions"`
		Subscriptions  []interface{} `json:"subscriptions"`
		Authorizations []struct {
			AuthorizationCode string `json:"authorization_code"`
			Bin               string `json:"bin"`
			Last4             string `json:"last4"`
			ExpMonth          string `json:"exp_month"`
			ExpYear           string `json:"exp_year"`
			Channel           string `json:"channel"`
			CardType          string `json:"card_type"`
			Bank              string `json:"bank"`
			CountryCode       string `json:"country_code"`
			Brand             string `json:"brand"`
			Reusable          bool   `json:"reusable"`
			Signature         string `json:"signature"`
			AccountName       string `json:"account_name"`
		} `json:"authorizations"`
		FirstName             string      `json:"first_name"`
		LastName              string      `json:"last_name"`
		Email                 string      `json:"email"`
		Phone                 string      `json:"phone"`
		Metadata              interface{} `json:"metadata"`
		Domain                string      `json:"domain"`
		CustomerCode          string      `json:"customer_code"`
		RiskAction            string      `json:"risk_action"`
		ID                    int         `json:"id"`
		Integration           int         `json:"integration"`
		CreatedAt             string      `json:"createdAt"`
		UpdatedAt             string      `json:"updatedAt"`
		TotalTransactions     int         `json:"total_transactions"`
		TotalTransactionValue interface{} `json:"total_transaction_value"`
		DedicatedAccount      interface{} `json:"dedicated_account"`
		Identified            bool        `json:"identified"`
		Identifications       interface{} `json:"identifications"`
	} `json:"data"`
}

// UpdateCustomerRequest represents the request body for updating a customer.
type UpdateCustomerRequest struct {
	FirstName string                 `json:"first_name,omitempty"`
	LastName  string                 `json:"last_name,omitempty"`
	Phone     string                 `json:"phone,omitempty"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
}

// UpdateCustomerResponse represents the response from the Paystack API for updating a customer.
type UpdateCustomerResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Integration     int                    `json:"integration"`
		FirstName       string                 `json:"first_name"`
		LastName        string                 `json:"last_name"`
		Email           string                 `json:"email"`
		Phone           string                 `json:"phone"`
		Metadata        map[string]interface{} `json:"metadata"`
		Identified      bool                   `json:"identified"`
		Identifications interface{}            `json:"identifications"`
		Domain          string                 `json:"domain"`
		CustomerCode    string                 `json:"customer_code"`
		ID              int                    `json:"id"`
		CreatedAt       string                 `json:"createdAt"`
		UpdatedAt       string                 `json:"updatedAt"`
		Transactions    []interface{}          `json:"transactions"`
		Subscriptions   []interface{}          `json:"subscriptions"`
		Authorizations  []interface{}          `json:"authorizations"`
	} `json:"data"`
}

// ValidateCustomerRequest represents the request body for validating a customer.
type ValidateCustomerRequest struct {
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Type          string `json:"type"`
	Value         string `json:"value"`
	Country       string `json:"country"`
	BVN           string `json:"bvn"`
	BankCode      string `json:"bank_code"`
	AccountNumber string `json:"account_number"`
}

// ValidateCustomerResponse represents the response from the Paystack API for validating a customer.
type ValidateCustomerResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

// SetCustomerRiskActionRequest represents the request body for setting a customer's risk action.
type SetCustomerRiskActionRequest struct {
	Customer   string `json:"customer"`
	RiskAction string `json:"risk_action,omitempty"`
}

// SetCustomerRiskActionResponse represents the response from the Paystack API for setting a customer's risk action.
type SetCustomerRiskActionResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    struct {
		FirstName       string                 `json:"first_name"`
		LastName        string                 `json:"last_name"`
		Email           string                 `json:"email"`
		Phone           string                 `json:"phone"`
		Metadata        map[string]interface{} `json:"metadata"`
		Domain          string                 `json:"domain"`
		Identified      bool                   `json:"identified"`
		Identifications interface{}            `json:"identifications"`
		CustomerCode    string                 `json:"customer_code"`
		RiskAction      string                 `json:"risk_action"`
		ID              int                    `json:"id"`
		Integration     int                    `json:"integration"`
		CreatedAt       string                 `json:"createdAt"`
		UpdatedAt       string                 `json:"updatedAt"`
	} `json:"data"`
}

// DeactivateAuthorizationRequest represents the request body for deactivating an authorization.
type DeactivateAuthorizationRequest struct {
	AuthorizationCode string `json:"authorization_code"`
}

// DeactivateAuthorizationResponse represents the response from the Paystack API for deactivating an authorization.
type DeactivateAuthorizationResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}
