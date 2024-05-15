package paystack_transactions

type Transaction struct {
	ID     int    `json:"id"`
	Amount int    `json:"amount"`
	Status string `json:"status"`
}

type InitializeTransactionRequest struct {
	Amount            int      `json:"amount"`
	Email             string   `json:"email"`
	Currency          string   `json:"currency,omitempty"`
	Reference         string   `json:"reference,omitempty"`
	CallbackURL       string   `json:"callback_url,omitempty"`
	Plan              string   `json:"plan,omitempty"`
	InvoiceLimit      int      `json:"invoice_limit,omitempty"`
	Metadata          string   `json:"metadata,omitempty"`
	Channels          []string `json:"channels,omitempty"`
	SplitCode         string   `json:"split_code,omitempty"`
	SubAccount        string   `json:"subaccount,omitempty"`
	TransactionCharge int      `json:"transaction_charge,omitempty"`
	Bearer            string   `json:"bearer,omitempty"`
}

type InitializeTransactionResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    struct {
		AuthorizationURL string `json:"authorization_url"`
		AccessCode       string `json:"access_code"`
		Reference        string `json:"reference"`
	} `json:"data"`
}

// VerifyTransactionResponse represents the response from the Paystack API for verifying a transaction.
type VerifyTransactionResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    struct {
		ID              uint64 `json:"id"`
		Domain          string `json:"domain"`
		Status          string `json:"status"`
		Reference       string `json:"reference"`
		Amount          int    `json:"amount"`
		Message         string `json:"message"`
		GatewayResponse string `json:"gateway_response"`
		Channel         string `json:"channel"`
		Currency        string `json:"currency"`
		IPAddress       string `json:"ip_address"`
		Metadata        string `json:"metadata"`
		Log             struct {
			StartTime int           `json:"start_time"`
			TimeSpent int           `json:"time_spent"`
			Attempts  int           `json:"attempts"`
			Errors    int           `json:"errors"`
			Success   bool          `json:"success"`
			Mobile    bool          `json:"mobile"`
			Input     []interface{} `json:"input"`
			History   []struct {
				Type    string `json:"type"`
				Message string `json:"message"`
				Time    int    `json:"time"`
			} `json:"history"`
		} `json:"log"`
		Fees          int `json:"fees"`
		Authorization struct {
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
		} `json:"authorization"`
		Customer struct {
			ID                       uint64 `json:"id"`
			FirstName                string `json:"first_name"`
			LastName                 string `json:"last_name"`
			Email                    string `json:"email"`
			CustomerCode             string `json:"customer_code"`
			Phone                    string `json:"phone"`
			Metadata                 string `json:"metadata"`
			RiskAction               string `json:"risk_action"`
			InternationalFormatPhone string `json:"international_format_phone"`
		} `json:"customer"`
		Plan               interface{} `json:"plan"`
		Subaccount         interface{} `json:"subaccount"`
		OrderID            interface{} `json:"order_id"`
		PaidAt             string      `json:"paid_at"`
		CreatedAt          string      `json:"created_at"`
		RequestedAmount    int         `json:"requested_amount"`
		PosTransactionData interface{} `json:"pos_transaction_data"`
		Source             interface{} `json:"source"`
		FeesBreakdown      interface{} `json:"fees_breakdown"`
		TransactionDate    string      `json:"transaction_date"`
		PlanObject         interface{} `json:"plan_object"`
		Split              interface{} `json:"split"`
	} `json:"data"`
}

// ListTransactionsResponse represents the response from the Paystack API for listing transactions.
type ListTransactionsResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    []struct {
		ID              uint64 `json:"id"`
		Domain          string `json:"domain"`
		Status          string `json:"status"`
		Reference       string `json:"reference"`
		Amount          int    `json:"amount"`
		Message         string `json:"message"`
		GatewayResponse string `json:"gateway_response"`
		PaidAt          string `json:"paid_at"`
		CreatedAt       string `json:"created_at"`
		Channel         string `json:"channel"`
		Currency        string `json:"currency"`
		IPAddress       string `json:"ip_address"`
		Metadata        string `json:"metadata"`
		Customer        struct {
			FirstName    string `json:"first_name"`
			LastName     string `json:"last_name"`
			Email        string `json:"email"`
			Phone        string `json:"phone"`
			CustomerCode string `json:"customer_code"`
		} `json:"customer"`
		Authorization struct {
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
		} `json:"authorization"`
		Log struct {
			StartTime int           `json:"start_time"`
			TimeSpent int           `json:"time_spent"`
			Attempts  int           `json:"attempts"`
			Errors    int           `json:"errors"`
			Success   bool          `json:"success"`
			Mobile    bool          `json:"mobile"`
			Input     []interface{} `json:"input"`
			History   []struct {
				Type    string `json:"type"`
				Message string `json:"message"`
				Time    int    `json:"time"`
			} `json:"history"`
		} `json:"log"`
		Fees                 int         `json:"fees"`
		Timeline             interface{} `json:"timeline"`
		RequestedAmount      int         `json:"requested_amount"`
		AuthorizationDetails struct {
			Code string `json:"code"`
			Pin  string `json:"pin"`
		} `json:"authorization_details"`
	} `json:"data"`
	Meta struct {
		Total     int       `json:"total"`
		Skipped   int       `json:"skipped"`
		PerPage   CustomInt `json:"perPage"`
		Page      CustomInt `json:"page"`
		PageCount CustomInt `json:"pageCount"`
	} `json:"meta"`
}

