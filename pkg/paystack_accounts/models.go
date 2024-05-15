package paystack_accounts

// CreateDedicatedAccountRequest represents the request body for creating a dedicated virtual account.
type CreateDedicatedAccountRequest struct {
	Customer      string `json:"customer"`
	PreferredBank string `json:"preferred_bank,omitempty"`
}

// CreateDedicatedAccountResponse represents the response from the Paystack API for creating a dedicated virtual account.
type CreateDedicatedAccountResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Bank struct {
			Name string `json:"name"`
			ID   int    `json:"id"`
			Slug string `json:"slug"`
		} `json:"bank"`
		AccountName   string `json:"account_name"`
		AccountNumber string `json:"account_number"`
		Assigned      bool   `json:"assigned"`
		Currency      string `json:"currency"`
		Metadata      string `json:"metadata"`
		Active        bool   `json:"active"`
		ID            int    `json:"id"`
		CreatedAt     string `json:"created_at"`
		UpdatedAt     string `json:"updated_at"`
		Assignment    struct {
			Integration  int    `json:"integration"`
			AssigneeID   int    `json:"assignee_id"`
			AssigneeType string `json:"assignee_type"`
			Expired      bool   `json:"expired"`
			AccountType  string `json:"account_type"`
			AssignedAt   string `json:"assigned_at"`
		} `json:"assignment"`
		Customer struct {
			ID                       int    `json:"id"`
			FirstName                string `json:"first_name"`
			LastName                 string `json:"last_name"`
			Email                    string `json:"email"`
			CustomerCode             string `json:"customer_code"`
			Phone                    string `json:"phone"`
			RiskAction               string `json:"risk_action"`
			InternationalFormatPhone string `json:"international_format_phone"`
		} `json:"customer"`
	} `json:"data"`
}

// AssignDedicatedAccountRequest represents the request body for assigning a dedicated virtual account.
type AssignDedicatedAccountRequest struct {
	Email         string `json:"email"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Phone         string `json:"phone"`
	PreferredBank string `json:"preferred_bank"`
	Country       string `json:"country"`
}

// AssignDedicatedAccountResponse represents the response from the Paystack API for assigning a dedicated virtual account.
type AssignDedicatedAccountResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}
