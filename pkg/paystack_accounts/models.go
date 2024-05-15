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

// DedicatedAccount represents a dedicated virtual account.
type DedicatedAccount struct {
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
	Bank struct {
		Name string `json:"name"`
		ID   int    `json:"id"`
		Slug string `json:"slug"`
	} `json:"bank"`
	ID            int    `json:"id"`
	AccountName   string `json:"account_name"`
	AccountNumber string `json:"account_number"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
	Currency      string `json:"currency"`
	SplitConfig   struct {
		Subaccount string `json:"subaccount"`
	} `json:"split_config"`
	Active   bool `json:"active"`
	Assigned bool `json:"assigned"`
}

// ListDedicatedAccountsResponse represents the response from the Paystack API for listing dedicated virtual accounts.
type ListDedicatedAccountsResponse struct {
	Status  bool               `json:"status"`
	Message string             `json:"message"`
	Data    []DedicatedAccount `json:"data"`
	Meta    struct {
		Total     int `json:"total"`
		Skipped   int `json:"skipped"`
		PerPage   int `json:"perPage"`
		Page      int `json:"page"`
		PageCount int `json:"pageCount"`
	} `json:"meta"`
}

// FetchDedicatedAccountResponse represents the response from the Paystack API for fetching a dedicated virtual account.
type FetchDedicatedAccountResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Transactions          []interface{} `json:"transactions"`
		Subscriptions         []interface{} `json:"subscriptions"`
		Authorizations        []interface{} `json:"authorizations"`
		FirstName             string        `json:"first_name"`
		LastName              string        `json:"last_name"`
		Email                 string        `json:"email"`
		Phone                 string        `json:"phone"`
		Metadata              interface{}   `json:"metadata"`
		Domain                string        `json:"domain"`
		CustomerCode          string        `json:"customer_code"`
		RiskAction            string        `json:"risk_action"`
		ID                    int           `json:"id"`
		Integration           int           `json:"integration"`
		CreatedAt             string        `json:"createdAt"`
		UpdatedAt             string        `json:"updatedAt"`
		CreatedAtAlt          string        `json:"created_at"`
		UpdatedAtAlt          string        `json:"updated_at"`
		TotalTransactions     int           `json:"total_transactions"`
		TotalTransactionValue []interface{} `json:"total_transaction_value"`
		DedicatedAccount      struct {
			ID            int    `json:"id"`
			AccountName   string `json:"account_name"`
			AccountNumber string `json:"account_number"`
			CreatedAt     string `json:"created_at"`
			UpdatedAt     string `json:"updated_at"`
			Currency      string `json:"currency"`
			Active        bool   `json:"active"`
			Assigned      bool   `json:"assigned"`
			Provider      struct {
				ID           int    `json:"id"`
				ProviderSlug string `json:"provider_slug"`
				BankID       int    `json:"bank_id"`
				BankName     string `json:"bank_name"`
			} `json:"provider"`
			Assignment struct {
				AssigneeID   int    `json:"assignee_id"`
				AssigneeType string `json:"assignee_type"`
				AccountType  string `json:"account_type"`
				Integration  int    `json:"integration"`
			} `json:"assignment"`
		} `json:"dedicated_account"`
	} `json:"data"`
}

// RequeryDedicatedAccountResponse represents the response from the Paystack API for requerying a dedicated virtual account.
type RequeryDedicatedAccountResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

// DeactivateDedicatedAccountResponse represents the response from the Paystack API for deactivating a dedicated virtual account.
type DeactivateDedicatedAccountResponse struct {
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
			AssigneeID   int    `json:"assignee_id"`
			AssigneeType string `json:"assignee_type"`
			AssignedAt   string `json:"assigned_at"`
			Integration  int    `json:"integration"`
			AccountType  string `json:"account_type"`
		} `json:"assignment"`
	} `json:"data"`
}