// FetchTransactionResponse represents the response from the Paystack API for fetching a transaction.
type FetchTransactionResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    struct {
		ID              uint64   `json:"id"`
		Domain          string   `json:"domain"`
		Status          string   `json:"status"`
		Reference       string   `json:"reference"`
		Amount          int      `json:"amount"`
		Message         string   `json:"message"`
		GatewayResponse string   `json:"gateway_response"`
		PaidAt          string   `json:"paid_at"`
		CreatedAt       string   `json:"created_at"`
		Channel         string   `json:"channel"`
		Currency        string   `json:"currency"`
		IPAddress       string   `json:"ip_address"`
		Metadata        Metadata `json:"metadata"`
		Log             struct {
			StartTime int           `json:"start_time"`
			TimeSpent int           `json:"time_spent"`
			Attempts  int           `json:"attempts"`
			Errors    int           `json:"errors"`
			Success   bool          `json:"success"`
			Mobile    bool          `json:"mobile"`
			Input     []interface{} `json:"input"`
			History   []struct {
				Type    string `json:"type"`
				Message string `json:"message"`
				Time    int    `json:"time"`
			} `json:"history"`
		} `json:"log"`
		Fees          int         `json:"fees"`
		FeesSplit     interface{} `json:"fees_split"`
		Authorization struct {
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
		} `json:"authorization"`
		Customer struct {
			ID                       uint64 `json:"id"`
			FirstName                string `json:"first_name"`
			LastName                 string `json:"last_name"`
			Email                    string `json:"email"`
			CustomerCode             string `json:"customer_code"`
			Phone                    string `json:"phone"`
			Metadata                 string `json:"metadata"`
			RiskAction               string `json:"risk_action"`
			InternationalFormatPhone string `json:"international_format_phone"`
		} `json:"customer"`
		Plan            interface{} `json:"plan"`
		Subaccount      interface{} `json:"subaccount"`
		OrderID         interface{} `json:"order_id"`
		RequestedAmount int         `json:"requested_amount"`
	} `json:"data"`
}

// ChargeAuthorizationRequest represents the request body for charging an authorization.
type ChargeAuthorizationRequest struct {
	Amount            string   `json:"amount"`
	Email             string   `json:"email"`
	AuthorizationCode string   `json:"authorization_code"`
	Reference         string   `json:"reference"`
	Currency          string   `json:"currency"`
	Metadata          string   `json:"metadata"`
	Channels          []string `json:"channels"`
	SubAccount        string   `json:"subaccount"`
	TransactionCharge int      `json:"transaction_charge"`
	Bearer            string   `json:"bearer"`
	Queue             bool     `json:"queue"`
}

// ChargeAuthorizationResponse represents the response from the Paystack API for charging an authorization.
type ChargeAuthorizationResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Amount          int    `json:"amount"`
		Currency        string `json:"currency"`
		TransactionDate string `json:"transaction_date"`
		Status          string `json:"status"`
		Reference       string `json:"reference"`
		Domain          string `json:"domain"`
		Metadata        string `json:"metadata"`
		GatewayResponse string `json:"gateway_response"`
		Message         string `json:"message"`
		Channel         string `json:"channel"`
		IPAddress       string `json:"ip_address"`
		Log             string `json:"log"`
		Fees            int    `json:"fees"`
		Authorization   struct {
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
		} `json:"authorization"`
		Customer struct {
			ID           uint64 `json:"id"`
			FirstName    string `json:"first_name"`
			LastName     string `json:"last_name"`
			Email        string `json:"email"`
			CustomerCode string `json:"customer_code"`
			Phone        string `json:"phone"`
			Metadata     string `json:"metadata"`
			RiskAction   string `json:"risk_action"`
		} `json:"customer"`
		Plan interface{} `json:"plan"`
		ID   uint64      `json:"id"`
	} `json:"data"`
}

// TransactionTimelineResponse represents the response from the Paystack API for viewing the transaction timeline.
type TransactionTimelineResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    struct {
		TimeSpent      int           `json:"time_spent"`
		Attempts       int           `json:"attempts"`
		Authentication interface{}   `json:"authentication"`
		Errors         int           `json:"errors"`
		Success        bool          `json:"success"`
		Mobile         bool          `json:"mobile"`
		Input          []interface{} `json:"input"`
		Channel        string        `json:"channel"`
		History        []struct {
			Type    string `json:"type"`
			Message string `json:"message"`
			Time    int    `json:"time"`
		} `json:"history"`
	} `json:"data"`
}

// TransactionTotalsResponse represents the response from the Paystack API for fetching transaction totals.
type TransactionTotalsResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    struct {
		TotalTransactions     int `json:"total_transactions"`
		UniqueCustomers       int `json:"unique_customers"`
		TotalVolume           int `json:"total_volume"`
		TotalVolumeByCurrency []struct {
			Currency string `json:"currency"`
			Amount   int    `json:"amount"`
		} `json:"total_volume_by_currency"`
		PendingTransfers           int `json:"pending_transfers"`
		PendingTransfersByCurrency []struct {
			Currency string `json:"currency"`
			Amount   int    `json:"amount"`
		} `json:"pending_transfers_by_currency"`
	} `json:"data"`
}

// ExportTransactionsResponse represents the response from the Paystack API for exporting transactions.
type ExportTransactionsResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Path string `json:"path"`
	} `json:"data"`
}
